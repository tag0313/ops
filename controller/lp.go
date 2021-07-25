package controller

import (
	"errors"
	"net/http"
	modelLp "ops/pkg/model/lp"
	"ops/pkg/utils"
	pbContract "ops/proto/contract"

	"github.com/gin-gonic/gin"
)

func convertLpType(lpType modelLp.LpType) (pbContract.LpType, error) {
	_, ok := pbContract.LpType_name[int32(lpType)]
	if !ok {
		return 0, errors.New("no such Liquidity Provided Type")
	}
	return pbContract.LpType(lpType), nil
}

func getLpTypeStr(lpType modelLp.LpType) string {
	str := pbContract.LpType_name[int32(lpType)]
	//if !ok{
	//	return "", errors.New("no such Liquidity Provided Type")
	//}
	return str
}

// GetOpsUsdtApy godoc
// @Summary 查询 ops usdt 矿池 apy, 获取到参数后需要自行计算百分比。
// @Description
// @ID GetOpsUsdtApy
// @tags 流动挖矿
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param GetOpsUsdtApy body modelLp.GetOpsUsdtApyReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} modelLp.GetOpsUsdtApyResponse 成功
// @Failure 4005 {object} modelLp.GetOpsUsdtApyResponse "微服务可能挂了"
// @Failure 4004 {object} modelLp.GetOpsUsdtApyResponse "输入参数错误"
// @Router /lp/get-ops-usdt-apy [POST]
func GetOpsUsdtApy(ctx *gin.Context) {
	var (
		statusCode int
		//req modelLp.GetOpsUsdtApyReq
		response modelLp.GetOpsUsdtApyResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	//if err := ctx.ShouldBind(&req); err != nil {
	//	response.SetError(utils.RECODE_DATAERR, "bind data failed in GetOpsUsdtApy", err )
	//	statusCode = http.StatusBadRequest
	//	return
	//}
	//Your code logic here
	pbResp, err := RpcLpService.GetOpsUsdtApy(ctx, new(pbContract.GetOpsUsdtApyRequest))
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "bind data failed in GetOpsUsdtApy", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data.Apy = pbResp.Capitalization

	statusCode = http.StatusOK
	response.NewSuccess()
}

// GetOpsFluxApy godoc
// @Summary 查询 ops flux 矿池 apy, 获取到参数后需要自行计算百分比。
// @Description
// @ID GetOpsFluxApy
// @tags 流动挖矿
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param GetOpsFluxApy body modelLp.GetOpsFluxApyReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} modelLp.GetOpsFluxApyResponse 成功
// @Failure 4005 {object} modelLp.GetOpsFluxApyResponse "微服务可能挂了"
// @Failure 4004 {object} modelLp.GetOpsFluxApyResponse "输入参数错误"
// @Router /lp/get-ops-flux-apy [POST]
func GetOpsFluxApy(ctx *gin.Context) {
	var (
		statusCode int
		//req modelLp.GetOpsFluxApyReq
		response modelLp.GetOpsFluxApyResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	//if err := ctx.ShouldBind(&req); err != nil {
	//	response.SetError(utils.RECODE_DATAERR, "bind data failed in GetOpsFluxApy", err )
	//	statusCode = http.StatusBadRequest
	//	return
	//}
	//Your code logic here
	pbResp, err := RpcLpService.GetOpsFluxAPY(ctx, new(pbContract.GetOpsFluxApyRequest))
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "bind data failed in GetOpsFluxApy", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data.Apy = pbResp.Capitalization

	statusCode = http.StatusOK
	response.NewSuccess()
}

// GetOpsPriceUsdt godoc
// @Summary 查询 ops 的价格，以 usdt 表示
// @Description
// @ID GetOpsPriceUsdt
// @tags 流动挖矿
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param GetOpsPriceUsdt body modelLp.GetOpsPriceUsdtReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} modelLp.GetOpsPriceUsdtResponse 成功
// @Failure 4005 {object} modelLp.GetOpsPriceUsdtResponse "微服务可能挂了"
// @Failure 4004 {object} modelLp.GetOpsPriceUsdtResponse "输入参数错误"
// @Router /lp/get-ops-price-usdt [POST]
func GetOpsPriceUsdt(ctx *gin.Context) {
	var (
		statusCode int
		//req modelLp.GetOpsPriceUsdtReq
		response modelLp.GetOpsPriceUsdtResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	//if err := ctx.ShouldBind(&req); err != nil {
	//	response.SetError(utils.RECODE_DATAERR, "bind data failed in GetOpsPriceUsdt", err )
	//	statusCode = http.StatusBadRequest
	//	return
	//}
	//Your code logic here
	pbResp, err := RpcLpService.GetOpsPriceUsdt(ctx, new(pbContract.GetOpsPriceRequest))
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "bind data failed in GetOpsFluxApy", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data.Price = pbResp.Wroth

	statusCode = http.StatusOK
	response.NewSuccess()
}

// GetMiningPoolWorthUsdt godoc
// @Summary 查询对应矿池的价值，以 usdt 表示.
// @Description 请求参数pool_type 0 代表 LP（OPS-USDT）; 1 代表（OPS-FLUX）
// @ID GetMiningPoolWorthUsdt
// @tags 流动挖矿
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param GetMiningPoolWorthUsdt body modelLp.GetMiningPoolWorthUsdtReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} modelLp.GetMiningPoolWorthUsdtResponse 成功
// @Failure 4005 {object} modelLp.GetMiningPoolWorthUsdtResponse "微服务可能挂了"
// @Failure 4004 {object} modelLp.GetMiningPoolWorthUsdtResponse "输入参数错误"
// @Router /lp/get-mining-pool-worth-usdt [POST]
func GetMiningPoolWorthUsdt(ctx *gin.Context) {
	var (
		statusCode int
		req        modelLp.GetMiningPoolWorthUsdtReq
		response   modelLp.GetMiningPoolWorthUsdtResponse
		err        error
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in GetMiningPoolWorthUsdt", err)
		statusCode = http.StatusBadRequest
		return
	}
	pbReq := &pbContract.GetPoolWorthRequest{}
	pbReq.LpType, err = convertLpType(req.PoolType)
	if err != nil {
		response.SetError(utils.RECODE_DATAERR, "no such Liquidity Provided Type", err)
		statusCode = http.StatusBadRequest
		return
	}
	pbResp, err := RpcLpService.GetPoolWorth(ctx, pbReq)
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "no such Liquidity Provided Type", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data = modelLp.GetMiningPoolWorthUsdtData{
		Worth:       pbResp.GetWroth(),
		PoolType:    req.PoolType,
		PoolTypeStr: getLpTypeStr(req.PoolType),
	}

	statusCode = http.StatusOK
	response.NewSuccess()
}

// GetOpsPriceFlux godoc
// @Summary 查询对应 OPS 的价格，以 flux 表示
// @Description
// @ID GetOpsPriceFlux
// @tags 流动挖矿
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param GetOpsPriceFlux body modelLp.GetOpsPriceFluxReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} modelLp.GetOpsPriceFluxResponse 成功
// @Failure 4005 {object} modelLp.GetOpsPriceFluxResponse "微服务可能挂了"
// @Failure 4004 {object} modelLp.GetOpsPriceFluxResponse "输入参数错误"
// @Router /lp/get-ops-price-flux [POST]
func GetOpsPriceFlux(ctx *gin.Context) {
	var (
		statusCode int
		//req modelLp.GetOpsPriceFluxReq
		response modelLp.GetOpsPriceFluxResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	//if err := ctx.ShouldBind(&req); err != nil {
	//	response.SetError(utils.RECODE_DATAERR, "bind data failed in GetOpsPriceFlux", err )
	//	statusCode = http.StatusBadRequest
	//	return
	//}
	//Your code logic here
	pbResp, err := RpcLpService.GetOpsPriceFlux(ctx, new(pbContract.GetOpsPriceRequest))
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "calling rpc GetOpsPriceFlux failed", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data.Price = pbResp.Wroth

	statusCode = http.StatusOK
	response.NewSuccess()
}

// GetUserLPInfo godoc
// @Summary 查询对应矿池用户地址的余额(以 USDT 表示)，lp数量，ops 奖励。
// @Description 请求参数pool_type 0 代表 LP（OPS-USDT）; 1 代表（OPS-FLUX）
// @ID GetUserLPInfo
// @tags 流动挖矿
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param GetUserLPInfo body modelLp.GetUserLPInfo true "传入相关参数"
// @Produce  json
// @Success 0 {object} modelLp.GetUserLPInfoResponse 成功
// @Failure 4005 {object} modelLp.GetUserLPInfoResponse "微服务可能挂了"
// @Failure 4004 {object} modelLp.GetUserLPInfoResponse "输入参数错误"
// @Router /lp/get-user-lp-info [POST]
func GetUserLPInfo(ctx *gin.Context) {
	var (
		statusCode int
		err        error
		req        modelLp.GetUserLPInfo
		response   modelLp.GetUserLPInfoResponse

		balanceRep = new(pbContract.BalanceResponse)
		amountRep  = new(pbContract.LpAmountResponse)
		rewardRep  = new(pbContract.GetOpsRewardsResponse)
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in GetUserLPInfo", err)
		statusCode = http.StatusBadRequest
		return
	}
	//Your code logic here

	pbReq := &pbContract.UserLpRequest{
		AccountAddress: req.Address,
	}
	pbReq.LpType, err = convertLpType(req.PoolType)
	if err != nil {
		response.SetError(utils.RECODE_DATAERR, "no such Liquidity Provided Type", err)
		statusCode = http.StatusBadRequest
		return
	}

	balanceRep, err = RpcLpService.GetUserLpBalance(ctx, pbReq)
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "calling rpc GetUserLpBalance failed", err)
		statusCode = http.StatusInternalServerError
		return
	}

	amountRep, err = RpcLpService.GetUserLpAmount(ctx, pbReq)
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "calling rpc GetUserLpAmount failed", err)
		statusCode = http.StatusInternalServerError
		return
	}

	rewardRep, err = RpcLpService.GetOpsReward(ctx, pbReq)
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "calling rpc GetOpsReward failed", err)
		statusCode = http.StatusInternalServerError
		return
	}

	response.Data = modelLp.GetUserLPInfoData{
		Balance:     balanceRep.BalanceStr,
		Amount:      amountRep.Amount,
		OpsEarned:   rewardRep.Rewards,
		PoolType:    req.PoolType,
		PoolTypeStr: getLpTypeStr(req.PoolType),
	}

	statusCode = http.StatusOK
	response.NewSuccess()
}
