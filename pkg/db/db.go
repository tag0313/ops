package db

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ops/pkg/logger"
	"ops/pkg/utils"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	db *mongo.Client
	once sync.Once
)

func InitMongoDB(uri string)error{
	var err error
	once.Do(func() {
		db, err = setConnect(uri)
	})
	return err
}

// 连接设置
func setConnect(uri string) (*mongo.Client, error) {
	logger.Info("connecting mongodb, url is: ", uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(utils.GetConfigUint("mongodb.pool_size"))) // 连接池
	if err != nil{
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Errorf("connecting mongodb failed, err=%v, uri=%s", err, uri)
		return nil, err
	}
	logger.Infof("connecting mongodb: %s successfully", uri)
	return client, nil
}

type Record interface {
	DbName()	string
	CollectionName()string
	SetOID(id string)
}

func getOperateCollection(r Record)(*mongo.Collection, error){
	dbName := r.DbName()
	collection := r.CollectionName()
	if dbName == "" || collection == ""{
		return nil, errors.New("database name or collection should not be nil")
	}
	return db.Database(dbName).Collection(collection), nil
}

type table struct{
}

func NewOpTable()*table{
	return &table{}
}

func (t *table)InsertOne(ctx context.Context, r Record)error{
	var (
		c *mongo.Collection
		err error
	)
	if c, err = getOperateCollection(r); err != nil{
		return err
	}

	var insertResult *mongo.InsertOneResult
	insertResult, err = c.InsertOne(ctx, r)
	if err != nil{
		return err
	}
	oid := insertResult.InsertedID.(primitive.ObjectID)
	r.SetOID(oid.String())
	return nil
}

func (t *table)InsertMany(ctx context.Context, rs []interface{Record})error{
	if len(rs) == 0{
		return errors.New("the slice of records cannot be empty")
	}
	c, err := getOperateCollection(rs[0])
	if err != nil{
		return err
	}
	var docs []interface{}
	for _, r := range rs {
		docs = append(docs, r)
	}
	var result *mongo.InsertManyResult
	result, err = c.InsertMany(ctx, docs)
	if err != nil{
		return err
	}

	if len(rs) != len(result.InsertedIDs){
		return errors.New("insert error, " +
			"the insert result cannot be equal records data")
	}

	for i, oid := range result.InsertedIDs {
		rs[i].SetOID(oid.(primitive.ObjectID).String())
	}

	return nil
}