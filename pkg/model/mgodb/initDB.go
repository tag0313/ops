package mgodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"ops/pkg/logger"
	"ops/pkg/utils"
	"sync"
	"time"
)

type Database struct {
	Mongo *mongo.Client
}

var (
	DB *Database
	once sync.Once
)

func InitMongoDB(uri string)error{
	var (
		err error
		client *mongo.Client
	)
	once.Do(func() {
		client, err = setConnect(uri)
		DB = &Database{
			Mongo: client,
		}
	})
	return err
}

// 连接设置
func setConnect(uri string) (*mongo.Client, error) {
	if uri == ""{
		uri = utils.GetConfigStr("mongodb.url")
	}
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
