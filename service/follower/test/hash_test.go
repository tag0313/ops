package test

import (
	"context"
	"fmt"
	"ops/pkg/model/mgodb"
	"ops/pkg/utils"
	"sort"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func TestRelation(t *testing.T) {
	t.Run("relation", func(t *testing.T) {
		list := make([]string, 3)
		list[0] = "k5PuyZpeo7GxJZVLGlgN"
		list[1] = "yArSC5HmIMFjYbeFdAil"
		list[2] = "DfjMEHLrzDewv3pa7i0f"
		follow := mgodb.NewMgo(utils.GetConfigStr("mongodb.db.user_info"), utils.GetConfigStr("mongodb.collection.following"))
		many, _ := follow.FindMany(bson.M{"uid": "k5PuyZpeo7GxJZVLGlgN", "following": bson.M{"$in": list}})

		var currentUid string
		resultMap := make(map[string]string)
		for many.Next(context.TODO()) {
			currentUid = many.Current.Lookup("following").StringValue()
			i := sort.Search(len(list), func(i int) bool {
				return list[i] >= currentUid
			})
			if i < len(list) && list[i] == currentUid { //这里可以采用 strings.EqualFold(arrString[i],target)
				resultMap[currentUid] = currentUid
			}
		}

		resultList := make([]string, len(resultMap))
		var index int
		for key := range resultMap {
			resultList[index] = key
			index++
		}
		fmt.Println(resultList)
	})
}
