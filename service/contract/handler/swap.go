package handler

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"ops/pkg/logger"
	"ops/pkg/utils"
	"ops/proto/swap"
	"ops/service/contract/ercpkg/erc20basic"
	"ops/service/contract/ercpkg/ops"
	contractSwap "ops/service/contract/ercpkg/swap"
)

type swapHandler struct {
	client *ethclient.Client
	contractInstance *contractSwap.Swap

	ops2usdtPairAddr common.Address
	eth2usdtPairAddr common.Address

	ops2usdtContractIns *contractSwap.Swap
	eth2usdtContractIns *contractSwap.Swap

	opsClient *ops.Ops
}

//NewSwapHandler Warning: check the pair address carefully, don't passed mix them.
func NewSwapHandler(opts *Options) (*swapHandler, error){
	var(
		err error
		s = new(swapHandler)
	)
	s.client, err = ethclient.Dial(opts.BlockChainURLHttp)
	if err != nil{
		return nil, err
	}
	s.ops2usdtPairAddr = opts.SwapOpsUsdtContractAddr
	s.eth2usdtPairAddr = opts.SwapEthUsdtContractAddr

	s.ops2usdtContractIns, err = contractSwap.NewSwap(s.ops2usdtPairAddr, s.client)
	if err != nil{
		return nil, err
	}
	s.eth2usdtContractIns, err = contractSwap.NewSwap(s.eth2usdtPairAddr, s.client)
	if err != nil{
		return nil, err
	}
	return s, err
}

func (s swapHandler)CheckConnection()error{
	name, err := s.ops2usdtContractIns.Name(nil)
	if err != nil{
		return err
	}
	sym, _ := s.eth2usdtContractIns.Symbol(nil)
	logger.Info("ops2usdt contract name", name, sym)
	name, err = s.eth2usdtContractIns.Name(nil)
	if err != nil{
		return err
	}
	sym, _ = s.eth2usdtContractIns.Symbol(nil)
	logger.Info("eth2usdt contract name", name, sym)
	return nil
}

func (s swapHandler) getUSDTDecimals()(decimals uint8, err error){
	usdtContractAddr, err := s.ops2usdtContractIns.Token1(nil)
	if err != nil{
		return 0, err
	}
	ins, err := erc20basic.NewErc20basic(usdtContractAddr, s.client )
	if err != nil{
		return
	}
	return ins.Decimals(nil)
}

func (s swapHandler) getOPSDecimals()(decimals uint8, err error){
	opsAddr, err := s.ops2usdtContractIns.Token0(nil)
	if err !=nil{
		return 0, err
	}
	ins, err := ops.NewOps(opsAddr, s.client )
	if err != nil{
		return
	}
	return ins.Decimals(nil)
}

func (s swapHandler) getETHDecimals()(decimals uint8,err error){
	usdtContractAddr, err := s.eth2usdtContractIns.Token0(nil)
	if err != nil{
		return 0, err
	}
	ins, err := erc20basic.NewErc20basic(usdtContractAddr, s.client )
	if err != nil{
		return
	}
	return ins.Decimals(nil)
}

//exchangeFormula
/*  example
	reserve1 is the count of usdt
	reserve0 is the count of ops
	R1 = reserve1_num /(10 **  token1_decimals())
	R0 = reserve0_num /(10 **  token0_decimals())
	Ops(USDT) = R1 / R0
 */
func exchangeFormula(r0, r1 *big.Int, decimals0, decimals1 uint8)*big.Rat{
	p0 := utils.BigIntPow10(int(decimals0))
	p1 := utils.BigIntPow10(int(decimals1))

	rat0 := new(big.Rat).SetFrac(r0, p0)
	rat1 := new(big.Rat).SetFrac(r1, p1)

	return new(big.Rat).Quo(rat1, rat0)
}

func (s swapHandler) OPS2USDT(ctx context.Context, request *pbSwap.MoneyRequest, response *pbSwap.MoneyResponse) error {
	m, valid := new(big.Rat).SetString(request.Money)
	if !valid || m.Cmp(new(big.Rat).SetInt64(0))==0{
		return errors.New("invalid float number")
	}

	opsDecimals, err :=  s.getOPSDecimals()
	if err != nil{
		return err
	}
	usdtDecmals, err := s.getUSDTDecimals()
	if err != nil{
		return err
	}
	r, err := s.ops2usdtContractIns.GetReserves(nil)
	if err != nil{
		return err
	}

	//1 is USDT，0 isOps
	price := exchangeFormula(r.Reserve0, r.Reserve1, opsDecimals, usdtDecmals)
	logger.Infof("ops reserves: %+v", r)

	if request.Money == "" {
		response.Money = price.FloatString(int(usdtDecmals))
	}else{
		response.Money = new(big.Rat).Mul(price, m).FloatString(int(usdtDecmals))
	}
	logger.Infof("price = %s", response.Money)
	return nil
}

func (s swapHandler) ETH2USDT(ctx context.Context, request *pbSwap.MoneyRequest, response *pbSwap.MoneyResponse) error {
	m, valid := new(big.Rat).SetString(request.Money)
	if !valid || m.Cmp(new(big.Rat).SetInt64(0))==0{
		return errors.New("invalid float number")
	}

	usdtDecmals, err := s.getUSDTDecimals()
	if err != nil{
		return err
	}
	ethDecmals, err := s.getETHDecimals()
	if err != nil{
		return err
	}

	r, err := s.eth2usdtContractIns.GetReserves(nil)
	if err != nil{
		return err
	}
	logger.Infof("eth reserves: %+v", r)
	//1 is USDT，0 is eth
	price := exchangeFormula(r.Reserve0, r.Reserve1, ethDecmals, usdtDecmals)

	if request.Money == "" {
		response.Money = price.FloatString(int(usdtDecmals))
	}else{
		response.Money = new(big.Rat).Mul(price, m).FloatString(int(usdtDecmals))
	}
	return nil
}

func (s swapHandler) ETH2OPS(ctx context.Context, request *pbSwap.MoneyRequest, response *pbSwap.MoneyResponse) error {
	m, valid := new(big.Rat).SetString(request.Money)
	if !valid || m.Cmp(new(big.Rat).SetInt64(0))==0{
		return errors.New("invalid float number")
	}

	opsDecimalsValue, err :=  s.getOPSDecimals()
	if err != nil{
		return err
	}
	usdtDecmals, err := s.getUSDTDecimals()
	if err != nil{
		return err
	}
	ethDecmals, err := s.getETHDecimals()
	if err != nil{
		return err
	}

	opsReserves, err := s.ops2usdtContractIns.GetReserves(nil)
	if err != nil{
		return err
	}
	logger.Infof("eth reserves: %+v", opsReserves)

	//1 is USDT，0 isOps
	opsPrice := exchangeFormula(opsReserves.Reserve0, opsReserves.Reserve1,
		opsDecimalsValue, usdtDecmals)

	ethReserves, err := s.eth2usdtContractIns.GetReserves(nil)
	if err != nil{
		return err
	}
	logger.Infof("eth reserves: %+v", ethReserves)
	//1 is USDT，0 is eth
	ethPrice := exchangeFormula(ethReserves.Reserve0,
		ethReserves.Reserve1, ethDecmals, usdtDecmals)

	logger.Infof("ethPrice=%.18f, opsPrice=%.18f", ethPrice, opsPrice)
	if request.Money == "" {
		price := new(big.Rat).Quo(ethPrice, opsPrice)
		response.Money = price.FloatString(int(opsDecimalsValue))
	}else{
		ethMoney := new(big.Rat).Mul(ethPrice, m)
		response.Money = new(big.Rat).Quo(ethMoney, opsPrice).FloatString(int(opsDecimalsValue))
	}
	return nil
}


