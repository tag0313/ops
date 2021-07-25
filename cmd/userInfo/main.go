package main

import (
	"github.com/go-redis/redis/v8"
	"ops/pkg/daemon"
	"ops/pkg/logger"
	"ops/pkg/model/azureStorage"
	"ops/pkg/model/consulreg"
	"ops/pkg/model/mgodb"
	"ops/pkg/model/rdb"
	"ops/pkg/utils"
	"ops/pkg/version"
	"ops/proto/userInfo"
	"ops/service/userInfo/handler"
	"time"
)

func main() {
	flag := daemon.NewCmdFlags()
	if err := flag.Parse(); err != nil{
		daemon.Exit(-1, err.Error())
	}
	err := utils.LoadConfigFile(flag.ConfigFile)
	if err != nil{
		daemon.Exit(-1, err.Error())
	}

	loggerOpt := logger.NewOpts(utils.GetConfigStr("log_path"))
	logger.InitDefault(loggerOpt)
	defer logger.Sync()
	logger.Info(version.Long)

	//init outside resources
	err = mgodb.InitMongoDB(utils.GetConfigStr("mongodb.url"))
	if err != nil{
		daemon.Exit(-1, err.Error())
	}

	err = rdb.InitRedis(&redis.Options{
		Addr:         utils.GetConfigStr("redis.addr"),
		Password:     utils.GetConfigStr("redis.pwd"),
		DB:           utils.GetConfigInt("redis.db"), // use default DB
		DialTimeout:  time.Duration(utils.GetConfigInt64("redis.dial_timeout")) * time.Second,
		ReadTimeout:  time.Duration(utils.GetConfigInt64("redis.read_timeout")) * time.Second,
		WriteTimeout: time.Duration(utils.GetConfigInt64("redis.write_timeout")) * time.Second,
		PoolSize:     utils.GetConfigInt("redis.pool_size"),
		PoolTimeout:  time.Duration(utils.GetConfigInt64("redis.pool_timeout")) * time.Second,
	})
	if err != nil{
		daemon.Exit(-1, err.Error())
	}

	if err = azureStorage.InitAzureStorage(); err != nil{
		daemon.Exit(-1, err.Error())
	}

	err = consulreg.InitMicro(utils.GetConfigStr("micro.addr"), utils.GetConfigStr("micro.name"))
	if err != nil{
		daemon.Exit(-1, err.Error())
	}
	// Create service
	srv := consulreg.MicroSer
	err = pbUserInfo.RegisterOperateUserInfoHandler(srv.Server(), new(handler.UserInfo))
	if err != nil {
		return
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}

}
