package handler

import (
	"context"
	"encoding/json"
	"github.com/asim/go-micro/v3/logger"
	"go.mongodb.org/mongo-driver/bson"
	"ops/pkg/model/mgodb"
	"ops/pkg/model/rdb"
	"ops/pkg/utils"
	pbMessage "ops/proto/message"
	"time"
)

type Message struct {
	MsgType   int64  `json:"msg_type"`
	Timestamp string `json:"timestamp"`
}

type Txn struct {
	Message
	TxnHash     string  `json:"txn_hash"`
	TxnStatus   string  `json:"txn_status"`
	FromAddress string  `json:"from_address"`
	ToAddress   string  `json:"to_address"`
	Amount      float64 `json:"amount"`
}

type MintedOcard struct {
	TransactionHash string `bson:"transaction_hash"`
	Uid             string `bson:"uid"`
	Timestamp       string `bson:"timestamp"`
}

type Follow struct {
	Message
	Uid string `json:"uid"`
}

type Like struct {
	Message
	Oid     string `json:"oid"`
	Uid     string `json:"uid"`
	LikeNum int64  `json:"like_num"`
	Content string `json:"content"`
}

type Foword struct {
	Message
	Oid string `json:"oid"`
	Uid string `json:"uid"`
}

type Comment struct {
	Message
	Oid        string `json:"oid"`
	Uid        string `json:"uid"`
	Content    string `json:"content"`
	ContentOid string `json:"contentOid"`
}

type PrivateMessage struct {
	Message
	Uid     string `json:"uid"`
	Content string `json:"content"`
}

type OcardNotEnough Message

type CreateCard struct {
	Message
	TxnHash   string `json:"txn_hash"`
	Amount    int64  `json:"amount"`
	TxnStatus string `json:"txn_status"`
}

type Mention struct {
	Message
	Mentioning string `json:"mentioning"`
	Oid        string `json:"oid"`
	Uid        string `json:"uid"`
	LikeNum    int64  `json:"like_num"`
	Content    string `json:"content"`
}

func (m Message) Get(ctx context.Context, num *pbMessage.Num, results *pbMessage.Msgs) error {
	defer func() {
		logger.Infof("calling Get success, request=%+v,  response=%+v", num, results)
	}()
	redisListName := getUserMqName(num.Uid)
	for i := 0; i < int(num.Num); i++ {
		result, err := rdb.RPop(redisListName)
		if err != nil {
			logger.Error("从redis取出失败", err)
		} else {
			msg, err := m.MessageUnmarshal2Msg(result)
			if err != nil {
				logger.Error("MessageUnmarshal2Msg失败", err)
				return err
			}
			results.Data = append(results.Data, msg)
		}
	}
	return nil
}

func (m Message) GetAll(ctx context.Context, userID *pbMessage.UserID, results *pbMessage.Msgs) error {
	defer func() {
		logger.Infof("calling GetAll success, request=%+v,  response=%+v", userID, results)
	}()
	redisList := getUserMqName(userID.Uid)
	lenth, err := rdb.LLen(redisList)
	if err != nil {
		logger.Error(err)
		return err
	}
	for i := 0; i < lenth; i++ {
		result, err := rdb.RPop(redisList)
		if err != nil {
			logger.Error(err)
		} else {
			msg, err := m.MessageUnmarshal2Msg(result)
			if err != nil {
				logger.Error(err, "unmarshal error")
				return err
			}
			results.Data = append(results.Data, msg)
		}
	}
	return nil
}

func (m Message) PutOne(ctx context.Context, msg *pbMessage.Msg, result *pbMessage.Result) error {
	defer func() {
		logger.Info(msg)
	}()
	if msg.MsgType == 8 && msg.Uid == "" {
		// MsgType为8 msg.Uid为空 需要用address字段在表找到uid
		var (
			mintedOcard MintedOcard
			err         error
		)
		mongoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.property"), utils.GetConfigStr("mongodb.collection.minted_ocard"))
		err = mongoClient.FindOne(bson.M{"transaction_hash": msg.CreateCard.TxnHash}).Decode(&mintedOcard)

		if err != nil {
			logger.Error(err, "通过txn解析uid错误", msg.CreateCard.TxnHash)
			return err
		}
		msg.Uid = mintedOcard.Uid

	} else if msg.MsgType == 1 {
		if msg.Uid == "" {
			// 去redis查 没有带uid就是success或者failed
			uid, errStr := rdb.Get(msg.Txn.TxnHash)
			if errStr != "" {
				logger.Info(errStr)
			}

			msg.Uid = uid
			err := rdb.Del(msg.Txn.TxnHash)
			if err != nil {
				logger.Error(err)
			}
		} else {
			// 带uid 存取哈希和uid映射
			err := rdb.SetM(msg.Txn.TxnHash, msg.Uid, time.Hour*2400)
			if err != nil {
				logger.Error(err, "set失败")
			}
		}
	}

	redisListName := getUserMqName(msg.Uid)
	err := m.PushMsg2Redis(msg, redisListName)
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info("存入redis成功")
	}
	return err
}

func (m Message) PutBatch(ctx context.Context, msgs *pbMessage.Msgs, result *pbMessage.Result) error {
	defer func() {
		logger.Infof("calling PutBatch success, request=%+v,  response=%+v", msgs, "no return")
	}()
	for _, msg := range msgs.Data {
		if (msg.MsgType == 8 || msg.MsgType == 1) && msg.Uid == "" {
			// MsgType为1 msg.Uid为空 需要用address字段在表找到uid
			var (
				mintedOcard MintedOcard
				err         error
				collection  string
				txnHash     string
			)
			if msg.MsgType == 8 {
				collection = "mongodb.collection.minted_ocard"
				txnHash = msg.CreateCard.TxnHash
			} else if msg.MsgType == 1 {
				collection = "mongodb.collection.withdraw_opspoint_history"
				txnHash = msg.Txn.TxnHash
			}
			mongoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.property"), utils.GetConfigStr(collection))
			err = mongoClient.FindOne(bson.M{"transaction_hash": txnHash}).Decode(&mintedOcard)
			if err != nil {
				logger.Error(err, "通过txn解析uid错误")
				return err
			}
			logger.Info(mintedOcard)
			msg.Uid = mintedOcard.Uid

			redisListName := getUserMqName(msg.Uid)
			err = m.PushMsg2Redis(msg, redisListName)
			if err != nil {
				logger.Error("message微服务PushMsg2Redis报错", err)
			}
		}
	}
	return nil
}

func (m Message) PushMsg2Redis(msg *pbMessage.Msg, redisListName string) error {
	msgType := msg.MsgType
	var model interface{}
	var err error
	switch msgType {
	case 1:
		model = Txn{
			Message: Message{
				MsgType:   msgType,
				Timestamp: msg.Txn.Timestamp,
			},
			TxnHash:     msg.Txn.TxnHash,
			TxnStatus:   msg.Txn.TxnStatus,
			ToAddress:   msg.Txn.ToAddress,
			FromAddress: msg.Txn.FromAddress,
			Amount:      msg.Txn.Amount,
		}
	case 2:
		model = Follow{
			Message: Message{
				MsgType:   msgType,
				Timestamp: msg.Follow.Timestamp,
			},
			Uid: msg.Follow.Uid,
		}
	case 3:
		model = Like{
			Message: Message{
				MsgType:   msgType,
				Timestamp: msg.Like.Timestamp,
			},
			Oid:     msg.Like.Oid,
			Uid:     msg.Like.Uid,
			LikeNum: msg.Like.LikeNum,
			Content: msg.Like.Content,
		}
	case 4:
		model = Foword{
			Message: Message{
				MsgType:   msgType,
				Timestamp: msg.Foword.Timestamp,
			},
			Oid: msg.Foword.Oid,
			Uid: msg.Foword.Uid,
		}
	case 5:
		model = Comment{
			Message: Message{
				MsgType:   msgType,
				Timestamp: msg.Comment.Timestamp,
			},
			Oid:        msg.Comment.Oid,
			Uid:        msg.Comment.Uid,
			Content:    msg.Comment.Content,
			ContentOid: msg.Comment.ContentOid,
		}
	case 6:
		model = PrivateMessage{
			Message: Message{
				MsgType:   msgType,
				Timestamp: msg.PrivateMessage.Timestamp,
			},
			Uid:     msg.PrivateMessage.Uid,
			Content: msg.PrivateMessage.Content,
		}
	case 7:
		model = OcardNotEnough{
			MsgType:   msgType,
			Timestamp: msg.OcardNotEnough.Timestamp,
		}
	case 8:
		model = CreateCard{
			Message: Message{
				MsgType:   msgType,
				Timestamp: msg.CreateCard.Timestamp,
			},
			TxnHash:   msg.CreateCard.TxnHash,
			Amount:    msg.CreateCard.Amount,
			TxnStatus: msg.CreateCard.TxnStatus,
		}
	case 9:
		model = Mention{
			Message: Message{
				MsgType:   msgType,
				Timestamp: msg.Mention.Timestamp,
			},
			Mentioning: msg.Mention.Mentioning,
			Oid:        msg.Mention.Oid,
			Uid:        msg.Mention.Uid,
			Content:    msg.Mention.Content,
		}
	}
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(model)
	if err != nil {
		logger.Error("字符串生成失败", err)
	}
	err = rdb.LPush(redisListName, bytes)
	if err != nil {
		logger.Error("Redis可能出现问题", err)
	}
	return err
}

func (m Message) MessageUnmarshal2Msg(redisResult string) (*pbMessage.Msg, error) {
	var msg pbMessage.Msg
	err := json.Unmarshal([]byte(redisResult), &m)
	if err != nil {
		logger.Error(err)
		return &msg, err
	}
	msg.MsgType = m.MsgType
	switch m.MsgType {
	case 1:
		txn := &Txn{}
		err := json.Unmarshal([]byte(redisResult), txn)
		if err != nil {
			return &msg, err
		}
		msg.Txn = &pbMessage.Txn{
			TxnHash:     txn.TxnHash,
			TxnStatus:   txn.TxnStatus,
			MsgType:     1,
			Timestamp:   txn.Timestamp,
			ToAddress:   txn.ToAddress,
			FromAddress: txn.FromAddress,
			Amount:      txn.Amount,
		}

	case 2:
		follow := &Follow{}
		err := json.Unmarshal([]byte(redisResult), follow)
		if err != nil {
			return &msg, err
		}
		msg.Follow = &pbMessage.Follow{
			Uid:       follow.Uid,
			MsgType:   2,
			Timestamp: follow.Timestamp,
		}
	case 3:
		like := &Like{}
		err := json.Unmarshal([]byte(redisResult), like)
		if err != nil {
			return &msg, err
		}
		msg.Like = &pbMessage.Like{
			Oid:       like.Oid,
			Uid:       like.Uid,
			LikeNum:   like.LikeNum,
			MsgType:   3,
			Timestamp: like.Timestamp,
			Content:   like.Content,
		}
	case 4:
		foword := &Foword{}
		err := json.Unmarshal([]byte(redisResult), foword)
		if err != nil {
			return &msg, err
		}
		msg.Foword = &pbMessage.Foword{
			Uid:       foword.Uid,
			Oid:       foword.Oid,
			MsgType:   4,
			Timestamp: foword.Timestamp,
		}
	case 5:
		comment := &Comment{}
		err := json.Unmarshal([]byte(redisResult), comment)
		if err != nil {
			return &msg, err
		}
		msg.Comment = &pbMessage.Comment{
			Uid:        comment.Uid,
			Oid:        comment.Oid,
			Content:    comment.Content,
			ContentOid: comment.ContentOid,
			MsgType:    5,
			Timestamp:  comment.Timestamp,
		}
	case 6:
		privateMessage := &PrivateMessage{}
		err := json.Unmarshal([]byte(redisResult), privateMessage)
		if err != nil {
			return &msg, err
		}
		msg.PrivateMessage = &pbMessage.PrivateMessage{
			Uid:       privateMessage.Uid,
			Content:   privateMessage.Content,
			MsgType:   6,
			Timestamp: privateMessage.Timestamp,
		}
	case 7:
		ocardNotEnough := &OcardNotEnough{}
		err := json.Unmarshal([]byte(redisResult), ocardNotEnough)
		if err != nil {
			return &msg, err
		}
		msg.OcardNotEnough = &pbMessage.OcardNotEnough{
			MsgType:   7,
			Timestamp: ocardNotEnough.Timestamp,
		}
	case 8:
		createCard := &CreateCard{}
		err := json.Unmarshal([]byte(redisResult), createCard)
		if err != nil {
			return &msg, err
		}
		msg.CreateCard = &pbMessage.CreateCard{
			MsgType:   8,
			Timestamp: createCard.Timestamp,
			TxnHash:   createCard.TxnHash,
			TxnStatus: createCard.TxnStatus,
			Amount:    createCard.Amount,
		}
	case 9:
		mention := &Mention{}
		err := json.Unmarshal([]byte(redisResult), mention)
		if err != nil {
			return &msg, err
		}
		msg.Mention = &pbMessage.Mention{
			MsgType:    9,
			Timestamp:  mention.Timestamp,
			Mentioning: mention.Mentioning,
			Uid:        mention.Uid,
			Oid:        mention.Oid,
			Content:    mention.Content,
		}
	}
	return &msg, err
}

func getUserMqName(uid string) string {
	return uid + "mq"
}
