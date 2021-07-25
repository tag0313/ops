package controller

import (
	"context"
	"math/big"
	"net/http"
	"ops/pkg/model/consulreg"
	model "ops/pkg/model/contract"
	myjwt "ops/pkg/model/jwt"
	utils3 "ops/pkg/utils"
	pbContract "ops/proto/contract"
	pbNft1155 "ops/proto/nft1155"
	pbSwap "ops/proto/swap"

	"github.com/asim/go-micro/v3/logger"
	"github.com/gin-gonic/gin"
)

// SubscribeOPSRecharge godoc
// @Summary 订阅用户 OPS 充值信息
// @Description 从后端获取用户的OPS充值确认信息，在交易成功后会推送到前端
// @ID SubscribeOPSRecharge
// @tags OPS
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param SubscribeOPSRecharge body model.SubscribeOPSRechargeReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} model.SubscribeOPSRechargeResponse 成功
// @Failure 4005 {object} model.SubscribeOPSRechargeResponse "微服务可能挂了"
// @Failure 4004 {object} model.SubscribeOPSRechargeResponse "输入参数错误"
// @Router /contract/ops/subscribe-ops-recharge [POST]
func SubscribeOPSRecharge(ctx *gin.Context) {
	var (
		statusCode int
		req        model.SubscribeOPSRechargeReq
		response   model.SubscribeOPSRechargeResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils3.RECODE_DATAERR, "bind data failed in SubscribeOPSRecharge", err)
		statusCode = http.StatusBadRequest
		return
	}
	microClient := pbContract.NewContractService("contract", consulreg.MicroSer.Client())
	pbReq := &pbContract.GetTransactionByHashRequest{
		TransactionHash: req.TransactionHash,
	}
	PBresp, err := microClient.GetTransactionByHash(ctx, pbReq)
	if err != nil {
		return
	}
	//Your code logic here
	response.Data.ToAddress = PBresp.ContractTo
	response.Data.State = int(PBresp.Status)
	response.Data.Amount = PBresp.ContractAmount
	statusCode = http.StatusOK
	response.NewSuccess()

	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	//err = utils.SendTradeMsgPending(PBresp.Hash, toAddress, fromAddress, uid string, amount float64)

	//send notify message
	err = utils3.SendTradeMsgPending(PBresp.Hash, PBresp.To,
		"", claims.Uid, ByteToFloat64([]byte(PBresp.ContractAmount)))
	if err != nil {
		logger.Error(err)
		return
	}
	//claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	//err = utils.SendTradeMsgPending(PBresp.Hash, toAddress, fromAddress, uid string, amount float64)

	//err = utils.SendTradeMsgPending(PBresp.Hash, PBresp.To,
	//	erc20ContractAddr, claims.Uid, -1)
	//if err != nil{
	//	logger.Error(err)
	//	return
	//}
	//response.Data.FromAddress = PBresp.F
}

// GetGasFee godoc
// @Summary 获取 OPS 需要的交易手续费 （单位为：OPS）
// @Description 从后端获取交易手续费，返回两个值，分别是建议值和最高值。（单位为：OPS）
// @ID GetGasFee
// @tags ETH
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param GetGasFee body model.GetGasFeeReq true "to 和 from 地址可以选择一个传空字符串，表示是管理钱包的地址，请勿传 null 或不填。"
// @Produce  json
// @Success 0 {object} model.GetGasFeeResponse 成功
// @Failure 4005 {object} model.GetGasFeeResponse "微服务可能挂了"
// @Failure 4004 {object} model.GetGasFeeResponse "输入参数错误"
// @Router /contract/get-gas-fee [POST]
func GetGasFee(ctx *gin.Context) {
	var (
		statusCode int
		req        model.GetGasFeeReq
		response   model.GetGasFeeResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils3.RECODE_DATAERR, "bind data failed in GetGasFee", err)
		statusCode = http.StatusBadRequest
		return
	}
	//Your code logic here
	opsReq := &pbContract.GetGasFeeRequest{
		FromAddress: req.FromAddress,
		ToAddress:   req.ToAddress,
		Amount:      stringDecimalToBytes(req.Amount),
	}
	microClient := pbContract.NewContractService("contract", consulreg.MicroSer.Client())
	resp, err := microClient.GetGasFee(ctx, opsReq)
	if err != nil {
		response.SetError(utils3.RECODE_MICROERR, "internal server error", err)
		statusCode = http.StatusInternalServerError
		return
	}

	ETH := GasToEth(resp.GasLimit, new(big.Int).SetBytes(resp.GasPrice))
	logger.Infof("eth is: %s, gasLimit=%d gasPrice=%s", ETH.Text('f', 18),
		resp.GasLimit, new(big.Int).SetBytes(resp.GasPrice).String())
	swapClient := pbSwap.NewSwapService("contract", consulreg.MicroSer.Client())
	priceResponse, err := swapClient.ETH2OPS(ctx, &pbSwap.MoneyRequest{Money: ETH.String()})
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils3.RECODE_MICROERR, "call GetCreateBatchPrice failed", err)
		return
	}
	response.Data.Money = priceResponse.Money

	statusCode = http.StatusOK
	response.NewSuccess()
}

// WithdrawOps godoc
// @Summary OPS 提现(仅仅测试使用）。
// @Description 提取 OPS 代币，需传入用户的钱包地址和手续费
// @ID WithdrawOps
// @tags OPS
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param WithdrawOps body model.WithdrawOpsReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} model.WithdrawOpsResponse 成功
// @Failure 4005 {object} model.WithdrawOpsResponse "微服务可能挂了"
// @Failure 4004 {object} model.WithdrawOpsResponse "输入参数错误"
// @Failure 4122 {object} model.WithdrawOpsResponse "手续费过低"
// @Failure 4501 {object} model.WithdrawOpsResponse "未知错误"
// @Router /contract/ops/withdraw [POST]
func WithdrawOps(ctx *gin.Context) {
	var (
		statusCode int
		req        model.WithdrawOpsReq
		response   model.WithdrawOpsResponse
		err        error
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err = ctx.ShouldBind(&req); err != nil {
		response.SetError(utils3.RECODE_DATAERR, "bind data failed in WithdrawOps", err)
		statusCode = http.StatusBadRequest
		return
	}

	pbReq := &pbNft1155.TransferERC20Request{}
	pbReq.AddressTo = req.ToAddress
	pbReq.TokenContract = req.TokenAddress
	pbReq.Amount, err = model.AmountToFloatString(model.AmountDataType(req.AmountDataType))
	if err != nil {
		response.SetError(utils3.RECODE_DATAERR, "bind data failed in WithdrawOps", err)
		statusCode = http.StatusBadRequest
		return
	}

	microClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	resp, err := microClient.TransferERC20(context.TODO(), pbReq)
	if err != nil {
		response.SetError(utils3.RECODE_MICROERR, "internal server error", err)
		statusCode = http.StatusInternalServerError
		return
	}
	logger.Info("http response %+v", resp)
	response.Data.TransactionHash = resp.TransactionHash

	statusCode = http.StatusOK
	response.NewSuccess()
}

// WithDrawGasFee godoc
// @Summary 获取提现手续手续费，参数和/property/withdraw一致，返回提现所需的手续费。
// @Description 获取OPS提现手续费
// @ID WithDrawGasFee
// @tags OPS
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param WithDrawGasFee body model.WithdrawFeeReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} model.GetGasFeeResponse 成功
// @Failure 4005 {object} model.GetGasFeeResponse "微服务可能挂了"
// @Failure 4004 {object} model.GetGasFeeResponse "输入参数错误"
// @Router /contract/ops/withdraw-fee [POST]
func WithDrawGasFee(ctx *gin.Context) {
	var (
		statusCode int
		req        model.WithdrawFeeReq
		response   model.GetGasFeeResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils3.RECODE_DATAERR, "bind data failed in WithDrawGasFee", err)
		statusCode = http.StatusBadRequest
		return
	}
	bfStr, err := stringDecimalToBigFloat(req.Amount)
	if err != nil {
		statusCode = http.StatusBadRequest
		response.SetError(utils3.RECODE_PARAMERR, "convert big float failed", err)
		return
	}
	//Your code logic here
	withDrawRequest := &pbNft1155.TransferERC20Request{
		TokenContract: utils3.GetConfigStr("erc20_contract_address"),
		AddressTo:     req.ToAddress,
		Amount:        bfStr.String(),
	}
	microClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	resp, err := microClient.GetTransferERC20Price(context.TODO(), withDrawRequest)
	if err != nil {
		response.SetError(utils3.RECODE_MICROERR, "GetTransferERC20Price", err)
		statusCode = http.StatusInternalServerError
		return
	}

	ETH := GasToEth(resp.GasLimit, new(big.Int).SetBytes(resp.GasPrice))
	logger.Infof("eth is: %s, gasLimit=%d gasPrice=%s", ETH.Text('f', 18),
		resp.GasLimit, new(big.Int).SetBytes(resp.GasPrice).String())
	swapClient := pbSwap.NewSwapService("contract", consulreg.MicroSer.Client())
	priceResponse, err := swapClient.ETH2OPS(ctx, &pbSwap.MoneyRequest{Money: ETH.String()})
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils3.RECODE_MICROERR, "call GetCreateBatchPrice failed", err)
		return
	}
	response.Data.Money = priceResponse.Money

	statusCode = http.StatusOK
	response.NewSuccess()
}

// OpsBalance godoc
// @Summary 获取 ops 账户余额
// @Description 获取指定 OPS 账户余额
// @ID OpsBalance
// @tags OPS
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param OpsBalance body model.OpsBalanceReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} model.OpsBalanceResponse 成功
// @Failure 4005 {object} model.OpsBalanceResponse "微服务可能挂了"
// @Failure 4004 {object} model.OpsBalanceResponse "输入参数错误"
// @Router /contract/ops/balance [POST]
func OpsBalance(ctx *gin.Context) {
	var (
		statusCode int
		req        model.OpsBalanceReq
		response   model.OpsBalanceResponse
	)
	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils3.RECODE_DATAERR, "bind data failed in OpsBalance", err)
		statusCode = http.StatusBadRequest
		return
	}
	opsReq := &pbContract.BalanceRequest{
		Address: req.Address,
	}
	microClient := pbContract.NewContractService("contract", consulreg.MicroSer.Client())
	balance, err := microClient.BalanceOf(context.TODO(), opsReq)
	if err != nil {
		response.SetError(utils3.RECODE_MICROERR, "call balance of error", err)
		statusCode = http.StatusInternalServerError
	} else {
		response.Data.Balance = intBytesToString(balance.GetBalance())
		response.Data.Address = req.Address
		response.Data.Decimals = balance.GetDecimals()
		response.NewSuccess()
		statusCode = http.StatusOK
	}

	statusCode = http.StatusOK
	response.NewSuccess()
}

// OpsInfo godoc
// @Summary 查询 ERC20 合约的 owner, name, decimals, total supply, symbol
// @Description OpsInfo 查询合约的基本信息
// @ID OpsInfo
// @tags OPS
// @Accept json
// @Produce  json
// @Success 0 {object} model.OpsInfo "返回合约的基本信息"
// @Failure 4005 {object} model.OpsInfo "微服务可能挂了"
// @Router /contract/ops/info [POST]
func OpsInfo(ctx *gin.Context) {
	var (
		response   model.OpsInfo
		statusCode int
	)
	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	microClient := pbContract.NewContractService("contract", consulreg.MicroSer.Client())
	nameResponse, err := microClient.Name(context.TODO(), new(pbContract.NameRequest))
	if err != nil {
		response.SetError(utils3.RECODE_MICROERR, "call Name error", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data.Name = nameResponse.GetName()

	ownerResponse, err := microClient.Owner(context.TODO(), new(pbContract.OwnerRequest))
	if err != nil {
		response.SetError(utils3.RECODE_MICROERR, "call Owner error", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data.Owner = ownerResponse.GetAddress()

	symbolResponse, err := microClient.Symbol(context.TODO(), new(pbContract.SymbolRequest))
	if err != nil {
		response.SetError(utils3.RECODE_MICROERR, "call Symbol error", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data.Symbol = symbolResponse.GetSymbol()

	totalSupplyResponse, err := microClient.TotalSupply(context.TODO(), new(pbContract.TotalSupplyRequest))
	if err != nil {
		response.SetError(utils3.RECODE_MICROERR, "call TotalSupply error", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data.TotalSupply = totalSupplyResponse.GetTotalSupply()

	decimals, err := microClient.Decimals(context.TODO(), new(pbContract.DecimalsRequest))
	if err != nil {
		response.SetError(utils3.RECODE_MICROERR, "call Decimals error", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data.Decimals = decimals.GetDecimals()
	statusCode = http.StatusOK
	response.NewSuccess()
}
