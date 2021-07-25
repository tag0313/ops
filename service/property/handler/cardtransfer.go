package handler

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/logger"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"ops/pkg/model/consulreg"
	"ops/pkg/model/mgodb"
	"ops/pkg/utils"
	"ops/proto/follower"
	pbMessage "ops/proto/message"
	"ops/proto/nft1155"
	"ops/proto/property"
	"ops/proto/userInfo"
	"strconv"
	"strings"
	"time"
)

const (
	followOP   = true
	unfollowOP = false
)

func handleCreateBatch(info *pbProperty.MintOCardsSuccessInfo, result *pbProperty.OperateResult) error {
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.property"), utils.GetConfigStr("mongodb.collection.minted_ocard"))
	createCardsClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	var opss client.CallOption = func(o *client.CallOptions) {
		// o.RequestTimeout = time.Minute * 30
		o.RequestTimeout = time.Second * 10
	}
	var (
		totalAmount int64
		uid         string
		transaction string
	)
	for index, value := range info.GetGroupIds() {
		pbRequest := new(pbNft1155.UriRequest)
		fmt.Println(bytes2bigInt(value))
		pbRequest.Id = value
		resp, err := createCardsClient.Uri(context.TODO(), pbRequest, opss)
		if err != nil {
			logger.Error("through group_id to find the json_id from chain is failed ", err)
			result.Code = utils.RECODE_MICROERR
			return nil
		}
		jsonId := strings.Split(resp.Uri, "/")[4][:19]
		//jsonId := strings.Split(resp.Uri, ".")[0]
		one := mgoClient.FindOneAndUpdate(bson.M{"json_id": jsonId},
			bson.M{"$set": bson.M{"group_id": intBytesToString(value),
				"amount": intBytesToInt(info.Amounts[index])}})
		if one.Err() != nil {
			logger.Error("execute StoreMintOCardsSuccessInfo failed")
			result.Code = utils.RECODE_STOREDATA_FAILED
			return nil
		}
		bytes, err := one.DecodeBytes()
		if err != nil {
			return err
		}
		uid = bytes.Lookup("uid").StringValue()
		totalAmount += intBytesToInt(info.Amounts[index])

		//对接message服务
		transaction = bytes.Lookup("transaction_hash").StringValue()
	}

	fmt.Println("totalamount: ", totalAmount)
	//set the total amount for user
	req := &pbUserInfo.TotalReleaseOCardNum{}
	req.ReleaseCardNum = totalAmount
	req.Uid = uid
	setTotalClient := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
	num, err := setTotalClient.SetTotalReleaseOCardNum(context.TODO(), req)
	if err != nil {
		logger.Info("execute SetTotalReleaseOCardNum")
		result.Code = utils.RECODE_MICROERR
		return nil
	}

	// sent the notification to message service 交易生成 成功
	msg := &pbMessage.Msg{
		MsgType: 8,
		Uid:     uid,
		CreateCard: &pbMessage.CreateCard{
			TxnHash:   transaction,
			TxnStatus: "success",
			Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
			Amount:    totalAmount,
		},
	}
	messageClient := pbMessage.NewOperateMessageService("message", consulreg.MicroSer.Client())
	_, err = messageClient.PutOne(context.TODO(), msg)
	if err != nil {
		logger.Error(err, "消息进入redis队列失败")
	}
	result.Code = num.Code
	return nil
}

//handleTransferBatchID 从服务器到用户钱包，导出 OCard
func handleTransferBatchID(info *pbProperty.MintOCardsSuccessInfo, result *pbProperty.OperateResult) error {
	logger.Info("handleTransferBatchID.....")
	mgoUserIdClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.user_id"))
	toAddr := common.HexToAddress(info.ToAddress)
	filter := mgoFilterCaseInsensitive("pubkaddr", toAddr.String())
	findResult := mgoUserIdClient.FindOne(filter)
	ui := UserInfo{}
	if err := findResult.Decode(&ui); err != nil {
		logger.Error(err)
		return err
	}
	var amount int64
	for i := range info.Amounts {
		amount += bytes2bigInt(info.Amounts[i]).Int64()
	}
	//sending message to redis by using message service
	txn := &pbMessage.Msg{
		MsgType: 1,
		Uid:     ui.Uid,
		Txn: &pbMessage.Txn{
			TxnHash:     info.TransactionHash,
			TxnStatus:   "success",
			Amount:      float64(amount),
			Timestamp:   strconv.FormatInt(time.Now().Unix(), 10),
			ToAddress:   info.ToAddress,
			FromAddress: info.FromAddress,
		},
	}
	messageClient := pbMessage.NewOperateMessageService("message", consulreg.MicroSer.Client())
	_, er := messageClient.PutOne(context.TODO(), txn)
	if er != nil {
		logger.Error(er)
	}

	return nil
}

//handleBatchTransferFrom 从用户钱包到服务器， 导入 Ocard.
func handleBatchTransferFrom(info *pbProperty.MintOCardsSuccessInfo, result *pbProperty.OperateResult) error {
	logger.Info("handleBatchTransferFrom.....")
	fromAddr := common.HexToAddress(info.FromAddress)
	mgoUserIdClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.user_id"))

	filter := mgoFilterCaseInsensitive("pubkaddr", fromAddr.String())
	findResult := mgoUserIdClient.FindOne(filter)
	if findResult.Err() != nil {
		logger.Error(findResult.Err())
		return findResult.Err()
	}
	ui := UserInfo{}
	if err := findResult.Decode(&ui); err != nil {
		logger.Error(err)
		return err
	}
	logger.Infof("userInfo is %+v", ui)

	err := storeBuyerInfo(ui.Uid, info)
	if err != nil {
		logger.Error(err)
		return err
	}

	var amount int64
	for i := range info.Amounts {
		amount += bytes2bigInt(info.Amounts[i]).Int64()
	}
	//sending message to redis by using message service
	txn := &pbMessage.Msg{
		MsgType: 1,
		Uid:     ui.Uid,
		Txn: &pbMessage.Txn{
			TxnHash:     info.TransactionHash,
			TxnStatus:   "success",
			Amount:      float64(amount),
			Timestamp:   strconv.FormatInt(time.Now().Unix(), 10),
			ToAddress:   info.ToAddress,
			FromAddress: info.FromAddress,
		},
	}
	messageClient := pbMessage.NewOperateMessageService("message", consulreg.MicroSer.Client())
	_, er := messageClient.PutOne(context.TODO(), txn)
	if er != nil {
		logger.Error(er)
	}

	return nil
}

func HandleUserTransaction(info *pbProperty.MintOCardsSuccessInfo,
	result *pbProperty.OperateResult) (err error) {
	defer func() {
		logger.Info("handle user transaction info=%+v, result=%+v, err=%v",
			info, result, err)
	}()
	fromAddr := common.HexToAddress(info.FromAddress)
	toAddr := common.HexToAddress(info.ToAddress)

	var ocs []MintedOCard
	ocs, err = fetchOcardsCreator(info.Amounts)
	if err != nil {
		logger.Error(err)
		return err
	}

	mgoUserIdClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.user_id"))
	filter := mgoFilterCaseInsensitive("pubkaddr", fromAddr.String())
	findResult := mgoUserIdClient.FindOne(filter)
	if findResult.Err() != mongo.ErrNoDocuments {
		fromUser := UserInfo{}
		if err = findResult.Decode(&fromUser); err != nil {
			logger.Error(err)
		} else {
			for _, oc := range ocs {
				hasCard, err := checkUserHasCard(fromUser.Uid, []string{oc.GroupId})
				if err != nil {
					logger.Error(err)
				} else if hasCard == false {
					err = operateFollowing(fromUser.Uid, oc.Uid, unfollowOP)
				}
			}
		}
	}

	filter = mgoFilterCaseInsensitive("pubkaddr", toAddr.String())
	findResult = mgoUserIdClient.FindOne(filter)
	if findResult.Err() != mongo.ErrNoDocuments {
		toUser := UserInfo{}
		if err = findResult.Decode(&toUser); err != nil {
			logger.Error(err)
		} else {
			for _, oc := range ocs {
				hasCard, err := checkUserHasCard(toUser.Uid, []string{oc.GroupId})
				if err != nil {
					logger.Error(err)
				} else if hasCard == true {
					err = operateFollowing(toUser.Uid, oc.Uid, followOP)
				}
			}
		}
	}

	return err
}

func operateFollowing(userID, followingID string,
	followingOrUnfollowing bool) (err error) {
	srv := consulreg.MicroSer
	f := pbFollower.NewOperateFollowService("follower", srv.Client())
	data := &pbFollower.Follower{
		Uid:       userID,
		Following: followingID,
	}
	if followingOrUnfollowing { //following
		_, err = f.Follow(context.TODO(), data)
		return err
	} else { //unfollowing
		_, err = f.CanalFollow(context.TODO(), data)
		return err
	}
	return nil
}
