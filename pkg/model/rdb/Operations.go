package rdb

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	utils2 "ops/pkg/utils"
	"time"
)

func Get(key string) (string, string) {
	result, err := REDISDB.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return "", utils2.RECODE_DATAINEXISTENCE
	}
	return result, utils2.RECODE_OK
}

func SetS(key string, value interface{}, expiration time.Duration) string {
	err := REDISDB.RedisClient.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return utils2.RECODE_STOREDATA_FAILED
	}
	return utils2.RECODE_STOREDATA_OK
}
func LPushS(key string, value interface{}) string {
	err := REDISDB.RedisClient.LPush(context.Background(), key, value).Err()
	if err != nil {
		logger.Error(err)
		return utils2.RECODE_STOREDATA_FAILED
	}
	return utils2.RECODE_STOREDATA_OK
}

func LRem(key string, count int64, value interface{}) string {
	err := REDISDB.RedisClient.LRem(context.Background(), key, count, value)
	if err.Err() != nil {
		logger.Error(err.Err())
		return utils2.RECODE_NODATA
	}
	return utils2.RECODE_OK
}

func GetList(key string, start, stop int64) []string {
	result := REDISDB.RedisClient.LRange(context.Background(), key, start, stop)
	if result.Err() != nil {
		return nil
	}
	return result.Val()
}

func Del(key string) error {
	err := REDISDB.RedisClient.Del(context.Background(), key).Err()
	return err
}

func LPush(key string, value interface{}) error {
	err := REDISDB.RedisClient.LPush(context.Background(), key, value).Err()
	return err
}

func RPop(key string) (string, error) {
	result, err := REDISDB.RedisClient.RPop(context.Background(), key).Result()
	return result, err
}

func LLen(key string) (int, error) {
	result, err := REDISDB.RedisClient.LLen(context.Background(), key).Result()
	return int(result), err
}
func SetM(key string, value interface{}, expiration time.Duration) error {
	err := REDISDB.RedisClient.Set(context.Background(), key, value, expiration).Err()
	return err
}
