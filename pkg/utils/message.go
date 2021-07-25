package utils

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	pbMessage "ops/proto/message"
	"strconv"
	"time"
)

var microClient pbMessage.OperateMessageService

func init() {
	consulReg := consul.NewRegistry(
		registry.Addrs(GetConfigStr("micro.addr")))
	MicroSer := micro.NewService(
		micro.Registry(consulReg),
	)
	microClient = pbMessage.NewOperateMessageService("message", MicroSer.Client())
}

// 建卡pending
func SendCreateCardMsgPending(txnHash, timestamp string, amount int64) error {
	err := SendCreateCardMsg(txnHash, timestamp, "pending", amount)
	return err
}

// 建卡failed
func SendCreateCardMsgFailed(txnHash, timestamp string, amount int64) error {
	err := SendCreateCardMsg(txnHash, timestamp, "failed", amount)
	return err
}

// 建卡success
func SendCreateCardMsgSuccess(txnHash, timestamp string, amount int64) error {
	err := SendCreateCardMsg(txnHash, timestamp, "success", amount)
	return err
}

// 卡函数
func SendCreateCardMsg(txnHash, timestamp, txnStatus string, amount int64) error {
	card := &pbMessage.Msg{
		MsgType: 8,
		CreateCard: &pbMessage.CreateCard{
			TxnHash:   txnHash,
			TxnStatus: txnStatus,
			Timestamp: timestamp,
			Amount:    amount,
		},
	}
	_, err := microClient.PutOne(context.TODO(), card)
	return err
}

// 交易pending
func SendTradeMsgPending(txnHash, toAddress, fromAddress, uid string, amount float64) error {
	logger.Info("send message is: txn=%s, to=%s, from=%s uid=%s ",
		txnHash, toAddress, fromAddress, uid, amount)
	err := SendTradeMsg(txnHash, "pending", toAddress, fromAddress, uid, amount)
	if err != nil {
		logger.Error(err)
	}
	return err
}

// 交易失败
func SendTradeMsgFailed(txnHash, toAddress, fromAddress, uid string, amount float64) error {
	err := SendTradeMsg(txnHash, "failed", toAddress, fromAddress, uid, amount)
	return err
}

// 交易成功
func SendTradeMsgSuccess(txnHash, toAddress, fromAddress, uid string, amount float64) error {
	err := SendTradeMsg(txnHash, "success", toAddress, fromAddress, uid, amount)
	return err
}

// 交易
func SendTradeMsg(txnHash, txnStatus, toAddress, fromAddress, uid string, amount float64) error {
	txn := &pbMessage.Msg{
		MsgType: 1,
		Txn: &pbMessage.Txn{
			TxnHash:     txnHash,
			TxnStatus:   txnStatus,
			Amount:      amount,
			Timestamp:   strconv.FormatInt(time.Now().Unix(), 10),
			ToAddress:   toAddress,
			FromAddress: fromAddress,
		},
		Uid: uid,
	}
	_, err := microClient.PutOne(context.TODO(), txn)
	return err
}

// 关注
func SendFollowMsg(uid, follower string) error {
	follow := &pbMessage.Msg{
		Uid:     uid,
		MsgType: 2,
		Follow: &pbMessage.Follow{
			Uid:       follower,
			Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		},
	}
	_, err := microClient.PutOne(context.TODO(), follow)
	return err
}

// 卡数量不足
func SendCardNotEnoughMsg(uid string) error {
	cardNotEnough := &pbMessage.Msg{
		Uid:     uid,
		MsgType: 7,
		OcardNotEnough: &pbMessage.OcardNotEnough{
			Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		},
	}
	_, err := microClient.PutOne(context.TODO(), cardNotEnough)
	return err
}

// 私信
func SendPrivateMsg(uid, content, MsgSenderUid string) error {
	privateMessage := &pbMessage.Msg{
		Uid:     uid,
		MsgType: 6,
		PrivateMessage: &pbMessage.PrivateMessage{
			Uid:       MsgSenderUid,
			Content:   content,
			Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		},
	}
	_, err := microClient.PutOne(context.TODO(), privateMessage)
	return err
}

// 评论
func SendCommentMsg(uid, oid, content, contentOid, CommentSenderUid string) error {
	comment := &pbMessage.Msg{
		Uid:     uid,
		MsgType: 5,
		Comment: &pbMessage.Comment{
			Uid:        CommentSenderUid,
			Oid:        oid,
			Content:    content,
			ContentOid: contentOid,
			Timestamp:  strconv.FormatInt(time.Now().Unix(), 10),
		},
	}
	_, err := microClient.PutOne(context.TODO(), comment)
	return err
}

// 转发
func SendForwordMsg(uid, ForworderUid, Oid string) error {
	foword := &pbMessage.Msg{
		Uid:     uid,
		MsgType: 4,
		Foword: &pbMessage.Foword{
			Uid:       ForworderUid,
			Oid:       Oid,
			Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		},
	}
	_, err := microClient.PutOne(context.TODO(), foword)
	return err
}

// 喜欢oop
func SendLikeMsg(uid, oid, liker string, likeNum int64, content string) error {
	like := &pbMessage.Msg{
		Uid:     uid,
		MsgType: 3,
		Like: &pbMessage.Like{
			Oid:       oid,
			Uid:       liker,
			LikeNum:   likeNum,
			Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
			Content: content,
		},
	}
	_, err := microClient.PutOne(context.TODO(), like)
	return err
}

//提及Msg
func SendMentionMsg(uid, mentioning, oid, Content, atUsers string) error {
	mention := &pbMessage.Msg{
		Uid:     uid,
		MsgType: 9,
		Mention: &pbMessage.Mention{
			Mentioning: mentioning,
			Timestamp:  strconv.FormatInt(time.Now().Unix(), 10),
			Uid: uid,
			Oid: oid,
			Content: Content,
			AtUsers: atUsers,
		},
	}
	_, err := microClient.PutOne(context.TODO(), mention)
	return err
}
