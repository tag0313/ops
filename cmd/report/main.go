package main

import (
	"ops/pkg/daemon"
	"ops/pkg/logger"
	"ops/pkg/model/consulreg"
	"ops/pkg/model/mgodb"
	"ops/pkg/utils"
	"ops/proto/report"
	"ops/service/report/handler"
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

	//Create service
	srv := consulreg.MicroSer
	err = pbReport.RegisterOperateReportHandler(srv.Server(), new(handler.Report))
	if err != nil {
		return
	}

	// Run service
	if err = srv.Run(); err != nil {
		logger.Fatal(err)
	}

}
