package test

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"math/big"
	"ops/pkg/model/mgodb"
	"ops/pkg/model/rdb"
	utils2 "ops/pkg/utils"
	"testing"
)

func TestHash(t *testing.T) {
	t.Run("rdb", func(t *testing.T) {
		mgoClient := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db"), utils2.GetConfigStr("mongodb.collection.user_detail"))
		many, _ := mgoClient.FindMany(bson.M{"is_private": true})
		for many.Next(context.Background()) {
			fmt.Println(many.Current.Lookup("uid").StringValue())

			rdb.LPush("test_1", many.Current.Lookup("uid").StringValue())
			rdb.LPush("test_1", many.Current.Lookup("uid").StringValue())
		}
		rdb.LRem("test_1", 1, "yArSC5HmIMFjYbeFdAil")
	})

	t.Run("queryUser", func(t *testing.T) {
		//userInfo := mgodb.NewMgo(utils2.GetConfigStr("mongodb.db"), utils2.GetConfigStr("mongodb.collection.user_detail"))
		//andDelete := pbUserInfo.QueryAndDelete{Uid: "yArSC5HmIMFjYbeFdAil1"}
		//singleResult := userInfo.FindOne(bson.M{"uid": andDelete.Uid})
		//fmt.Println(singleResult.Err() != nil)
		//fmt.Println(singleResult.DecodeBytes())

		donateAmount := big.NewRat(2, 100)
		fmt.Println(donateAmount)
	})
}
