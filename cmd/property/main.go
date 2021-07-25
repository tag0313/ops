package main

import (
	"ops/pkg/daemon"
	"ops/pkg/db"
	"ops/pkg/logger"
	"ops/pkg/model/consulreg"
	"ops/pkg/model/mgodb"
	"ops/pkg/utils"
	"ops/proto/property"
	"ops/service/property/handler"
)

func main() {
	flag := daemon.NewCmdFlags()
	if err := flag.Parse(); err != nil {
		panic(err)
	}
	err := utils.LoadConfigFile(flag.ConfigFile)
	if err != nil {
		panic(err)
	}

	loggerOpt := logger.NewOpts(utils.GetConfigStr("log_path"))
	logger.InitDefault(loggerOpt)
	defer logger.Sync()

	//init outside resources
	err = mgodb.InitMongoDB(utils.GetConfigStr("mongodb.url"))
	if err != nil {
		daemon.Exit(-1, err.Error())
	}

	err = db.InitMongoDB(utils.GetConfigStr("mongodb.url"))
	if err != nil {
		daemon.Exit(-1, err.Error())
	}

	//Create service
	err = consulreg.InitMicro(utils.GetConfigStr("micro.addr"), utils.GetConfigStr("micro.name"))
	if err != nil {
		daemon.Exit(-1, err.Error())
	}

	// Create service
	srv := consulreg.MicroSer
	err = pbProperty.RegisterOperatePropertyHandler(srv.Server(), new(handler.Property))
	if err != nil {
		return
	}

	// Run service
	if err = srv.Run(); err != nil {
		logger.Fatal(err)
	}

}
