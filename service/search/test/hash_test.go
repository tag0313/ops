package test

import (
	"context"
	"fmt"
	"ops/pkg/model/mgodb"
	"ops/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestHash(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		followClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.following"))
		oopClient := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.ops"), utils.GetConfigStr("mongodb.collection.oop"))
		findOptions := options.Find()
		findOptions.SetProjection(bson.M{"following": 1, "_id": 0})
		many, _ := followClient.FindMany(bson.M{"uid": "oiJIhWYWJDNBubK1mQ38"}, findOptions)

		var followingList []string
		for many.Next(context.Background()) {
			followingList = append(followingList, many.Current.Lookup("following").StringValue())
		}

		fmt.Println(followingList)

		oopOptions := options.Find()
		oopOptions.SetProjection(bson.M{"_id": 0})
		findMany, _ := oopClient.FindMany(bson.M{"uid": bson.M{"$in": followingList}}, oopOptions)

		for findMany.Next(context.Background()) {
			fmt.Println(findMany.Current.String())
		}
	})

}
