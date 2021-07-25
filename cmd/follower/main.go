package main

import (
	"ops/pkg/daemon"
	"ops/pkg/logger"
	"ops/pkg/model/consulreg"
	"ops/pkg/model/mgodb"
	"ops/pkg/utils"
	"ops/proto/follower"
	"ops/service/follower/handler"
)

func main() {
	flag := daemon.NewCmdFlags()
	if err := flag.Parse(); err != nil{
		panic(err)
	}
	err := utils.LoadConfigFile(flag.ConfigFile)
	if err != nil{
		panic(err)
	}

	loggerOpt := logger.NewOpts(utils.GetConfigStr("log_path"))
	logger.InitDefault(loggerOpt)
	defer logger.Sync()

	//init outside resources
	err = mgodb.InitMongoDB(utils.GetConfigStr("mongodb.url"))
	if err != nil{
		daemon.Exit(-1, err.Error())
	}

	//Create service
	err = consulreg.InitMicro(utils.GetConfigStr("micro.addr"), utils.GetConfigStr("micro.name"))
	if err != nil{
		daemon.Exit(-1, err.Error())
	}
	srv := consulreg.MicroSer
	err = pbFollower.RegisterOperateFollowHandler(srv.Server(), new(handler.Follower))
	if err != nil {
		daemon.Exit(-1, err.Error())
	}

	// Run service
	if err = srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
