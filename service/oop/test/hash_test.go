package test

import (
	"context"
	"fmt"
	"ops/pkg/model/mgodb"
	"ops/pkg/model/rdb"
	"ops/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"

	"testing"
)

func TestHash(t *testing.T) {
	t.Run("skip", func(t *testing.T) {
		get := rdb.GetList("private_user", 0, -1)
		//test := make([]string, 0)
		mgoClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db"), utils.GetConfigStr("mongodb.collection.oop"))
		many, _ := mgoClient.FindMany(bson.M{"is_private": true, "uid": bson.M{"$nin": get}})
		for many.Next(context.Background()) {
			fmt.Println(many.Current)
		}
	})
}
