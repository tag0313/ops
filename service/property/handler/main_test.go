package handler

import (
	"fmt"
	"ops/pkg/daemon"
	"ops/pkg/db"
	"ops/pkg/logger"
	"ops/pkg/model/consulreg"
	"ops/pkg/utils"
	"os"
	"testing"
)

var(
	configStr =`
	micro:
	  addr: localhost
	  name: property
	  version: 1.0
	mongodb:
	  url: mongodb://localhost:27017
	  pool_size: 20
	  db:
		property: property
		user_info: user_info
	  collection:
		minted_ocard: minted_ocard
		buyer_info_ops: buyer_info_ops
		user_opspoint: user_opspoint
		buyer_info_chain: buyer_info_chain
		transfer_erc20_history: transfer_erc20_history
		withdraw_opspoint_history: withdraw_opspoint_history
		withdraw_ocard_history: withdraw_ocard_history
`
)

func TestMain(m *testing.M) {
	fmt.Println("running test main")
	logger.InitDefault(&logger.Options{
		AbsLogDir: "",
		Debug: true,
	})
	defer logger.Sync()

	if err := utils.LoadConfigString(configStr);err!=nil{
		daemon.Exit(-1, err.Error())
	}

	logger.Info("test main running before")
	err := db.InitMongoDB(utils.GetConfigStr("mongodb.url"))

	if err != nil{
		daemon.Exit(-1, err.Error())
	}

	//Create service
	err = consulreg.InitMicro("localhost", "property")
	if err != nil{
		daemon.Exit(-1, err.Error())
	}

	code := m.Run()

	logger.Info("test main running after")

	os.Exit(code)
}