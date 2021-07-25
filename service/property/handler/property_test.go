package handler

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
	mgodb2 "ops/pkg/model/mgodb"
	utils2 "ops/pkg/utils"
	pbProperty "ops/proto/property"
	"testing"
)

func TestStoreMintOCardsSuccessInfo(t *testing.T) {
	testAddress := "0xabcdeec7a9C15643411A2584438E9D085d4c54ee"
	p := &Property{}
	info := new(pbProperty.MintOCardsSuccessInfo)
	info.FromAddress = utils2.GetConfigStr("contract_addr.holder")
	info.ToAddress = testAddress
	for i := 0; i < 10; i++ {
		info.Amounts = append(info.Amounts, new(big.Int).SetUint64(22).Bytes())
		info.GroupIds = append(info.GroupIds, new(big.Int).SetUint64(uint64(i)).Bytes())
	}
	result := new(pbProperty.OperateResult)
	err := p.StoreMintOCardsSuccessInfo(context.Background(), info, result)
	if err == nil {
		t.Fatal(err)
		return
	}

}

func TestGetInfoFromMintedCard(t *testing.T) {
	mgoClient := mgodb2.NewMgo(utils2.GetConfigStr("mongodb.db.property"), utils2.GetConfigStr("mongodb.collection.minted_ocard"))
	update := mgoClient.FindOne(bson.M{"group_id": "1298"})
	if update.Err() != nil {
		logger.Error(update.Err())
	}
	bytes, err := update.DecodeBytes()
	if err != nil {
		logger.Error(err)
	}
	amount := bytes.Lookup("amount").Int64()
	sold := bytes.Lookup("sold").Int64()
	if amount-sold < 0 {
		fmt.Println("success")
	}
}
