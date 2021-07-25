
package handler

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"ops/pkg/logger"
	"ops/pkg/utils"
	pbContract "ops/proto/contract"
	"ops/service/contract/ercpkg/erc20basic"
	"ops/service/contract/ercpkg/ops"
	contractSwap "ops/service/contract/ercpkg/swap"
)

func getUSDTDecimals()(decimals uint8, err error){
	usdtContractAddr, err := swapOps2usdtIns.Token1(nil)
	if err != nil{
		return 0, err
	}
	ins, err := erc20basic.NewErc20basic(usdtContractAddr, ethereumClient)
	if err != nil{
		return
	}
	return ins.Decimals(nil)
}

func getOPSDecimals()(decimals uint8, err error){
	opsSwapAddr, err := swapOps2usdtIns.Token0(nil)
	if err !=nil{
		return 0, err
	}
	ins, err := ops.NewOps(opsSwapAddr, ethereumClient)
	if err != nil{
		return
	}
	return ins.Decimals(nil)
}

func getFluxDecimals()(decimals uint8, err error){
	fluxSwapAddr, err := swapOps2FluxIns.Token1(nil)
	if err !=nil{
		return 0, err
	}
	ins, err := ops.NewOps(fluxSwapAddr, ethereumClient)
	if err != nil{
		return
	}
	return ins.Decimals(nil)
}

func getETHDecimals()(decimals uint8,err error){
	usdtContractAddr, err := swapEth2usdtIns.Token0(nil)
	if err != nil{
		return 0, err
	}
	ins, err := erc20basic.NewErc20basic(usdtContractAddr, ethereumClient)
	if err != nil{
		return
	}
	return ins.Decimals(nil)
	return
}

func getETHPrice()(*big.Rat, error){
	return calculatePrice(swapEth2usdtIns)
}

func GetOpsPrice()(*big.Rat, error){
	return calculatePrice(swapOps2usdtIns)
}

func GetFluxPrice()(*big.Rat, error){
	return calculatePrice(swapOps2FluxIns)
}

func CalculateApy(opsDailyNew, dailyPer *big.Rat,
	totalSupply, r0, r1, lpAmount *big.Int,
	opsPrice, r0Price, r1Price *big.Rat,
	decimals int )(*big.Rat, error){

	logger.Infof("opsDailyNew=%s, dailyPer=%s, " +
		"totalSupply=%d, r0=%d, r1=%d, lpAmount=%d, " +
		"opsPrice=%s, r0Price=%s, r1Price=%s, " +
		"decimals=%d",
		opsDailyNew.FloatString(18), dailyPer.FloatString(18),
		totalSupply, r0, r1, lpAmount,
		opsPrice.FloatString(18), r0Price.FloatString(18), r1Price.FloatString(18),
		decimals,
	)
	if lpAmount.Cmp(big.NewInt(0)) == 0{
		return big.NewRat(0, 1), nil
	}
	lpAmountR := new(big.Rat).Quo(
		new(big.Rat).SetInt(lpAmount),
		new(big.Rat).SetInt(utils.BigIntPow10(decimals)),
		)

	//lpPrice = (R0 *Price0+R1 *Price1)/totalSupply()
	r0Sum := new(big.Rat).Mul(new(big.Rat).SetInt(r0), r0Price)
	r1Sum := new(big.Rat).Mul(new(big.Rat).SetInt(r1), r1Price)
	lpPrice := new(big.Rat).Quo(new(big.Rat).Add(r0Sum, r1Sum), new(big.Rat).SetInt(totalSupply))
	logger.Infof("r0=%s, r1=%s,  supply=%d, lpPrice=%s",
		r0Sum, r1Sum, totalSupply, lpPrice.FloatString(18))

	//挖矿每天产生OPS个数 *百分比* OPS价格
	opsValueF := new(big.Rat).Mul(opsDailyNew, dailyPer)
	opsValueF = new(big.Rat).Mul(opsValueF, opsPrice)

	//（lp代币价格 * 质押lp代币数量）
	logger.Infof("lpPrice(%s), lpAmountF(%s)", lpPrice.FloatString(18), lpAmountR.FloatString(18))
	lpValueF := new(big.Rat).Mul(lpPrice, lpAmountR)

	logger.Infof("opsValue(%s) / lpValue(%s)",
		opsValueF.FloatString(18), lpValueF.FloatString(18))
	//APY = (挖矿每天产生OPS个数 *百分比* OPS价格 / （lp代币价格 * 质押lp代币数量）)* 365天
	d := new(big.Rat).Quo(opsValueF, lpValueF)
	return new(big.Rat).Mul(new(big.Rat).SetInt64(365), d), nil
}


func getDecimals(contractAddr common.Address)(decimals uint8, err error){
	ins, err := ops.NewOps(contractAddr, ethereumClient)
	if err != nil{
		return
	}
	return ins.Decimals(nil)
}

func calculatePrice(swaIns *contractSwap.Swap) (*big.Rat, error){
	token0Addr, err := swaIns.Token0(nil)
	if err != nil{
		return nil, err
	}
	token1Addr, err := swaIns.Token1(nil)
	if err != nil{
		return nil, err
	}

	token0Decimals, err :=  getDecimals(token0Addr)
	if err != nil{
		return nil, err
	}
	token1Decimals, err :=  getDecimals(token1Addr)
	if err != nil{
		return nil, err
	}

	r, err := swaIns.GetReserves(nil)
	if err != nil{
		return nil, err
	}
	price := exchangeFormula(r.Reserve0, r.Reserve1, token0Decimals, token1Decimals)
	name, _ :=  swaIns.Name(nil)
	logger.Infof("%s price: %s", name, price.FloatString(int(token0Decimals)))
	return price, nil
}

func calculatePriceReverse(swaIns *contractSwap.Swap) (*big.Rat, error){
	token0Addr, err := swaIns.Token0(nil)
	if err != nil{
		return nil, err
	}
	token1Addr, err := swaIns.Token1(nil)
	if err != nil{
		return nil, err
	}

	token0Decimals, err :=  getDecimals(token0Addr)
	if err != nil{
		return nil, err
	}
	token1Decimals, err :=  getDecimals(token1Addr)
	if err != nil{
		return nil, err
	}

	r, err := swaIns.GetReserves(nil)
	if err != nil{
		return nil, err
	}
	price := exchangeFormula(r.Reserve1, r.Reserve0,  token1Decimals, token0Decimals)
	name, _ :=  swaIns.Name(nil)
	logger.Infof("%s reverse price: %s", name, price.FloatString(int(token0Decimals)))
	return price, nil
}

func CalculatePoolWorth(r0, r1 *big.Int,
	r0price, r1price *big.Rat,
	decimals int,
	) *big.Rat{
	logger.Infof("r0=%d, r1=%d, r0Price=%s, r1price=%s, decimals=%d",
		r0, r1,
		r0price.FloatString(decimals), r1price.FloatString(decimals),
		decimals)
	r0Worth := new(big.Rat).Mul(new(big.Rat).SetInt(r0), r0price)
	r1Worth := new(big.Rat).Mul(new(big.Rat).SetInt(r1), r1price)

	addedValue := new(big.Rat).Add(r0Worth, r1Worth)
	p := utils.BigIntPow10(decimals)
	return new(big.Rat).Quo(addedValue, new(big.Rat).SetInt(p))
}

func GetFluxPriceUsdt(opsPriceUsdt *big.Rat)(*big.Rat, error){
	fluxPrice, err := GetFluxPrice()
	if err != nil{
		return nil, err
	}
	return new(big.Rat).Mul(fluxPrice, opsPriceUsdt), nil
}

func getUserOpsUsdtLpBalance(ctx context.Context,
	request *pbContract.UserLpRequest,
	response *pbContract.BalanceResponse) error{
	balance, err := swapOps2usdtIns.BalanceOf(nil,
		common.HexToAddress(request.AccountAddress))
	if err != nil{
		return err
	}
	decimals, err := swapOps2usdtIns.Decimals(nil)
	if err != nil{
		return err
	}
	response.BalanceStr = utils.BigIntToFloatStr(balance, int(decimals))
	return nil
}

func getUserOpsFluxLpBalance(ctx context.Context,
	request *pbContract.UserLpRequest,
	response *pbContract.BalanceResponse) error{
	balance, err := swapOps2FluxIns.BalanceOf(nil,
		common.HexToAddress(request.AccountAddress))
	if err != nil{
		return err
	}
	decimals, err := swapOps2FluxIns.Decimals(nil)
	if err != nil{
		return err
	}
	response.BalanceStr = utils.BigIntToFloatStr(balance, int(decimals))
	return nil
}