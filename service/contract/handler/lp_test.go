package handler

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	pbContract "ops/proto/contract"
	"testing"
)

func TestLp_GetOpsUsdtAPY(t *testing.T) {
	type args struct {
		ctx      context.Context
		request  *pbContract.GetOpsUsdtApyRequest
		response *pbContract.GetApyResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "GetAPY", args: args{
			ctx:      nil,
			request:  nil,
			response: new(pbContract.GetApyResponse),
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lp{}
			if err := l.GetOpsUsdtApy(tt.args.ctx, tt.args.request, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("GetAPY() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("response: %+v", tt.args.response)
		})
	}
}

func TestLp_GetOpsFluxAPY(t *testing.T) {
	type args struct {
		ctx      context.Context
		request  *pbContract.GetOpsFluxApyRequest
		response *pbContract.GetApyResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "GetAPY", args: args{
			ctx:      nil,
			request:  nil,
			response: new(pbContract.GetApyResponse),
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lp{}
			if err := l.GetOpsFluxAPY(tt.args.ctx, tt.args.request, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("GetOpsFluxAPY() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLp_GetEthUsdtApy(t *testing.T) {
	type args struct {
		ctx      context.Context
		request  *pbContract.GetEthUsdtApyRequest
		response *pbContract.GetApyResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "GetAPY", args: args{
			ctx:      nil,
			request:  nil,
			response: new(pbContract.GetApyResponse),
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lp{}
			if err := l.GetEthUsdtApy(tt.args.ctx, tt.args.request, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("GetEthUsdtApy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLp_GetOpsPrice(t *testing.T) {
	var(
		req = new(pbContract.GetOpsPriceRequest)
		rep = new(pbContract.WorthResponse)
	)
	l := lp{}
	err := l.GetOpsPriceUsdt(nil, req, rep)
	if err != nil{
		t.Error(err)
	}
	t.Log(rep.Wroth)
}

func TestLp_GetPoolWorth(t *testing.T) {
	var(
		req = new(pbContract.GetPoolWorthRequest)
		rep = new(pbContract.WorthResponse)
	)
	req.LpType = pbContract.LpType_OpsUSDT
	l := lp{}
	err := l.GetPoolWorth(nil, req, rep)
	if err != nil{
		t.Error(err)
	}
	t.Log("ops usdt pool worth is: ", rep.Wroth)

	req.LpType = pbContract.LpType_OpsFlux
	err = l.GetPoolWorth(nil, req, rep)
	if err != nil{
		t.Error(err)
	}
	t.Log("ops flux pool worth is: ",rep.Wroth)
}

func TestLp_GetOpsPriceFlux(t *testing.T) {
	var(
		req = new(pbContract.GetOpsPriceRequest)
		rep = new(pbContract.WorthResponse)
	)
	l := lp{}
	err := l.GetOpsPriceFlux(nil, req, rep)
	if err != nil{
		t.Error(err)
	}
	t.Log(rep.Wroth)
}

func lpCall(req *pbContract.UserLpRequest)(
	Balance string,
	Amount string,
	Rewards string,
	err error){
	var(
		balanceRep = new(pbContract.BalanceResponse)
		amountRep = new(pbContract.LpAmountResponse)
		rewardRep = new(pbContract.GetOpsRewardsResponse)
	)
	l := lp{}
	err = l.GetUserLpBalance(nil, req, balanceRep)
	if err != nil{
		return
	}

	err = l.GetUserLpAmount(nil, req, amountRep)
	if err != nil{
		return
	}

	err = l.GetOpsReward(nil, req, rewardRep)
	if err != nil{
		return
	}
	return balanceRep.BalanceStr, amountRep.Amount, rewardRep.Rewards, nil
}

func TestLp_GetLpInfo(t *testing.T) {
	logger.Info("ops-flux addr",testOptions.SwapOpsUsdtContractAddr)
	var(
		req = new(pbContract.UserLpRequest)
		addr1 = "0x020EEEB7494BCa07F47EC5DBC6ADC5E3a4073263"
		addr2 = "0x66d59cA5721Ce058B706581d983bbD7c5bA366f1"
	)
	req.LpType = pbContract.LpType_OpsUSDT
	req.AccountAddress= addr1
	b, a, r, err := lpCall(req)
	if err != nil{
		t.Error(err)
	}
	t.Logf("address=%s, type=%s, balance=%s, " +
		"amount=%s, rewards=%s",
		req.AccountAddress, req.LpType.String(), b,a,r)

	newReq := &pbContract.UserLpRequest{
		AccountAddress:addr1,
		LpType: pbContract.LpType_OpsFlux,
	}
	b, a, r, err = lpCall(newReq)
	if err != nil{
		t.Error(err)
	}
	t.Logf("address=%s, type=%s, balance=%s, " +
		"amount=%s, rewards=%s",
		newReq.AccountAddress, newReq.LpType.String(), b,a,r)

	//address 2 ===============================
	req.AccountAddress= addr2
	req.LpType = pbContract.LpType_OpsUSDT
	b, a, r, err = lpCall(req)
	if err != nil{
		t.Error(err)
	}
	t.Logf("address=%s, type=%s, balance=%s, " +
		"amount=%s, rewards=%s",
		req.AccountAddress, req.LpType.String(), b,a,r)

	req.LpType = pbContract.LpType_OpsFlux
	b, a, r, err = lpCall(req)
	if err != nil{
		t.Error(err)
	}
	t.Logf("address=%s, type=%s, balance=%s, " +
		"amount=%s, rewards=%s",
		req.AccountAddress, req.LpType.String(), b,a,r)
}
