package handler

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/big"
	"strconv"
	"time"
)

type Property struct {
}

type UserOpsPoint struct {
	Uid      string  `bson:"uid"`
	OpsPoint float64 `bson:"ops_point,omitempty"`
}

type UserInfo struct {
	Uid     string `bson:"uid"`
	PubAddr string `bson:"pubkaddr"`
}

func intBytesToInt64(bytes []byte) int64 {
	bi := new(big.Int)
	bi.SetBytes(bytes)
	return bi.Int64()
}

func intBytesToString(bytes []byte) string {
	bi := new(big.Int)
	bi.SetBytes(bytes)
	return bi.String()
}

func intBytesToInt(bytes []byte) int64 {
	bi := new(big.Int)
	bi.SetBytes(bytes)
	return bi.Int64()
}

func stringDecimalToBytes(decimal string) []byte {
	bi := new(big.Int)
	bi.SetString(decimal, 10)
	return bi.Bytes()
}

func bytes2bigInt(bytes []byte) *big.Int {
	bi := new(big.Int)
	bi.SetBytes(bytes)
	return bi
}

func bigInt2Bytes(bi *big.Int) []byte {
	if bi == nil {
		return nil
	}
	return bi.Bytes()
}

func arrayBigInt2Bytes(bisArray []*big.Int) [][]byte {
	bts := make([][]byte, len(bisArray))
	if bisArray == nil {
		return nil
	}
	for index, bis := range bisArray {
		bts[index] = bigInt2Bytes(bis)
	}
	return bts
}

func getRegisterTime() string {
	timestamp := time.Now().Unix()
	return strconv.FormatInt(timestamp, 10)
}

func mgoFilterCaseInsensitive(key, value string) bson.D {
	//return bson.D{
	//	bson.E{
	//		Key:key, Value: bson.D{
	//			{"$regex", primitive.Regex{Pattern: value , Options:"$i"}},
	//		},
	//	},
	//}
	return bson.D{{key, primitive.Regex{Pattern: value, Options: "$i"}}}
}
