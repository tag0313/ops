package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
	"ops/pkg/db"
	"ops/pkg/model/consulreg"
	"ops/pkg/model/mgodb"
	"ops/pkg/utils"
	pbMessage "ops/proto/message"
	pbProperty "ops/proto/property"
	"strconv"
	"strings"
	"time"
)

func (p *Property) IncDecOpsPoint(ctx context.Context,
	req *pbProperty.IncDecOpsPointReq,
	resp *pbProperty.IncDecOpsPointResp) error {
	logger.Infof("op is %s, point is %f", req.Op.String(), req.OpsPoint)
	if req.GetOp() == pbProperty.OpsIncDec_Non {
		return fmt.Errorf("op field must set as %s or %s",
			pbProperty.OpsIncDec_Inc.String(), pbProperty.OpsIncDec_Dec.String())
	}
	if req.Uid == "" {
		return errors.New("uid cannot be null")
	}

	//find user acount
	mgoClient := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.user_opspoint"))
	record := mgoClient.FindOne(bson.M{"uid": req.Uid})
	if record.Err() != nil {
		logger.Error(record.Err())
		return fmt.Errorf("userID %s, record is not existed: %v", req.Uid, record.Err())
	}
	logger.Info("record %s", record)
	up := UserOpsPoint{}
	err := record.Decode(&up)
	if err != nil {
		logger.Error(err)
		return fmt.Errorf("record decoding failed %v", err)
	}
	logger.Infof("user ops_point is %+v", up)
	//if this is a withdraw operation, check the balance can be decreased
	if req.Op == pbProperty.OpsIncDec_Dec && up.OpsPoint-req.OpsPoint < 0 {
		return fmt.Errorf("account %s balance is not enough", req.Uid)
	}
	//operate account's balance
	if req.Op == pbProperty.OpsIncDec_Dec {
		_, err = mgoClient.UpsertRecordOne(ctx,
			bson.M{"uid": req.Uid}, bson.M{"$inc": bson.M{"ops_point": -req.OpsPoint}})
		resp.OpsBalance = up.OpsPoint - req.OpsPoint
		logger.Infof("computer balance(%f) - withdrawPoint(%f) = newBalance(%f)",
			up.OpsPoint, req.OpsPoint, resp.OpsBalance)
	} else {
		_, err = mgoClient.UpsertRecordOne(ctx,
			bson.M{"uid": req.Uid}, bson.M{"$inc": bson.M{"ops_point": req.OpsPoint}})
		resp.OpsBalance = up.OpsPoint + req.OpsPoint
		logger.Infof("computer balance(%f) + withdrawPoint(%f) = newBalance(%f)",
			up.OpsPoint, req.OpsPoint, resp.OpsBalance)
	}
	if err != nil {
		logger.Error(err)
		fmt.Errorf("update record failed %v", err)
		return err
	}

	return nil
}

func (p *Property) StoreTransferERC20History(ctx context.Context,
	request *pbProperty.TransferERC20Request,
	response *pbProperty.TransferERC20Response) error {
	defer func() {
		logger.Infof("calling StoreTransferERC20History success, request=%+v,  response=%+v", request, response)
	}()

	//save the transfer record into transfer_erc20_history
	mgoClient := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.transfer_erc20_history"))
	_, err := mgoClient.UpsertRecordOne(ctx, bson.M{"transaction_hash": request.TransactionHash},
		bson.M{"$set": bson.M{"transaction_hash": request.GetTransactionHash(),
			"from_address": request.FromAddress,
			"to_address":   request.ToAddress,
			"amount":       request.Amount},
		})
	if err != nil {
		logger.Error("upsert record failed", err)
		return fmt.Errorf("upsert record faild %v", err)
	}

	var fAmount64 float64
	fAmount64, err = strconv.ParseFloat(request.Amount, 32)
	if err != nil {
		logger.Error("string to float failed", err)
		return err
	}

	userInfo := &UserInfo{}
	mgoUserIdClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.user_id"))
	var record *mongo.SingleResult
	logger.Infof("to=%s,  manager=%s", request.ToAddress, utils.GetConfigStr("contract_addr.holder"))
	if strings.EqualFold(request.ToAddress, utils.GetConfigStr("contract_addr.holder")) {
		logger.Infof("charge the token, amount=%s", request.Amount)
		//充值
		record = mgoUserIdClient.FindOne(bson.M{"pubkaddr": strings.ToLower(request.FromAddress)})
		if record.Err() != nil {
			logger.Errorf("cant find the user info in the database. fromAddress=%s, err=%v",
				request.FromAddress, record.Err())
			return fmt.Errorf("find userID by fromAddress failed: %v", record.Err())
		}
		mgoOpsPointClient := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.user_opspoint"))
		err = record.Decode(&userInfo)
		if err != nil {
			logger.Info("the data decode to bytes by using mgo method is failed", err)
			return err

		}
		//充值加钱，提现不扣钱，因为提现在区块链转账之前钱已经扣掉了。
		logger.Infof("user=%+v", userInfo)
		_, err = mgoOpsPointClient.UpsertRecordOne(ctx,
			bson.M{"uid": userInfo.Uid}, bson.M{"$inc": bson.M{"ops_point": fAmount64}})
		if err != nil {
			logger.Info("add the opspoint for user is failed", userInfo.Uid, err)
			return err
		}
	} else if strings.EqualFold(request.FromAddress, utils.GetConfigStr("contract_addr.holder")) {
		//提现
		record = mgoUserIdClient.FindOne(bson.M{"pubkaddr": strings.ToLower(request.ToAddress)})
		//这里有可能拿不到用户ID，拿不到先不管，message 模块已经存了交易 hash 和 uid 关联
		if record.Err() != nil {
			logger.Errorf("cant find the user info in the database. fromAddress=%s, err=%v",
				request.FromAddress, record.Err())
		} else {
			err = record.Decode(&userInfo)
			if err != nil {
				logger.Error("the data decode to bytes by using mgo method is failed", err)
			}
		}
	} else {
		//用户到用户之间转账先不管，这是属于他们的私人间转账，我们只存交易记录。
		logger.Info("the transfer is user by user")
		return nil
	}

	//sending message to redis by using message service
	txn := &pbMessage.Msg{
		MsgType: 1,
		Uid:     userInfo.Uid,
		Txn: &pbMessage.Txn{
			TxnHash:     request.TransactionHash,
			TxnStatus:   "success",
			Amount:      fAmount64,
			Timestamp:   strconv.FormatInt(time.Now().Unix(), 10),
			ToAddress:   request.GetToAddress(),
			FromAddress: request.GetFromAddress(),
		},
	}
	messageClient := pbMessage.NewOperateMessageService("message", consulreg.MicroSer.Client())
	_, er := messageClient.PutOne(context.TODO(), txn)
	if er != nil {
		logger.Error(er)
	}

	return nil
}

func (p *Property) StoreWithdrawHistory(ctx context.Context,
	value *pbProperty.WithdrawOpspoint,
	result *pbProperty.OperateResult) error {
	defer func() {
		logger.Infof("calling OperateOpsPoint success,  card=%+v,  result=%+v", value, result)
	}()
	mgoClient := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.transfer_erc20_history"))
	err := mgoClient.UpsertOne(bson.M{"uid": value.Uid},
		bson.M{"$set": bson.M{"transaction_hash": value.TransactionHash, "ops_point": value.Opspoint, "gasfee": value.Gasfee}})
	if err == utils.RECODE_STOREDATA_FAILED {
		result.Code = err
	}
	result.Code = utils.RECODE_OK
	return nil
}

func (p *Property) CheckOCardAmountOps(ctx context.Context, ops *pbProperty.OCardsOnOps, result *pbProperty.OCardsOnOps) error {
	defer func() {
		logger.Infof("calling CheckOCardAmountOps success,  card=%+v,  result=%+v", ops, result)
	}()
	mgoClient := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.buyer_info_ops"))
	one := mgoClient.FindOne(bson.M{"group_id": ops.GroupId, "buyer_id": ops.BuyerUid})
	if one.Err() != nil {
		logger.Errorf("cant find the ocard info, card=%+v", ops.GroupId)
		return fmt.Errorf("cant find the ocard info, card=%+v", ops.GroupId)
	}
	OCardOnOps := new(BuyerInfoOps)
	err := one.Decode(&OCardOnOps)
	if err != nil {
		logger.Errorf("decode problem, card=%+v", ops.GroupId)
		return fmt.Errorf("decode problem card=%+v", ops.GroupId)
	}
	marshal, err := json.Marshal(OCardOnOps)
	if err != nil {
		logger.Errorf("decode problem, card=%+v", ops.GroupId)
		return fmt.Errorf("decode problem card=%+v", ops.GroupId)
	}
	err = json.Unmarshal(marshal, result)
	if err != nil {
		logger.Errorf("decode problem, card=%+v", ops.GroupId)
		return fmt.Errorf("decode problem card=%+v", ops.GroupId)
	}
	return nil
}

func (p *Property) OperateOCardAmountOps(ctx context.Context, ops *pbProperty.OCardsOnOps, result *pbProperty.OperateResult) error {
	defer func() {
		logger.Infof("calling OperateOCardAmountOps success,  card=%+v,  result=%+v", ops, result)
	}()
	mgoClient := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.buyer_info_ops"))
	update := mgoClient.UpsertOne(bson.M{"group_id": ops.GroupId, "buyer_id": ops.BuyerUid}, bson.M{"$inc": bson.M{"amount": ops.Amount}})
	if update != utils.RECODE_OK {
		logger.Errorf("cant find the ocard info, card=%+v", ops.GroupId)
		result.Code = update
		return nil
	}
	result.Code = utils.RECODE_OK
	return nil
}

func (p *Property) OperateOpsPoint(ctx context.Context, point *pbProperty.OpsPoint, result *pbProperty.OperateResult) error {
	defer func() {
		logger.Infof("calling OperateOpsPoint success,  card=%+v,  result=%+v", point, result)
	}()
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.property"), utils.GetConfigStr("mongodb.collection.user_opspoint"))
	err := mgoClient.UpsertOne(bson.M{"uid": point.Uid}, bson.M{"$inc": bson.M{"ops_point": point.OpsPoint}})
	if err == utils.RECODE_STOREDATA_FAILED {
		result.Code = err
	}
	result.Code = utils.RECODE_OK
	return nil
}

func (p *Property) CheckOpsPoint(ctx context.Context, point *pbProperty.OpsPoint, result *pbProperty.OpsPoint) error {
	defer func() {
		logger.Infof("calling CheckOpsPoint , point=%+v,  result=%+v", point, result)
	}()
	mgoClient := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.user_opspoint"))
	one := mgoClient.FindOne(bson.M{"uid": point.Uid})
	if one.Err() != nil {
		logger.Error(one.Err())
		result.Uid = point.Uid
		result.OpsPoint = 0
		return nil
	}
	bytes, err := one.DecodeBytes()
	if err != nil {
		logger.Error(err)
		return nil
	}
	result.Uid = bytes.Lookup("uid").StringValue()
	result.OpsPoint = bytes.Lookup("ops_point").Double()
	return nil
}

func (p *Property) StoreMintOCardInfo(ctx context.Context, card *pbProperty.OCard, result *pbProperty.OperateResult) error {
	defer func() {
		logger.Infof("calling StoreMintOCardInfo,  card=%+v,  result=%+v", card, result)
	}()
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.property"), utils.GetConfigStr("mongodb.collection.minted_ocard"))
	if card.Uid != "" && card.TransactionHash != "" {
		one := mgoClient.UpsertOne(
			bson.M{"uid": card.Uid},
			bson.M{"$set": bson.M{
				"amount":           intBytesToString(card.Amount),
				"card_type":        card.CardType,
				"transaction_hash": card.TransactionHash}})
		if one == utils.RECODE_STOREDATA_FAILED {
			result.Code = utils.RECODE_STOREDATA_FAILED
			return nil
		}
		result.Code = utils.RECODE_OK
		return nil
	} else if card.TransactionHash != "" {
		one := mgoClient.UpsertOne(
			bson.M{"transaction_hash": card.TransactionHash},
			bson.M{"$set": bson.M{
				"amount":    intBytesToString(card.Amount),
				"group_id":  card.GroupId,
				"mint_date": card.MintDate}})
		if one == utils.RECODE_STOREDATA_FAILED {
			result.Code = utils.RECODE_STOREDATA_FAILED
			return nil
		}
		result.Code = utils.RECODE_OK
		return nil
	}
	return nil
}

func (p *Property) StoreMintOCardsInfo(ctx context.Context, cards *pbProperty.OCards, result *pbProperty.OperateResult) error {
	defer func() {
		logger.Infof("calling StoreMintOCardsInfo,  card=%+v,  result=%+v", cards, result)
	}()
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.property"), utils.GetConfigStr("mongodb.collection.minted_ocard"))
	for _, value := range cards.Ocard {
		if value.JsonId != "" && value.TransactionHash != "" {
			one := mgoClient.UpsertOne(
				bson.M{"json_id": value.JsonId},
				bson.M{"$set": bson.M{
					"uid":              value.Uid,
					"amount":           intBytesToInt(value.Amount),
					"card_type":        value.CardType,
					"transaction_hash": value.TransactionHash,
					"mint_date":        getRegisterTime()}})
			if one == utils.RECODE_STOREDATA_FAILED {
				result.Code = utils.RECODE_STOREDATA_FAILED
				return nil
			}
		}
	}
	result.Code = utils.RECODE_OK
	return nil
}

func (p *Property) StoreMintOCardsSuccessInfo(ctx context.Context, info *pbProperty.MintOCardsSuccessInfo,
	result *pbProperty.OperateResult) (err error) {
	defer func() {
		logger.Infof("calling StoreMintOCardsSuccessInfo,  card=%+v,  result=%+v, err=%v", info, result, err)
		result.Code = utils.RECODE_OK
	}()
	managerHolderAddr := common.HexToAddress(utils.GetConfigStr("contract_addr.holder"))
	fromAddr := common.HexToAddress(info.FromAddress)
	toAddr := common.HexToAddress(info.ToAddress)
	logger.Infof("from=%s, to=%s, manager=%s", fromAddr.String(), toAddr.String(), managerHolderAddr.String())

	t := db.NewOpTable()
	var hs []interface{ db.Record }
	for i, gid := range info.GroupIds {
		hs = append(hs, &WithdrawOcardHistory{
			GroupID:         bytes2bigInt(gid).String(),
			Amount:          bytes2bigInt(info.Amounts[i]).Uint64(),
			TransactionHash: info.TransactionHash,
			FromAddress:     fromAddr.String(),
			ToAddress:       toAddr.String(),
			OperatorAddress: info.OperatorAddress,
			TransactionTime: utils.NowUnix(),
		})
	}
	err = t.InsertMany(ctx, hs)
	if err != nil {
		logger.Error(err)
		return err
	}

	//from address is "0x0000000000000000000000000000000000000000"
	if fromAddr.Hash().Big().Cmp(big.NewInt(0)) == 0 {
		err = handleCreateBatch(info, result)
	} else if strings.EqualFold(fromAddr.String(), managerHolderAddr.String()) {
		//从服务器到用户钱包，导出 OCard
		err = handleTransferBatchID(info, result)
		//store the chain information
		err = storeOCardChain(info)
		if err != nil {
			logger.Error()
			return err
		}
	} else if strings.EqualFold(toAddr.String(), managerHolderAddr.String()) {
		//从用户钱包到服务器 导入 Ocard
		err = handleBatchTransferFrom(info, result)
	} else {
		//err = HandleUserTransaction(info, result)
	}

	return err
}

func (p *Property) StoreOCardFromChain(ctx context.Context, card *pbProperty.OCardsOnOps, result *pbProperty.OperateResult) error {
	defer func() {
		logger.Infof("calling StoreOCardFromChain,  card=%+v,  result=%+v, err=%v", card, result)
		result.Code = utils.RECODE_OK
	}()

	mgoClient := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.minted_ocard"))
	one := mgoClient.FindOne(bson.M{"group_id": card.GroupId})
	if one.Err() != nil {
		logger.Error(one.Err())
		return one.Err()
	}
	ops := MintedOCard{}
	err := one.Decode(&ops)
	if err != nil {
		logger.Error(err)
		return err
	}
	mgoChainClient := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.buyer_info_chain"))
	mgoChainClient.UpsertOne(bson.M{"buyer_id": card.BuyerUid, "group_id": card.GroupId},
		bson.M{"$set": bson.M{"seller_uid": ops.Uid,
			"card_type":     ops.CardType,
			"mint_date":     ops.MintDate,
			"purchase_time": getRegisterTime()}})
	mgoChainClient.UpsertOne(bson.M{"buyer_id": card.BuyerUid, "group_id": card.GroupId},
		bson.M{"$inc": bson.M{"amount": card.Amount}})

	//delete ocard info in buyer_info_ops
	mgoOpsClient := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.buyer_info_ops"))
	deleteOne := mgoOpsClient.Delete(bson.M{"buyer_id": card.BuyerUid, "group_id": card.GroupId, "amount": bson.M{"$eq": 0}})
	if deleteOne != utils.RECODE_OK {
		return fmt.Errorf("delete ocard info in buyer_info_ops failed, err=%+v", deleteOne)
	}
	result.Code = utils.RECODE_OK
	return nil
}

func (p *Property) BuyOCardOnOps(ctx context.Context, req *pbProperty.BuyOCardOnOpsReq, result *pbProperty.OperateResult) error {
	defer func() {
		logger.Infof("calling BuyOCardOnOps,  card=%+v,  result=%+v", req, result)
		logger.Info("execute BuyOCardOnOps success")
	}()
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.property"), utils.GetConfigStr("mongodb.collection.buyer_info_ops"))
	for _, value := range req.PurchaseInfo {
		info := mgoClient.UpsertOne(bson.M{"group_id": value.GroupId, "buyer_id": value.BuyerUid},
			bson.M{"$set": bson.M{
				"card_type":     value.CardType,
				"mint_date":     value.MintDate,
				"seller_uid":    value.SellerUid,
				"purchase_time": value.PurchaseTime,
			},
			})
		if info == utils.RECODE_STOREDATA_FAILED {
			result.Code = utils.RECODE_STOREDATA_FAILED
			return nil
		}
		amount := mgoClient.UpsertOne(bson.M{"group_id": value.GroupId, "buyer_id": value.BuyerUid},
			bson.M{"$inc": bson.M{"amount": value.Amount}})
		if amount == utils.RECODE_STOREDATA_FAILED {
			result.Code = utils.RECODE_STOREDATA_FAILED
			return nil
		}
	}
	result.Code = utils.RECODE_OK
	return nil
}

func (p *Property) OperateOCardAmount(ctx context.Context, card *pbProperty.OCard, result *pbProperty.OperateResult) error {
	defer func() {
		logger.Infof("calling OperateOCardAmount,  card=%+v,  result=%+v", card, result)
		logger.Info("execute OperateOCardAmount success")
	}()
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.property"), utils.GetConfigStr("mongodb.collection.minted_ocard"))
	update := mgoClient.FindOneAndUpdate(bson.M{"group_id": card.GroupId}, bson.M{"$inc": bson.M{"sold": intBytesToInt(card.Amount)}})
	if update.Err() != nil {
		logger.Error(update.Err())
		result.Code = utils.RECODE_STOREDATA_FAILED
		return nil
	}
	bytes, err := update.DecodeBytes()
	if err != nil {
		logger.Error(err)
		return nil
	}
	amount := bytes.Lookup("amount").Int64()
	sold := bytes.Lookup("sold").Int64()
	if amount-sold < 0 {
		result.Code = utils.RECODE_INSUFFICIENT_FUND
		return nil
	}
	result.Code = utils.RECODE_OK
	return nil
}

func (p *Property) QueryMintedOCard(ctx context.Context, field *pbProperty.QueryField, card *pbProperty.ListOCardOnMongo) error {
	defer func() {
		logger.Infof("calling QueryMintedOCard,  card=%+v,  result=%+v", field, card)
		logger.Info("execute QueryMintedOCard success")
	}()
	var (
		oCardOnMongo []*MintedOCard
		listOCard    []*pbProperty.OCardOnMongo
	)
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.property"), utils.GetConfigStr("mongodb.collection.minted_ocard"))
	optionsFind := options.Find().SetProjection(bson.M{"_id": 0, "json_id": 0, "transaction_hash": 0}).SetSort(bson.M{"mint_date": -1})
	many, errCode := mgoClient.FindMany(bson.M{"uid": field.Uid}, optionsFind)
	if errCode == utils.RECODE_DATAINEXISTENCE {
		card.Code = errCode
		return nil
	}

	err := many.All(context.TODO(), &oCardOnMongo)
	if err != nil {
		logger.Error("binding data failed", err)
		card.Code = utils.RECODE_DATAERR
		return nil
	}
	marshal, err := json.Marshal(&oCardOnMongo)
	if err != nil {
		logger.Error("binding data failed", err)
		card.Code = utils.RECODE_DATAERR
		return nil
	}
	err = json.Unmarshal(marshal, &listOCard)
	if err != nil {
		logger.Error("binding data failed", err)
		card.Code = utils.RECODE_DATAERR
		return nil
	}
	card.Code = utils.RECODE_OK
	card.Data = listOCard
	return nil
}

func (p *Property) QueryOCardFromLocal(ctx context.Context, field *pbProperty.QueryField, card *pbProperty.ListOCardOnOps) error {
	defer func() {
		logger.Infof("calling QueryOCardFromLocal,  card=%+v,  result=%+v", field, card)
		logger.Info("execute QueryOCardFromLocal success")
	}()
	var (
		oCardOnMongo []*BuyerInfoOps
		listOCard    []*pbProperty.OCardsOnOps
	)
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.property"), utils.GetConfigStr("mongodb.collection.buyer_info_ops"))
	optionsFind := options.Find().SetProjection(bson.M{"_id": 0, "json_id": 0, "transaction_hash": 0}).SetSort(bson.M{"purchase_time": -1})
	many, errCode := mgoClient.FindMany(bson.M{"buyer_id": field.Uid}, optionsFind)
	if errCode == utils.RECODE_DATAINEXISTENCE {
		card.Code = errCode
		return nil
	}

	err := many.All(context.TODO(), &oCardOnMongo)
	if err != nil {
		logger.Error("binding data failed", err)
		card.Code = utils.RECODE_DATAERR
		return nil
	}
	marshal, err := json.Marshal(&oCardOnMongo)
	if err != nil {
		logger.Error("binding data failed", err)
		card.Code = utils.RECODE_DATAERR
		return nil
	}
	err = json.Unmarshal(marshal, &listOCard)
	if err != nil {
		logger.Error("binding data failed", err)
		card.Code = utils.RECODE_DATAERR
		return nil
	}
	card.Code = utils.RECODE_OK
	card.PurchaseInfo = listOCard
	return nil
}

func (p *Property) QueryOCardFromChain(ctx context.Context, field *pbProperty.QueryField, card *pbProperty.ListOCardOnOps) error {
	defer func() {
		logger.Infof("calling QueryOCardFromLocal,  card=%+v,  result=%+v", field, card)
		logger.Info("execute QueryOCardFromLocal success")
	}()
	var (
		oCardOnMongo []*BuyerInfoOps
		listOCard    []*pbProperty.OCardsOnOps
	)
	mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.property"), utils.GetConfigStr("mongodb.collection.buyer_info_chain"))
	optionsFind := options.Find().SetProjection(bson.M{"_id": 0, "json_id": 0, "transaction_hash": 0}).SetSort(bson.M{"purchase_time": -1})
	many, errCode := mgoClient.FindMany(bson.M{"buyer_id": field.Uid}, optionsFind)
	if errCode == utils.RECODE_DATAINEXISTENCE {
		card.Code = errCode
		return nil
	}

	err := many.All(context.TODO(), &oCardOnMongo)
	if err != nil {
		logger.Error("binding data failed", err)
		card.Code = utils.RECODE_DATAERR
		return nil
	}
	marshal, err := json.Marshal(&oCardOnMongo)
	if err != nil {
		logger.Error("binding data failed", err)
		card.Code = utils.RECODE_DATAERR
		return nil
	}
	err = json.Unmarshal(marshal, &listOCard)
	if err != nil {
		logger.Error("binding data failed", err)
		card.Code = utils.RECODE_DATAERR
		return nil
	}
	card.Code = utils.RECODE_OK
	card.PurchaseInfo = listOCard
	return nil
}

func (p *Property) RelationshipJsonIdAndGroupId(ctx context.Context, card *pbProperty.JsonIdAndGroupId, result *pbProperty.OperateResult) error {
	mgodbClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.property"), utils.GetConfigStr("mongodb.collection.relationship"))
	one := mgodbClient.UpsertOne(bson.M{"json_id": card.JsonId}, bson.M{"$set": bson.M{"json_id": card.JsonId, "group_id": card.GroupId}})
	if one == utils.RECODE_STOREDATA_FAILED {
		result.Code = utils.RECODE_STOREDATA_FAILED
		return nil
	}
	result.Code = utils.RECODE_OK
	return nil
}

func (p *Property) RelationshipJsonIdAndGroupIds(ctx context.Context, ids *pbProperty.JsonIdAndGroupIds, result *pbProperty.OperateResult) error {
	mgodbClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.property"), utils.GetConfigStr("mongodb.collection.relationship"))
	for _, value := range ids.JsonIdAndGroupId {
		one := mgodbClient.UpsertOne(bson.M{"json_id": value.JsonId}, bson.M{"$set": bson.M{"json_id": value.JsonId, "group_id": value.GroupId}})
		if one == utils.RECODE_STOREDATA_FAILED {
			result.Code = utils.RECODE_STOREDATA_FAILED
			return nil
		}
	}
	result.Code = utils.RECODE_OK
	return nil
}
