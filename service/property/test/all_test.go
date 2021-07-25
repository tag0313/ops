package test

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"ops/pkg/logger"
	"ops/pkg/model/mgodb"
	"testing"
)

//func TestTime(t *testing.T) {
//	t.Run("get time", func(t *testing.T) {
//		fmt.Println(string(time.Now().Unix()))
//	})
//}

func TestMongo(t *testing.T) {
	t.Run("mongo", func(t *testing.T) {
		mgoClient := mgodb.NewMgo("property", "user_opspoint")
		one := mgoClient.FindOne(bson.M{"uid": "7JcS9N1Z3ksDChPRe0cN"})
		if one.Err() != nil {
			logger.Error(one.Err())
		}
		bytes, err := one.DecodeBytes()
		if err != nil {
			logger.Error(err)
		}
		fmt.Println(bytes.Lookup("uid").StringValue())
		fmt.Println(bytes.Lookup("ops_point").Double())
	})
}
