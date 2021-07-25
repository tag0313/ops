package handler

import (
	"ops/proto/swap"
	"testing"
)


func initNewSwapHandle()(*swapHandler, error){
	//server :="https://rinkeby.infura.io/v3/aacefa715bbb4a23b6ad98cf61cec2bc"
	//opsUsdtContract := "0xe8788cde4b81587386230b36093b96cbdf7f6a0a"
	//swapEthUsdtContract := "0x5308a481B2b65F6086083D2268acb73AADC757E0"
	h, err := NewSwapHandler(testOptions)
	if err != nil{
		return nil, err
	}
	if err = h.CheckConnection(); err != nil{
		return nil, err
	}
	return h, err
}

func TestSwap(t *testing.T) {
	h, err := initNewSwapHandle()
	if err != nil {
		t.Fatal(err)
	}
	req := &pbSwap.MoneyRequest{}
	rep := &pbSwap.MoneyResponse{}
	err = h.OPS2USDT(nil, req, rep)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("opsPrice=", rep)

	err = h.ETH2USDT(nil, req, rep)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ethPrice=", rep)
}

