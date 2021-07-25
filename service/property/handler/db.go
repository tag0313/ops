package handler

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"go.mongodb.org/mongo-driver/bson"
	"ops/pkg/utils"

	//"go.mongodb.org/mongo-driver/mongo/options"
	"ops/pkg/model/mgodb"
	pbProperty "ops/proto/property"
)

type incDecOP int

const (
	OCardInc incDecOP = 1
	OCardDec incDecOP = 2
)

func opOcardChain(GroupId string, Amount uint64, OwnerID string, CreatorID string,
	CardType string, op incDecOP) error {

	return nil
}

func storeOCardChain(info *pbProperty.MintOCardsSuccessInfo) error {
	userInfoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.user_id"))
	filter := mgoFilterCaseInsensitive("pubkaddr", info.ToAddress)
	findResult := userInfoClient.FindOne(filter)
	ui := UserInfo{}
	if err := findResult.Decode(&ui); err != nil {
		logger.Error(err)
		return err
	}

	mgoClient := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.withdraw_ocard_history"))
	chainInfo := BuyerInfoChain{
		//TransactionHash: info.TransactionHash,
		//FromAddress:     info.FromAddress,
		//ToAddress:       info.ToAddress,
		//OperatorAddress: info.ToAddress,
		OwnerID: ui.Uid,
	}

	var cs []interface{}
	for i := range info.GroupIds {
		c := chainInfo
		id := utils.Bytes2bigInt(info.GroupIds[i])
		amount := utils.Bytes2bigInt(info.Amounts[i])
		c.GroupId = id.String()
		c.Amount = amount.Uint64()
		cs = append(cs, c)
	}
	_, err := mgoClient.InsertMany(cs)
	if err != nil {
		return err
	}
	return nil
}

func storeBuyerInfo(uid string, info *pbProperty.MintOCardsSuccessInfo) error {
	mintedOCardDB := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.minted_ocard"))
	var groupIDs []string
	for i := range info.GroupIds {
		groupIDs = append(groupIDs, bytes2bigInt(info.GroupIds[i]).String())
	}

	filter := bson.M{"group_id": bson.M{"$in": groupIDs}}
	manyRecords, err := mintedOCardDB.FindManyRecords(filter, nil)
	if err != nil {
		logger.Error(err)
		return err
	}
	var ocs []MintedOCard
	err = manyRecords.All(context.TODO(), &ocs)
	if err != nil {
		return err
	}

	//upsert buyer info
	buyerInfoDB := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.buyer_info_ops"))
	//upsert buyer chain
	buyerChain := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.buyer_info_chain"))
	//var buyerInfos []interface{}
	for i := range info.GroupIds {
		var oc MintedOCard
		id := bytes2bigInt(info.GroupIds[i]).String()
		for j := range ocs {
			if ocs[j].GroupId == id {
				oc = ocs[i]
				_, err := buyerInfoDB.UpsertRecordOne(context.TODO(), bson.M{"group_id": id, "buyer_id": uid},
					bson.M{"$set": bson.M{
						"card_type":     oc.CardType,
						"mint_date":     oc.MintDate,
						"seller_uid":    oc.Uid,
						"purchase_time": getRegisterTime(),
					},
					})
				if err != nil {
					logger.Error(err)
					return err
				}
				_, err = buyerInfoDB.UpsertRecordOne(context.TODO(), bson.M{"group_id": id, "buyer_id": uid},
					bson.M{"$inc": bson.M{"amount": bytes2bigInt(info.Amounts[i]).Uint64()}})
				if err != nil {
					logger.Error(err)
					return err
				}
				//delete data from buyer chain
				one := buyerChain.FindOne(bson.M{"group_id": id, "buyer_id": uid})
				if one.Err() != nil {
					logger.Error(one.Err())
				}
				bytes, err := one.DecodeBytes()
				if err != nil {
					logger.Error(err)
				}
				amountOnMgo := uint64(bytes.Lookup("amount").Int64())
				amountOnCurrent := bytes2bigInt(info.Amounts[i]).Uint64()
				if amountOnMgo-amountOnCurrent == 0 {
					buyerChain.Delete(bson.M{"group_id": id, "buyer_id": uid})
				} else {
					buyerChain.UpsertRecordOne(context.TODO(), bson.M{"group_id": id, "buyer_id": uid},
						bson.M{"$inc": bson.M{"amount": -amountOnCurrent}})
				}
				break
			}
		}

		//buyerInfos = append(buyerInfos, BuyerInfoOps{
		//	GroupId:      id,
		//	BuyerUid:     uid,
		//	CardType:     oc.CardType,
		//	MintDate:     oc.MintDate,
		//	PurchaseTime: "",
		//	SellerUid:    oc.Uid,
		//	Amount:       bytes2bigInt(info.Amounts[i]).Uint64(),
		//})
	}

	//_, err = buyerInfoDB.InsertMany(buyerInfos)
	//if err != nil {
	//	return err
	//}
	return nil
}

func fetchOcardsCreator(ids [][]byte) (ocs []MintedOCard, err error) {
	mintedOCardDB := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.minted_ocard"))
	var groupIDs []string
	for i := range ids {
		groupIDs = append(groupIDs, bytes2bigInt(ids[i]).String())
	}

	filter := bson.M{"group_id": bson.M{"$in": groupIDs}}
	manyRecords, err := mintedOCardDB.FindManyRecords(filter, nil)
	if err != nil {
		logger.Error(err)
		return
	}
	//先找到所有得 creators
	var creators []MintedOCard
	err = manyRecords.All(context.TODO(), &creators)
	if err != nil {
		return
	}
	var uids []string
	for _, creator := range creators {
		uids = append(uids, creator.Uid)
	}
	filter = bson.M{"uid": bson.M{"$in": uids}}
	rs, err := mintedOCardDB.FindManyRecords(filter, nil)
	if err != nil {
		return
	}
	err = rs.All(context.TODO(), &ocs)
	if err != nil {
		return
	}

	return
}

func checkUserHasCard(buyerID string, groupIDs []string) (has bool, err error) {
	mintedOCardDB := mgodb.NewMgo(propertyDB, utils.GetConfigStr("mongodb.collection.buyer_info_ops"))
	filter := bson.M{"group_id": bson.M{"$in": groupIDs}, "buyer_id": buyerID}
	items, err := mintedOCardDB.CountDocument(filter)
	if err != nil {
		logger.Error(err)
		return
	}
	if items == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
