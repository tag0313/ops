package handler

import (
	pbProperty "ops/proto/property"
	"testing"
)

type prepareHandleUserTransaction struct{

}

func (p *prepareHandleUserTransaction) generateData(info *pbProperty.MintOCardsSuccessInfo){
	//t := db.NewOpTable()
	//fromUser := db.UserID{
	//	Uid:    utils.RandomStrLen(10),
	//	PubKey: info.FromAddress,
	//}
	//
	//err := t.InsertOne(context.TODO(), fromUser)
	//if err != nil{
	//	logger.Fatal(err)
	//}
	//
	//toUser := db.UserID{
	//	Uid:    utils.RandomStrLen(10),
	//	PubKey: info.ToAddress,
	//}
	//err = t.InsertOne(context.TODO(), toUser)
	//if err != nil{
	//	logger.Fatal(err)
	//}
	//
	//for _, groupID := range info.GroupIds {
	//	gid := bytes2bigInt(groupID).String()
	//	mo := db.MintedOCard{
	//		GroupId:         gid,
	//		CardType:        "",
	//		MintDate:        "",
	//		Uid:             utils.RandomStrLen(10),
	//		Amount:          1,
	//		TransactionHash: "",
	//		JsonId:          "0",
	//		Sold:            0,
	//	}
	//	err = t.InsertOne(context.TODO(), mo)
	//	if err != nil{
	//		logger.Fatal(err)
	//	}
	//
	//	toBi := db.BuyerInfoOps{
	//		BuyerUid:     toUser.Uid,
	//		GroupId:      gid,
	//		CardType:     "",
	//		MintDate:     "",
	//		SellerUid:    "",
	//		Amount:       1,
	//		PurchaseTime: "",
	//		UnitPrice:    0,
	//	}
	//	err = t.InsertOne(context.TODO(), toBi)
	//	if err != nil{
	//		logger.Fatal(err)
	//	}
	//
	//}
}
func TestHandleUserTransaction(t *testing.T) {
	//WTF....
}