package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"ops/pkg/model/consulreg"
	model "ops/pkg/model/contract"
	"ops/pkg/utils"
	pbEthereum "ops/proto/ethereum"
)

// EthBalance godoc
// @Summary 查询 Ethereum 指定地址余额
// @Description 查询 Ethereum 指定地址余额，单位为 ETH。
// @ID EthBalance
// @tags ETH
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param EthBalance body model.EthBalanceReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} model.EthBalanceResponse 成功
// @Failure 4005 {object} model.EthBalanceResponse "微服务可能挂了"
// @Failure 4004 {object} model.EthBalanceResponse "输入参数错误"
// @Router /contract/balance [POST]
func EthBalance(ctx *gin.Context) {
	var (
		statusCode int
		req        model.EthBalanceReq
		response   model.EthBalanceResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in EthBalance", err)
		statusCode = http.StatusBadRequest
		return
	}
	microClient := pbEthereum.NewEthereumService("contract", consulreg.MicroSer.Client())
	pbReq := new(pbEthereum.ETHBalanceRequest)
	pbReq.Address = req.Address

	pbRep, err := microClient.BalanceETH(context.TODO(), pbReq)
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "call balance of error", err)
		statusCode = http.StatusInternalServerError
	} else {
		response.Data.Balance = intBytesToString(pbRep.Balance)
		response.Data.Address = req.Address
		response.Data.Decimals = pbRep.Decimals
		response.NewSuccess()
		statusCode = http.StatusOK
	}
}
