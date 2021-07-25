package rdb

import (
	"context"
	"github.com/go-redis/redis/v8"
	"ops/pkg/logger"
	"sync"
)

type Redisdb struct {
	RedisClient *redis.Client
}

var (
	REDISDB *Redisdb
	once sync.Once
)

func InitRedis(opts *redis.Options)error{
	var err error
	once.Do(func() {
		rdb := redis.NewClient(opts)
		err = rdb.Ping(context.TODO()).Err()
		if err != nil {
			logger.Error("The connection of redis is failed.")

		}
		REDISDB = &Redisdb{
			RedisClient: rdb,
		}
		logger.Infof("connecting redis: %s successfully", opts.Addr)
	})

	return err
}