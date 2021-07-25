package handler

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	microErr "github.com/asim/go-micro/v3/errors"
	"github.com/asim/go-micro/v3/registry"
	"ops/pkg/utils"
	"ops/proto/contract"
	"testing"
	"time"
)


func newOPSHandler()(*opsHandler, error){
	//contractAddress := "0xb247BefC358BeDf8E0f305826350cC40184879c5"
	//adminAddress := "0x66d59cA5721Ce058B706581d983bbD7c5bA366f1"
	//testNetworkURL := "https://rinkeby.infura.io/v3/aacefa715bbb4a23b6ad98cf61cec2bc"
	h, err := NewHandler(testOptions)
	if err != nil{
		return nil, err
	}
	return h, nil
}

func TestOPSGetMethod(t *testing.T) {
	adminAddress := "0x66d59cA5721Ce058B706581d983bbD7c5bA366f1"
	h, err := newOPSHandler()
	if err != nil{
		t.Fatal(err)
	}
	ctx := context.Background()
	var nameResponse pbContract.NameResponse
	if err := h.Name(ctx,nil, &nameResponse); err!=nil{
		t.Fatal(err)
	}
	t.Logf("name is %s",nameResponse.Name)

	var symbolResponse pbContract.SymbolResponse
	if err = h.Symbol(ctx, nil, &symbolResponse); err != nil{
		t.Fatal(err)
	}
	t.Logf("symbol is %s", symbolResponse.Symbol)

	var totalSupplyResponse pbContract.TotalSupplyResponse
	if err = h.TotalSupply(ctx, nil, &totalSupplyResponse); err != nil{
		t.Fatal(err)
	}
	t.Logf("total supply is %s", totalSupplyResponse.TotalSupply)

	var decimals pbContract.DecimalsResponse
	if err = h.Decimals(ctx, nil, &decimals); err !=nil{
		t.Fatal(err)
	}
	t.Logf("decimals is %d", decimals.Decimals)

	var owner pbContract.OwnerResponse
	if err = h.Owner(ctx, nil, &owner); err != nil{
		t.Fatal(err)
	}
	t.Logf("owner address is %s",owner.Address)
	t.Logf("owner address is %s 0x%x",owner.Address, owner.GetAddress())

	var (
		balanceRequest pbContract.BalanceRequest
		balanceResponse pbContract.BalanceResponse
	)
	balanceRequest.Address = adminAddress
	if err = h.BalanceOf(ctx, &balanceRequest, &balanceResponse); err != nil{
		t.Fatal(err)
	}
	if balanceResponse.Decimals == 0 {
		t.Fatal("decimal calling error")
	}
	f := utils.Bytes2bigInt(balanceResponse.Balance)
	t.Logf("address %s balance is %s", balanceResponse.Address, f.String())
}


func TestOPSGetTransactionByHash(t *testing.T) {
	h, err := newOPSHandler()
	if err != nil{
		t.Fatal(err)
	}
	request := &pbContract.GetTransactionByHashRequest{}
	response := &pbContract.GetTransactionByHashResponse{}
	request.TransactionHash = "0xda605a795bee80e9528c77d6ce572598dc153f74f9c6f7cbcf1fb8da36e45a9c"
	err = h.GetTransactionByHash(context.TODO(), request, response)
	if err != nil{
		t.Fatal(err)
	}
}


func TestErrors(t *testing.T){
	consulReg := consul.NewRegistry(
		registry.Addrs(testOptions.ListenAddr))
	srv := micro.NewService(
		micro.Name(testOptions.MicroName),
		micro.Registry(consulReg),
		micro.Version(testOptions.Version),
	)

	t.Log("calling errors ")

	rpc := pbContract.NewContractService("contract", srv.Client())
	opts := func(o *client.CallOptions) {
		o.RequestTimeout = time.Second * 5
	}
	_, err := rpc.TestError(context.TODO(),new(pbContract.Empty), opts)
	if err != nil{
		t.Log("Got a error is: ",err)
	}
	verr, ok := err.(*microErr.Error)
	if ok {
		t.Fatalf("error received %#+v\n", verr)
	}
}

func TestOpsHandler_DonateCoin(t *testing.T) {
	h, err := newOPSHandler()
	if err != nil{
		t.Fatal(err)
	}
	var(
		request = &pbContract.DonateCoinRequest{
			//肥水不流外人田。。。。
			UserPublicKey: "0xC14c8ea6757d3504E1056c0A50563e71E33DfC15",
		}
		response pbContract.DonateCoinResponse
	)
	err = h.DonateCoin(context.TODO(), request, &response)
	if err != nil{
		t.Fatal(err)
	}
}