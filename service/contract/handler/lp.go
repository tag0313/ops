package handler

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"ops/pkg/logger"
	"ops/pkg/utils"
	pbContract "ops/proto/contract"
)

type lp struct {

}

func NewLP(opts *Options)(*lp, error){
	l := new(lp)

	return l, nil
}

func dailyOps(opsPerBlock *big.Int, decimals int)*big.Rat{
	//挖矿每天产生OPS个数：24*60*60*opsPerBlock()/(14*10**18)
	opsDailyNew := new(big.Int).Mul(big.NewInt(24*60*60), opsPerBlock)
	den := new(big.Int).Mul(big.NewInt(14), utils.BigIntPow10(decimals))
	return new(big.Rat).SetFrac(opsDailyNew, den)
}

func (l lp) GetOpsUsdtApy(ctx context.Context,
	request *pbContract.GetOpsUsdtApyRequest,
	response *pbContract.GetApyResponse) error {
	var err error
	defer func() {
		logger.Infof("calling GetAPY, request=%+v response=%+v, err=%+v",
			request, response, err)
	}()
	amounts, err := lpOpsUsdtIns.OpsPerBlock(nil)
	if err != nil{
		return err
	}
	//24*60*60*opsPerBlock()/(14*10**18)
	totalSupply, err := swapOps2usdtIns.TotalSupply(nil)
	if err != nil{
		return err
	}

	reserves, err := swapOps2usdtIns.GetReserves(nil)
	if err != nil{
		return err
	}
	opsPrice, err := GetOpsPrice()
	if err != nil{
		return err
	}
	opsDecimals, err := getOPSDecimals()
	if err != nil{
		return err
	}

	opsDailyNew := dailyOps(amounts, int(opsDecimals))

	lpAmount, err := swapOps2usdtIns.BalanceOf(nil, globalOptions.OpsLpUsdtContractAddr)
	if err != nil{
		return err
	}

	opsAPY, err  := CalculateApy(opsDailyNew, big.NewRat(6, 10), //0.6
		totalSupply, reserves.Reserve0, reserves.Reserve1, lpAmount,
		opsPrice, opsPrice, big.NewRat(1,1),
		int(opsDecimals))
	if err != nil{
		return err
	}
	response.Capitalization = opsAPY.FloatString(int(opsDecimals))

	logger.Infof("amounts=%s, opsDailyNew=%s, result=%s",
		amounts.String(), opsDailyNew.String(), response.Capitalization)

	return nil
}

func (l lp) GetOpsFluxAPY(ctx context.Context,
	request *pbContract.GetOpsFluxApyRequest,
	response *pbContract.GetApyResponse) error {
	var err error
	defer func() {
		logger.Infof("calling GetAPY, request=%+v response=%+v, err=%+v",
			request, response, err)
	}()
	amounts, err := lpOpsUsdtIns.OpsPerBlock(nil)
	if err != nil{
		return err
	}
	//24*60*60*opsPerBlock()/(14*10**18)
	totalSupply, err := swapOps2FluxIns.TotalSupply(nil)
	if err != nil{
		return err
	}

	reserves, err := swapOps2FluxIns.GetReserves(nil)
	if err != nil{
		return err
	}
	opsPrice, err := GetOpsPrice()
	if err != nil{
		return err
	}

	fluxPrice, err := GetFluxPrice()
	if err != nil{
		return err
	}
	fluxUsdtPrice := new(big.Rat).Mul(fluxPrice , opsPrice)

	opsDecimals, err := getOPSDecimals()
	if err != nil{
		return err
	}

	opsDailyNew := dailyOps(amounts, int(opsDecimals))
	logger.Infof("lp addr=%s\n", globalOptions.OpsLpUsdtContractAddr)
	lpAmount, err := swapOps2FluxIns.BalanceOf(nil, globalOptions.OpsLpUsdtContractAddr)
	if err != nil{
		return err
	}

	//r0price r1price is reserved
	opsAPY, err  := CalculateApy(opsDailyNew, big.NewRat(4, 10),
		totalSupply, reserves.Reserve0, reserves.Reserve1, lpAmount,
		opsPrice,fluxUsdtPrice, opsPrice,
		int(opsDecimals))
	if err != nil{
		return err
	}
	response.Capitalization = opsAPY.FloatString(int(opsDecimals))

	return nil
}

func (l lp) GetEthUsdtApy(ctx context.Context,
	request *pbContract.GetEthUsdtApyRequest,
	response *pbContract.GetApyResponse) error {
	var err error
	defer func() {
		logger.Infof("calling GetAPY, request=%+v response=%+v, err=%+v",
			request, response, err)
	}()
	price, err := getETHPrice()
	if err != nil{
		return err
	}
	response.Capitalization = price.FloatString(18)
	return nil
}

func (l *lp) GetOpsPriceUsdt(ctx context.Context,
	request *pbContract.GetOpsPriceRequest,
	response *pbContract.WorthResponse) (err error){
	defer func(){
		logger.Infof("calling GetOpsPriceUsdt req=%+v resp=%+v err=%v",
			request, response, err)
	}()
	opsPrice, err := calculatePrice(swapOps2usdtIns)
	if err != nil{
		return err
	}
	response.Wroth = opsPrice.FloatString(18)
	return nil
}

func (l lp) GetOpsPriceFlux(ctx context.Context,
	request *pbContract.GetOpsPriceRequest,
	response *pbContract.WorthResponse) (err error) {
	defer func(){
		logger.Infof("calling GetOpsPriceUsdt req=%+v resp=%+v err=%v",
			request, response, err)
	}()
	price, err := calculatePriceReverse(swapOps2FluxIns)
	if err != nil{
		return err
	}

	response.Wroth = price.FloatString(18)
	return nil
}

func (l lp) GetPoolWorth(ctx context.Context,
	request *pbContract.GetPoolWorthRequest,
	response *pbContract.WorthResponse) (err error) {
	defer func() {
		logger.Infof("calling GetUserLpBalance, " +
			"request=%+v response=%+v, err=%+v", request, response, err)
	}()
	opsPrice, err := calculatePrice(swapOps2usdtIns)
	if err != nil{
		return err
	}

	decimals, err := swapOps2FluxIns.Decimals(nil)
	if err != nil{
		return err
	}

	if request.LpType == pbContract.LpType_OpsFlux{
		fluxPrice, err := GetFluxPriceUsdt(opsPrice)
		if err != nil{
			return err
		}
		reserves, err := swapOps2FluxIns.GetReserves(nil)
		if err != nil{
			return err
		}
		worth := CalculatePoolWorth(reserves.Reserve0, reserves.Reserve1,
			opsPrice, fluxPrice, int(decimals))
		response.Wroth = worth.FloatString(int(decimals))
		return nil
	}else if request.LpType == pbContract.LpType_OpsUSDT{
		reserves, err := swapOps2usdtIns.GetReserves(nil)
		if err != nil{
			return err
		}

		worth := CalculatePoolWorth(reserves.Reserve0, reserves.Reserve1,
			opsPrice, big.NewRat(1,1), int(decimals))

		response.Wroth = worth.FloatString(int(decimals))
		return nil
	}
	return errors.New("no such mining pool")
}

func (l lp) GetUserLpBalance(ctx context.Context,
	request *pbContract.UserLpRequest,
	response *pbContract.BalanceResponse) (err error) {
	defer func() {
		logger.Infof("calling GetUserLpBalance, request=%+v response=%+v, err=%+v",
			request, response, err)
	}()
	if request.LpType == pbContract.LpType_OpsUSDT{
		return getUserOpsUsdtLpBalance(ctx, request, response)
	}else if request.LpType == pbContract.LpType_OpsFlux{
		return getUserOpsFluxLpBalance(ctx, request, response)
	}

	return errors.New("no such mining pool")
}

func (l lp) GetUserLpAmount(ctx context.Context,
	request *pbContract.UserLpRequest,
	response *pbContract.LpAmountResponse)(err error){
	defer func() {
		logger.Infof("calling GetUserLpAmount, request=%+v response=%+v, err=%+v",
			request, response, err)
	}()
	decimals, err := swapOps2usdtIns.Decimals(nil)
	if err != nil{
		logger.Error(err)
		return err
	}

	var (
		amount *big.Int
	)
	//_pid =0 代表 LP（OPS-USDT）1：（OPS-FLUX）
	if request.LpType == pbContract.LpType_OpsUSDT{
		lpPid := big.NewInt(0)
		info, err := lpOpsUsdtIns.UserInfo(nil, lpPid,
			common.HexToAddress(request.AccountAddress))
		if err != nil{
			logger.Errorf("addr=%s, type=%s", request.AccountAddress, request.LpType.String())
			return err
		}
		amount = info.Amount
	}else if request.LpType == pbContract.LpType_OpsFlux{
		lpPid := big.NewInt(1)
		info, err := lpOpsUsdtIns.UserInfo(nil, lpPid,
			common.HexToAddress(request.AccountAddress))
		if err != nil{
			logger.Errorf("addr=%s, type=%s, pid=%s, err=%v",
				request.AccountAddress, request.LpType.String(), lpPid, err)
			return err
		}
		amount = info.Amount
	}else{
		return errors.New("no such mining pool")
	}

	response.Amount = utils.BigIntToFloatStr(amount, int(decimals))
	logger.Debugf("lpType=%d, amount=%d f=%s", request.LpType, amount, response.Amount)
	return nil
}

func (l lp) GetOpsReward(ctx context.Context,
	request *pbContract.UserLpRequest,
	response *pbContract.GetOpsRewardsResponse)(err error) {
	defer func() {
		logger.Infof("calling GetOpsReward, request=%+v response=%+v, err=%+v",
			request, response, err)
	}()
	decimals, err := swapOps2usdtIns.Decimals(nil)
	if err != nil{
		return
	}

	var (
		amount *big.Int
	)
	//_pid =0 代表 LP（OPS-USDT）1：（OPS-FLUX）
	if request.LpType == pbContract.LpType_OpsUSDT{
		lpPid := big.NewInt(0)
		amount, err = lpOpsUsdtIns.PendingOps(nil, lpPid, common.HexToAddress(request.AccountAddress))
		if err != nil{
			return
		}
	}else if request.LpType == pbContract.LpType_OpsFlux{
		lpPid := big.NewInt(1)
		amount, err = lpOpsUsdtIns.PendingOps(nil, lpPid, common.HexToAddress(request.AccountAddress))
		if err != nil{
			return
		}
	}else{
		return errors.New("no such mining pool")
	}
	logger.Debugf("lpType=%d, amount=%d, decimals=%d", request.LpType, amount, decimals)
	response.Rewards = utils.BigIntToFloatStr(amount, int(decimals))
	return nil
}

