package mgodb

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ops/pkg/utils"
	"strconv"
	"time"
)

type mgo struct {
	database   string
	collection string
}

func NewMgo(database, collection string) *mgo {

	return &mgo{
		database,
		collection,
	}
}

// 查询单个
func (m *mgo) FindOne(filter interface{}) *mongo.SingleResult {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	singleResult := collection.FindOne(context.TODO(), filter)
	return singleResult
}

//查询多个
func (m *mgo) FindManyRecords(filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	result, err := collection.Find(context.TODO(), filter, opts...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return result, nil
}

//查询多个
func (m *mgo) FindMany(filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, string) {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	result, err := collection.Find(context.TODO(), filter, opts...)
	if err != nil {
		logger.Error(err)
		return nil, utils.RECODE_DATAINEXISTENCE
	}

	return result, utils.RECODE_OK
}

//插入单个
func (m *mgo) InsertOne(value interface{}) string {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	_, err := collection.InsertOne(context.TODO(), value)
	if err != nil {
		logger.Error(err)
		return utils.RECODE_STOREDATA_FAILED
	}
	return utils.RECODE_OK
}

//InsertMany 插入多个
func (m *mgo) InsertMany(values []interface{}) (*mongo.InsertManyResult, error) {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)

	result, err := collection.InsertMany(context.TODO(), values)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	return result, err
}

//更新单个
func (m *mgo) UpsertOne(filter interface{}, value interface{}) string {
	opitons := &options.UpdateOptions{}
	opitons.SetUpsert(true)

	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	_, err := collection.UpdateOne(context.TODO(),
		filter,
		value,
		opitons)

	if err != nil {
		logger.Error(err)
		return utils.RECODE_STOREDATA_FAILED
	}
	return utils.RECODE_OK
}

//更新单个
func (m *mgo) UpsertRecordOne(ctx context.Context,
	filter interface{}, value interface{}) (*mongo.UpdateResult, error) {
	options := &options.UpdateOptions{}
	options.SetUpsert(true)

	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	return collection.UpdateOne(ctx,
		filter,
		value,
		options)
}

//更新并返回更新后的结果
func (m *mgo) FindOneAndUpdate(filter interface{}, value interface{}) *mongo.SingleResult {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	update := options.FindOneAndUpdate()
	update.SetReturnDocument(options.After)
	update.SetProjection(bson.M{"_id": 0})
	andUpdate := collection.FindOneAndUpdate(context.TODO(), filter, value, update)
	return andUpdate
}

//查询集合里有多少数据
func (m *mgo) CollectionCount() (string, int64) {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	name := collection.Name()
	size, _ := collection.EstimatedDocumentCount(context.TODO())
	return name, size
}

//按选项查询集合 Skip 跳过 Limit 读取数量 sort 1 ，-1 . 1 为最初时间读取 ， -1 为最新时间读取
func (m *mgo) CollectionDocuments(Skip, Limit int64, Sort, Filter interface{}) (*mongo.Cursor, string) {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	findOptions := options.Find().SetSort(Sort).SetLimit(Limit).SetSkip(Skip)
	temp, err := collection.Find(context.Background(), Filter, findOptions)
	if err != nil {
		logger.Error(err)
		return nil, utils.RECODE_DATAERR
	}
	return temp, utils.RECODE_OK
}

//获取集合创建时间和编号
func (m *mgo) ParsingId(result string) (time.Time, uint64) {
	temp1 := result[:8]
	timestamp, _ := strconv.ParseInt(temp1, 16, 64)
	dateTime := time.Unix(timestamp, 0) //这是截获情报时间 时间格式 2019-04-24 09:23:39 +0800 CST
	temp2 := result[18:]
	count, _ := strconv.ParseUint(temp2, 16, 64) //截获情报的编号
	return dateTime, count
}

//删除文章和查询文章
func (m *mgo) DeleteAndFind(key string, value interface{}) (int64, *mongo.SingleResult) {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	filter := bson.D{{key, value}}
	singleResult := collection.FindOne(context.TODO(), filter)
	DeleteResult, err := collection.DeleteOne(context.TODO(), filter, nil)
	if err != nil {
		fmt.Println("删除时出现错误，你删不掉的~")
	}
	return DeleteResult.DeletedCount, singleResult
}

//删除文章
func (m *mgo) Delete(filter interface{}) string {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	_, err := collection.DeleteOne(context.TODO(), filter, nil)
	if err != nil {
		logger.Error(err)
		return utils.RECODE_DATAINEXISTENCE
	}
	return utils.RECODE_OK

}

//删除多个
func (m *mgo) DeleteMany(key string, value interface{}) int64 {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	filter := bson.D{{key, value}}

	count, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}
	return count.DeletedCount
}

//Count
func (m *mgo) Count(filter interface{}) (int64, string) {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	itemCount, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		logger.Error(err)
		return 0, utils.RECODE_USERNOEXISTSERR
	}
	return itemCount, ""

}

func (m *mgo) CountDocument(filter interface{}) (itemCount int64, err error) {
	client := DB.Mongo
	collection := client.Database(m.database).Collection(m.collection)
	return	collection.CountDocuments(context.TODO(), filter)
}
