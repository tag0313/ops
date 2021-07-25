package controller

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"ops/pkg/model"
	"ops/pkg/model/consulreg"
	contractModel "ops/pkg/model/contract"
	myjwt "ops/pkg/model/jwt"
	"ops/pkg/model/property"
	utils2 "ops/pkg/utils"
	pbContract "ops/proto/contract"
	pbFollower "ops/proto/follower"
	pbMessage "ops/proto/message"
	pbNft1155 "ops/proto/nft1155"
	pbProperty "ops/proto/property"
	pbSwap "ops/proto/swap"
	pbUserInfo "ops/proto/userInfo"
	"strconv"
	"time"
)

// DepositOpspoint godoc
// @Summary 充值 ops 到 opspoint 。
// @Description 当调用/property/deposit接口时，此方法会将用户的ops 转为opspoint并存入mongo
// @ID DepositOpspoint
// @tags 用户资产
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid"
// @Success 0 {object} property.OpsPointResp "返回当前用户的ops point 和uid"
// @Header 200 {string} token "token"
// @Failure 4002 {object} property.OpsPointResp "没有在mongodb里边找到对应的用户资产"
// @Failure 4005 {object} property.OpsPointResp "微服务可能挂了"
// @Router /property/deposit [post]
func DepositOpspoint(ctx *gin.Context) {
	var (
		statusCode int
		req        contractModel.SubscribeOPSRechargeReq
		response   contractModel.SubscribeOPSRechargeResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils2.RECODE_DATAERR, "bind data failed in SubscribeOPSRecharge", err)
		statusCode = http.StatusBadRequest
		return
	}
	microClient := pbContract.NewContractService("contract", consulreg.MicroSer.Client())
	pbReq := &pbContract.GetTransactionByHashRequest{
		TransactionHash: req.TransactionHash,
	}
	remote, err := microClient.GetTransactionByHash(ctx, pbReq)
	if err != nil {
		return
	}
	fmt.Println(remote.IsPending)
	fmt.Println(remote)
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	deposit64 := ByteToFloat64([]byte(remote.ContractAmount))
	if remote.IsPending {
		// pending 未确认成功还是失败
		err = utils2.SendTradeMsgPending(remote.Hash, remote.To, "remote.fromAddress", claims.Uid, deposit64)
	} else {
		if remote.Status == 0 {
			// 失败
			err = utils2.SendTradeMsgFailed(remote.Hash, remote.To, "remote.fromAddress", claims.Uid, deposit64)
		} else if remote.Status == 1 {
			// 成功
			err = utils2.SendTradeMsgSuccess(remote.Hash, remote.To, "remote.fromAddress", claims.Uid, deposit64)
		}
	}
	if err != nil {
		logger.Info(err)
	}

	//cache the uid and transaction hash into mongo transfer_erc20_history

	statusCode = http.StatusOK
	response.NewSuccess()
}

// WithdrawOpspoint godoc
// @Summary 提现opspoint 到ops 。
// @Description 当调用/property/withdraw接口时，此方法会将用户的opspoint转为ops并存入mongo
// @ID WithDrawOpspoint
// @tags 用户资产
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid"
// @Param WithdrawOpspoint body contractModel.WithdrawOpsReq true "header中的token自带了uid"
// @Success 0 {object} contractModel.WithdrawOpsResponse "返回当前用户的ops point 和uid"
// @Header 200 {string} token "token"
// @Failure 4002 {object} contractModel.WithdrawOpsResponse "没有在mongodb里边找到对应的用户资产"
// @Failure 4005 {object} contractModel.WithdrawOpsResponse "微服务可能挂了"
// @Router /property/withdraw [post]
func WithdrawOpspoint(ctx *gin.Context) {
	var (
		statusCode        int
		req               contractModel.WithdrawOpsReq
		response          contractModel.WithdrawOpsResponse
		erc20ContractAddr = utils2.GetConfigStr("erc20_contract_address")
		claims            = ctx.MustGet("claims").(*myjwt.CustomClaims)
	)
	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils2.RECODE_DATAERR, "bind data failed in WithDrawOpspoint", err)
		statusCode = http.StatusBadRequest
		return
	}

	reqGasfee, err := stringDecimalToBigFloat(req.GasFee)
	if err != nil {
		response.SetError(utils2.RECODE_DATAERR, "stringDecimalToBigFloat", err)
		statusCode = http.StatusBadRequest
		return
	}
	fAmount, err := stringDecimalToBigFloat(req.Amount)
	if err != nil {
		response.SetError(utils2.RECODE_DATAERR, "stringDecimalToBigFloat", err)
		statusCode = http.StatusBadRequest
		return
	}

	//check the opspoint that it enough for withdraw
	withdrawTotal := new(big.Float).Add(fAmount, reqGasfee)
	withdrawF64, _ := withdrawTotal.Float64()
	point, err := GetOpsPoint(claims.Uid, withdrawF64)
	if err != nil {
		response.SetError(utils2.RECODE_MICROERR, utils2.RecodeTest(utils2.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if point == utils2.RECODE_INSUFFICIENT_FUND {
		response.SetError(utils2.RECODE_INSUFFICIENT_FUND, utils2.RecodeTest(utils2.RECODE_INSUFFICIENT_FUND), err)
		statusCode = http.StatusBadRequest
		return
	}

	//compare the gasfee which is come from frontend with contract, the result is not great than %0.5
	pbGetGasfeeReq := &pbNft1155.TransferERC20Request{}
	pbGetGasfeeReq.TokenContract = erc20ContractAddr
	pbGetGasfeeReq.AddressTo = req.ToAddress
	pbGetGasfeeReq.Amount = fAmount.String()
	logger.Infof("getGasFeeReq=%+v  amount=%.18f", pbGetGasfeeReq, fAmount)
	pbGetGasfeeResp, err := RpcNft1155Service.GetTransferERC20Price(context.TODO(), pbGetGasfeeReq)
	if err != nil {
		response.SetError(utils2.RECODE_MICROERR, utils2.RecodeTest(utils2.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	}

	ETH := GasToEth(pbGetGasfeeResp.GasLimit, new(big.Int).SetBytes(pbGetGasfeeResp.GasPrice))
	logger.Infof("eth is: %s, gasLimit=%d gasPrice=%s", ETH.Text('f', 18),
		pbGetGasfeeResp.GasLimit, new(big.Int).SetBytes(pbGetGasfeeResp.GasPrice).String())
	priceResponse, err := RpcSwapService.ETH2OPS(ctx, &pbSwap.MoneyRequest{Money: ETH.String()})
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call ETH2OPS failed", err)
		return
	}
	logger.Info("eth to ops price is: ", priceResponse.Money)
	priceResponseGasfee, err := stringDecimalToBigFloat(priceResponse.Money)
	if err != nil {
		response.SetError(utils2.RECODE_DATAERR, utils2.RecodeTest(utils2.RECODE_DATAERR), err)
		statusCode = http.StatusBadRequest
		return
	}

	rate := new(big.Float).SetFloat64(0.005)
	if err = checkFloatGreatRate(reqGasfee, priceResponseGasfee, rate); err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_WITHDRAW_GASFEE, "checkFloatGreatRate", err)
		return
	}

	decOpsReq := &pbProperty.IncDecOpsPointReq{
		Op:       pbProperty.OpsIncDec_Dec,
		OpsPoint: withdrawF64,
		Uid:      claims.Uid,
	}
	decOpsResult, err := RpcPropertyService.IncDecOpsPoint(ctx, decOpsReq)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "checkFloatGreatRate", err)
		return
	}
	logger.Infof("user %s balance is: %f", claims.Uid, decOpsResult.OpsBalance)

	//call contract server withdraw opspoint
	pbReq := &pbNft1155.TransferERC20Request{}
	pbReq.AddressTo = req.ToAddress
	pbReq.TokenContract = erc20ContractAddr
	pbReq.Amount = fAmount.String()
	logger.Infof("TransferERC20=%+v, amount=%.18f", pbReq, fAmount)
	resp, err := RpcNft1155Service.TransferERC20(context.TODO(), pbReq)
	if err != nil {
		response.SetError(utils2.RECODE_MICROERR, "internal server error", err)
		statusCode = http.StatusInternalServerError
		return
	}
	logger.Info("http response %+v", resp)

	opspoint64, _ := fAmount.Float64()
	//if transaction hash is nil the failure message should be send into message
	logger.Info("respCreat1155: ", resp)
	if resp.TransactionHash == "" {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call TransferERC20 failed", err)
		// sent the notification to message service 交易失败
		err = utils2.SendTradeMsgFailed(resp.TransactionHash, req.ToAddress,
			erc20ContractAddr, claims.Uid, opspoint64)
		if err != nil {
			logger.Error(err, "消息进入redis队列失败")
		}

		//链上交易提交失败了 OPS point 加回去
		decOpsReq.Op = pbProperty.OpsIncDec_Inc
		decOpsResult, err = RpcPropertyService.IncDecOpsPoint(ctx, decOpsReq)
		if err != nil {
			statusCode = http.StatusInternalServerError
			response.SetError(utils2.RECODE_MICROERR, "checkFloatGreatRate", err)
			return
		}
		logger.Infof("user %s balance is: %f", claims.Uid, decOpsResult.OpsBalance)
		return
	}
	//success: send the message to the message service
	//the trade is pending
	err = utils2.SendTradeMsgPending(resp.TransactionHash, req.ToAddress,
		erc20ContractAddr, claims.Uid, opspoint64)
	if err != nil {
		logger.Error(err, "消息进入redis队列失败")
	}

	////cache those transaction information into mongo, this operation will creat a new collection, maybe called withdraw_opspoint_history
	gasfee64, _ := reqGasfee.Float64()
	pbWithdrawHistory := &pbProperty.WithdrawOpspoint{}
	pbWithdrawHistory.TransactionHash = resp.TransactionHash
	pbWithdrawHistory.Uid = claims.Uid
	pbWithdrawHistory.Opspoint = opspoint64
	pbWithdrawHistory.Gasfee = gasfee64
	resultStoreWithdrawHistory, err := RpcPropertyService.StoreWithdrawHistory(context.TODO(),
		pbWithdrawHistory)
	if err != nil {
		response.SetError(utils2.RECODE_MICROERR, utils2.RecodeTest(utils2.RECODE_MICROERR), err)
		statusCode = http.StatusInternalServerError
		return
	} else if resultStoreWithdrawHistory.Code != utils2.RECODE_OK {
		response.SetError(resultStoreWithdrawHistory.Code,
			utils2.RecodeTest(resultStoreWithdrawHistory.Code), err)
		statusCode = http.StatusInternalServerError
		return
	}
	//get the notification from contract to judge that the transaction is success or failed

	//failed: roll back the opspoint and return the transaction is failed
	response.Data.TransactionHash = resp.GetTransactionHash()
	statusCode = http.StatusOK
	response.NewSuccess()
}

// CheckOpsPoint godoc
// @Summary 用户查看自己的OpsPoint。
// @Description 当调用CheckOpsPoint接口时，此方法会返回当前用户的ops point
// @ID CheckOpsPoint
// @tags 用户资产
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，"
// @Success 0 {object} property.OpsPointResp "返回当前用户的ops point 和uid"
// @Header 200 {string} token "token"
// @Failure 4002 {object} property.OpsPointResp "没有在mongodb里边找到对应的用户资产"
// @Failure 4005 {object} property.OpsPointResp "微服务可能挂了"
// @Router /property/opspoint [post]
func CheckOpsPoint(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbProperty.OpsPoint{}
	resp := &property.OpsPointResp{}
	req.Uid = claims.Uid
	microClient := pbProperty.NewOperatePropertyService("property", consulreg.MicroSer.Client())
	remote, err := microClient.CheckOpsPoint(context.TODO(), req)
	if err != nil {
		logger.Error("The remote server may have some problem", utils2.GetConfigStr("mirco.addr"), " :CheckOpsPoint")
		logger.Error(err)
		ctx.JSON(http.StatusOK, resp.NewError(utils2.RECODE_MICROERR))
		return
	} else if remote.Uid != claims.Uid {
		logger.Error("Error of remote", utils2.RecodeTest(remote.Uid))
		ctx.JSON(http.StatusOK, resp.NewError(remote.Uid))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote))
}

// GetBKGImage godoc
// @Summary 重定向oCard到blob，并且返回图片的byte数据。
// @Description 当调用CheckOpsPoint接口时，此方法会重定向到oCard在容器中的图像数据。
// @ID GetBKGImage
// @tags 用户资产
// @Accept  json
// @Produce  json
// @Param style-code path string true "图片的code，例如1302.png"
// @Success 0 {string} string "返回图片的byte"
// @Header 200 {string} token "token"
// @Router /style/ [get]
func GetBKGImage(ctx *gin.Context) {
	code := ctx.Param("style-code")
	ctx.Redirect(http.StatusMovedPermanently, "https://storage.opsnft.net/"+code)
}

// QueryMintedOCard godoc
// @Summary 查询用户Mint了那些OCard。
// @Description 当调用QueryMintedOCard接口时，此方法会返回用户Mint的卡。
// @ID QueryMintedOCard
// @tags 用户资产
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid"
// @Success 0 {object} property.OCard "返回成功代码，已经铸造ocard的数组"
// @Failure 4004 {object} model.JSONResult "数据序列化出错"
// @Failure 4005 {object} model.JSONResult "micro服务端异常"
// @Failure 4007 {object} model.JSONResult "mongodb里边没数据"
// @Header 200 {string} token "token"
// @Router /property/ocard/query-mint [post]
func QueryMintedOCard(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbProperty.QueryField{}
	resp := &property.OCard{}
	_ = model.JSONResult{} //for swag init
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("bind data failed in QueryOop", err.Error())
		//ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		//return
		req.Uid = claims.Uid
	} else if req.Uid == "" {
		req.Uid = claims.Uid
	}

	microClient := pbProperty.NewOperatePropertyService("property", consulreg.MicroSer.Client())
	remote, err := microClient.QueryMintedOCard(context.TODO(), req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.NewError(utils2.RECODE_MICROERR))
		return
	} else if remote.Code != utils2.RECODE_OK {
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote.Data))
}

// QueryOCardOnOps godoc
// @Summary 查询用户在ops上买了那些ocard。
// @Description 当调用QueryOCardOnOps接口时，此方法会返回在Ops上，用户买的卡。
// @ID QueryOCardOnOps
// @tags 用户资产
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，"
// @Success 0 {object} property.OCardOnOps "返回成功代码，ops上用户拥有多少OCard，不包含他自己mint的卡"
// @Failure 4004 {object} model.JSONResult "数据序列化出错"
// @Failure 4005 {object} model.JSONResult "micro服务端异常"
// @Failure 4007 {object} model.JSONResult "mongodb里边没数据"
// @Header 200 {string} token "token"
// @Router /property/ocard/query-ops [post]
func QueryOCardOnOps(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbProperty.QueryField{}
	resp := &property.OCardOnOps{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("bind data failed in QueryOCardOnOps", err.Error())
		//ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		//return
		//req.Uid = claims.Uid
		req.Uid = claims.Uid
	} else if req.Uid == "" {
		req.Uid = claims.Uid
	}

	microClient := pbProperty.NewOperatePropertyService("property", consulreg.MicroSer.Client())
	remote, err := microClient.QueryOCardFromLocal(context.TODO(), req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.NewError(utils2.RECODE_MICROERR))
		return
	} else if remote.Code != utils2.RECODE_OK {
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote.PurchaseInfo))
}

// QueryOCardOnChain godoc
// @Summary 查询用户在chain上有那些OCard。
// @Description 当调用QueryOCardOnChain接口时，此方法会返回在chain上，用户拥有的卡。
// @ID QueryOCardOnChain
// @tags 用户资产
// @Accept  json
// @Produce  json
// @Param token header string true "header中的token自带了uid，"
// @Success 0 {object} property.OCard "返回成功代码，Chain上用户拥有多少OCard，不包含他自己mint的卡"
// @Failure 4002 {object} model.JSONResult "mongodb里边没数据"
// @Failure 4004 {object} model.JSONResult "数据序列化出错"
// @Failure 4501 {object} model.JSONResult "micro服务端异常"
// @Header 200 {string} token "token"
// @Router /property/ocard/chain [post]
func QueryOCardOnChain(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	req := &pbProperty.QueryField{}
	resp := &property.OCardOnOps{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Error("bind data failed in QueryOCardOnChain", err.Error())
		//ctx.JSON(http.StatusOK, resp.NewError(utils.RECODE_DATAERR))
		//return
		req.Uid = claims.Uid
	} else if req.Uid == "" {
		req.Uid = claims.Uid
	}
	microClient := pbProperty.NewOperatePropertyService("property", consulreg.MicroSer.Client())
	remote, err := microClient.QueryOCardFromChain(context.TODO(), req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.NewError(utils2.RECODE_MICROERR))
		return
	} else if remote.Code != utils2.RECODE_OK {
		ctx.JSON(http.StatusOK, resp.NewError(remote.Code))
		return
	}
	ctx.JSON(http.StatusOK, resp.NewSuccess(remote.PurchaseInfo))
}

// MintBatchCard godoc
// @Summary 批量生成新的 card
// @Description 批量生成新的 ocard
// @ID MintBatchCard
// @tags 用户资产
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param MintBatchCard body pbProperty.MintCard true "传入相关参数"
// @Produce  json
// @Success 0 {object} property.MintCardResponse 成功
// @Failure 4005 {object} property.MintCardResponse "微服务可能挂了"
// @Failure 4004 {object} property.MintCardResponse "输入参数错误"
// @Router /property/ocard/mint-batch [POST]
//data format {"info":[{"ocard":"1302","amount":"12"},{"ocard":"1302","amount":"12"},{"ocard":"1302","amount":"12"}],"gasfee":"0.000000"}
func MintBatchCard(ctx *gin.Context) {
	var (
		statusCode int
		req        property.MintCardReq
		response   property.MintCardResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()
	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils2.RECODE_DATAERR, "bind data failed in MintCard", err)
		statusCode = http.StatusBadRequest
		return
	}
	//Your code logic here
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	jsonIdReq := pbProperty.JsonIdAndGroupIds{}.JsonIdAndGroupId
	jsonIdAndGroupIds := make([]pbProperty.JsonIdAndGroupId, len(req.Info))
	uris := make([]string, len(req.Info))
	amount := make([][]byte, len(req.Info))

	//generate uris
	ocardsInfo := &pbProperty.OCards{}
	ocardInfo := make([]pbProperty.OCard, len(req.Info))
	var totalAmount int64
	var costOfMintOCard float64
	for index, info := range req.Info {
		jsonIdAndGroupIds[index].JsonId = utils2.NewJsonId()
		jsonIdReq = append(jsonIdReq, &jsonIdAndGroupIds[index])
		uris[index] = utils2.GetConfigStr("json_prefix") + jsonIdAndGroupIds[index].JsonId + ".json"
		amount[index] = stringDecimalToBytes(info.Amount)
		ocardInfo[index].Uid = claims.Uid
		ocardInfo[index].CardType = info.Ocard
		ocardInfo[index].Amount = stringDecimalToBytes(info.Amount)
		ocardInfo[index].JsonId = jsonIdAndGroupIds[index].JsonId
		amount, err := strconv.ParseInt(info.Amount, 10, 64)
		if err == nil {
			totalAmount += amount
		}
		ocardsInfo.Ocard = append(ocardsInfo.Ocard, &ocardInfo[index])
	}

	//compare the gas fee using the fee from frontend and contract
	pbGetGasfeeReq := &pbNft1155.CreateBatchRequest{}
	pbGetGasfeeReq.Uris = uris
	pbGetGasfeeReq.Data = []byte(nil)
	pbGetGasfeeReq.Quantities = amount
	pbGetGasfeeReq.InitOwnerAddress = utils2.GetConfigStr("owner_address")
	logger.Infof("getGasFeeReq=%+v  amount=%.18f", pbGetGasfeeReq, amount)
	pbGetGasfeeResp, err := RpcNft1155Service.GetCreateBatchPrice(context.TODO(), pbGetGasfeeReq)
	if err != nil {
		response.SetError(utils2.RECODE_MICROERR, utils2.RecodeTest(utils2.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	}

	ETH := GasToEth(pbGetGasfeeResp.GasLimit, new(big.Int).SetBytes(pbGetGasfeeResp.GasPrice))
	logger.Infof("eth is: %s, gasLimit=%d gasPrice=%s", ETH.Text('f', 18),
		pbGetGasfeeResp.GasLimit, new(big.Int).SetBytes(pbGetGasfeeResp.GasPrice).String())
	priceResponse, err := RpcSwapService.ETH2OPS(ctx, &pbSwap.MoneyRequest{Money: ETH.String()})
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call ETH2OPS failed", err)
		return
	}
	logger.Info("eth to ops price is: ", priceResponse.Money)
	priceResponseGasfee, err := stringDecimalToBigFloat(priceResponse.Money)
	if err != nil {
		response.SetError(utils2.RECODE_DATAERR, utils2.RecodeTest(utils2.RECODE_DATAERR), err)
		statusCode = http.StatusBadRequest
		return
	}

	rate := new(big.Float).SetFloat64(0.005)
	gasfeeFromClient, err := stringDecimalToBigFloat(req.Gasfee)
	if err != nil {
		response.SetError(utils2.RECODE_DATAERR, utils2.RecodeTest(utils2.RECODE_DATAERR), err)
		statusCode = http.StatusBadRequest
		return
	}

	if err = checkFloatGreatRate(gasfeeFromClient, priceResponseGasfee, rate); err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_WITHDRAW_GASFEE, "checkFloatGreatRate", err)
		return
	}

	//verfiey the total amount (the release card) is beyond 10000
	getTotalAmountClient := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
	info, err := getTotalAmountClient.QueryUserInfo(context.TODO(), &pbUserInfo.QueryAndDelete{Uid: claims.Uid})
	releaseTotalAmount := utils2.GetConfigInt64("total_release_amount")
	if err != nil {
		response.SetError(utils2.RECODE_MICROERR, "bind data failed in MintCard", err)
		statusCode = http.StatusBadRequest
		return
	} else if balance := info.TotalAmount + totalAmount - releaseTotalAmount; balance >= 0 {
		afterAbsBalance := strconv.FormatInt(abs(balance), 10)
		response.SetError(utils2.RECODE_AMOUNT, utils2.RecodeTest(utils2.RECODE_AMOUNT)+"当前余额:"+afterAbsBalance, err)
		statusCode = http.StatusBadRequest
		return
	}

	//get mint info
	contractClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(pbNft1155.CreateBatchRequest)
	pbRequest.Data = []byte(nil)
	pbRequest.Uris = uris
	pbRequest.Quantities = amount
	pbRequest.InitOwnerAddress = utils2.GetConfigStr("owner_address")

	respCreat1155, err := contractClient.CreateBatch(ctx, pbRequest)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call NFT1155CreateBatch failed", err)
		return
	}
	logger.Info("respCreat1155: ", respCreat1155)
	if respCreat1155.TransactionHash == "" {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call StoreMintOCardInfo failed", err)
		// sent the notification to message service 生成卡失败
		msg := &pbMessage.Msg{
			MsgType: 8,
			Uid:     claims.Uid,
			CreateCard: &pbMessage.CreateCard{
				TxnHash:   respCreat1155.TransactionHash,
				TxnStatus: "failed",
				Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
				Amount:    totalAmount,
			},
		}
		messageClient := pbMessage.NewOperateMessageService("message", consulreg.MicroSer.Client())
		_, err = messageClient.PutOne(context.TODO(), msg)
		if err != nil {
			logger.Error(err, "消息进入redis队列失败")
		}
		return
	}

	//operate the opspoint by gasfee
	f, _ := gasfeeFromClient.Float64()
	costOfMintOCard = float64(totalAmount)*0.1 + f
	opspoint, err := OperateOpspoint(claims.Uid, -costOfMintOCard)
	if err != nil {
		logger.Error("can't find the remote service in consul", utils2.GetConfigStr("mirco.addr"))
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call MintBatchCard failed", err)
		return
	} else if opspoint != utils2.RECODE_OK {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call MintBatchCard failed", err)
		return
	}

	//store mint card info to mongodb
	storeMintInfo := pbProperty.NewOperatePropertyService("property", consulreg.MicroSer.Client())
	for _, value := range ocardsInfo.Ocard {
		value.TransactionHash = respCreat1155.TransactionHash
	}
	storeResult, err := storeMintInfo.StoreMintOCardsInfo(context.TODO(), ocardsInfo)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call StoreMintOCardInfo failed", err)
		return
	} else if storeResult.Code != utils2.RECODE_OK {
		statusCode = http.StatusInternalServerError
		response.SetError(storeResult.Code, "call StoreMintOCardInfo failed", err)
		return
	}

	// sent the notification to message service 交易生成 等待
	msg := &pbMessage.Msg{
		MsgType: 8,
		Uid:     claims.Uid,
		CreateCard: &pbMessage.CreateCard{
			TxnHash:   respCreat1155.TransactionHash,
			TxnStatus: "pending",
			Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
			Amount:    totalAmount,
		},
	}
	messageClient := pbMessage.NewOperateMessageService("message", consulreg.MicroSer.Client())
	_, err = messageClient.PutOne(context.TODO(), msg)
	if err != nil {
		logger.Error(err, "消息进入redis队列失败")
	}

	statusCode = http.StatusOK
	response.NewSuccess()
}

// BuyOCardOnOps godoc
// @Summary ops上的卡交易
// @Description ops上的卡交易
// @ID BuyOCardOnOps
// @tags 用户资产
// @Accept json
// @Param token header string true "header中的token自带了buyer_uid"
// @Param BuyOCardOnOpsReq body pbProperty.BuyOCardOnOpsReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} property.MintCardResponse 成功
// @Failure 4005 {object} property.MintCardResponse "微服务可能挂了"
// @Failure 4004 {object} property.MintCardResponse "输入参数错误"
// @Failure 4301 {object} property.MintCardResponse "用户余额不足"
// @Failure 4006 {object} property.MintCardResponse "数据存储失败"
// @Failure 4001 {object} property.MintCardResponse "数据库错误"
// @Router /property/ocard/buy-ops [POST]
func BuyOCardOnOps(ctx *gin.Context) {
	var (
		statusCode       int
		buyOCardOnOpsReq *pbProperty.BuyOCardOnOpsReq
		response         property.MintCardResponse
	)
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	//bind the data from frontend
	if err := ctx.ShouldBind(&buyOCardOnOpsReq); err != nil {
		response.SetError(utils2.RECODE_DATAERR, "bind data failed in MintCard", err)
		statusCode = http.StatusBadRequest
		return
	}
	//订单生成发送通知消息

	//calculate the total opspoint that need to be minus from buyer
	var totalAmountOpspoint float64
	sellerUidsAndAmount := make(map[string]int64, len(buyOCardOnOpsReq.PurchaseInfo))
	groupIds := make(map[string]int64)
	sellerUidsAndOpspoint := make(map[string]float64)
	for index, value := range buyOCardOnOpsReq.PurchaseInfo {
		totalAmountOpspoint += float64(value.Amount) * value.UnitPrice
		groupIds[value.GroupId] = value.Amount
		sellerUidsAndAmount[value.SellerUid] += value.Amount
		sellerUidsAndOpspoint[value.SellerUid] += float64(value.Amount) * value.UnitPrice
		buyOCardOnOpsReq.PurchaseInfo[index].BuyerUid = claims.Uid
	}

	//check opspoint satisfy to buy a new card
	code, err := GetOpsPoint(claims.Uid, totalAmountOpspoint)
	if err != nil {
		response.SetError(code, utils2.RecodeTest(code), err)
		statusCode = http.StatusBadRequest
		return
	}

	//minus the opspoint from buyer
	code, err = OperateOpspoint(claims.Uid, -totalAmountOpspoint)
	if code != utils2.RECODE_OK {
		response.SetError(code, utils2.RecodeTest(code), err)
		statusCode = http.StatusBadRequest
		return
	}

	//minus the amount ocard from each ocard by using groupid
	for groupId, amount := range groupIds {
		RpcSetOCardAmount := pbProperty.NewOperatePropertyService("property", consulreg.MicroSer.Client())
		cardAmount, err := RpcSetOCardAmount.OperateOCardAmount(context.TODO(), &pbProperty.OCard{GroupId: groupId, Amount: int64ToByte(amount)})
		if err != nil {
			response.SetError(utils2.RECODE_MICROERR, utils2.RecodeTest(utils2.RECODE_MICROERR), err)
			statusCode = http.StatusBadRequest
			return
		} else if cardAmount.Code == utils2.RECODE_STOREDATA_FAILED { //4006
			response.SetError(cardAmount.Code, utils2.RecodeTest(cardAmount.Code), err)
			statusCode = http.StatusBadRequest
			return
		} else if cardAmount.Code == utils2.RECODE_INSUFFICIENT_FUND {
			response.SetError(cardAmount.Code, "卡可出售数量不足", err)
			statusCode = http.StatusBadRequest
			return
		}
	}
	//minus the total amount in user_detail
	for sellerUid, amount := range sellerUidsAndAmount {
		RpcSetTotalAmountOCard := pbUserInfo.NewOperateUserInfoService("userInfo", consulreg.MicroSer.Client())
		minusTotalOCard, err := RpcSetTotalAmountOCard.SetTotalReleaseOCardNum(context.TODO(), &pbUserInfo.TotalReleaseOCardNum{Uid: sellerUid, ReleaseCardNum: -amount})
		if err != nil { //4005
			response.SetError(utils2.RECODE_MICROERR, utils2.RecodeTest(utils2.RECODE_MICROERR), err)
			statusCode = http.StatusBadRequest
			return
		} else if minusTotalOCard.Code == utils2.RECODE_DBERR { //4001
			response.SetError(minusTotalOCard.Code, utils2.RecodeTest(minusTotalOCard.Code), err)
			statusCode = http.StatusBadRequest
			return
		}
	}

	//start transfer ocard from seller to buyer
	RpcTransferOCard := pbProperty.NewOperatePropertyService("property", consulreg.MicroSer.Client())
	transferOCard, err := RpcTransferOCard.BuyOCardOnOps(context.TODO(), buyOCardOnOpsReq)
	if err != nil {
		response.SetError(utils2.RECODE_MICROERR, utils2.RecodeTest(utils2.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	} else if transferOCard.Code == utils2.RECODE_STOREDATA_FAILED { //4006
		response.SetError(transferOCard.Code, utils2.RecodeTest(transferOCard.Code), err)
		statusCode = http.StatusBadRequest
		return
	}
	//transfer success
	//add the opspoint to seller
	for sellerUid, opspoint := range sellerUidsAndOpspoint {
		code, err = OperateOpspoint(sellerUid, opspoint)
		if code != utils2.RECODE_OK {
			response.SetError(code, utils2.RecodeTest(code), err)
			statusCode = http.StatusBadRequest
			return
		}
	}
	//add buyer into follow list with seller
	RpcFollowList := pbFollower.NewOperateFollowService("follower", consulreg.MicroSer.Client())
	for sellerUid, _ := range sellerUidsAndAmount {
		follow, err := RpcFollowList.Follow(context.TODO(), &pbFollower.Follower{Uid: claims.Uid, Following: sellerUid})
		if err != nil {
			response.SetError(utils2.RECODE_MICROERR, utils2.RecodeTest(utils2.RECODE_MICROERR), err)
			statusCode = http.StatusBadRequest
			return
		} else if follow.Uid == utils2.RECODE_STOREDATA_FAILED {
			response.SetError(utils2.RECODE_STOREDATA_FAILED, utils2.RecodeTest(utils2.RECODE_STOREDATA_FAILED), err)
			statusCode = http.StatusBadRequest
			return
		}
		//在ops上买卡成功，发送通知消息
		err = utils2.SendFollowMsg(sellerUid, claims.Uid)
		if err != nil {
			logger.Error(err, "消息进入redis队列失败")
		}

		response.NewSuccess()
		statusCode = http.StatusOK
		return
	}

	//transfer failed
	//roll back the opspoint

}

func MintCard(ctx *gin.Context) {
	var (
		statusCode int
		req        property.MintCardReq
		response   property.MintCardResponse
	)
	ocardInfo := &pbProperty.OCard{}
	jsonIdReq := &pbProperty.JsonIdAndGroupId{}
	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()
	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils2.RECODE_DATAERR, "bind data failed in MintCard", err)
		statusCode = http.StatusBadRequest
		return
	}
	//Your code logic here
	//compare the gas fee using the fee from frontend and contract

	//store json_id to mongodb
	jsonIdReq.JsonId = utils2.NewJsonId()

	//generate uri
	contractClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(pbNft1155.CreateRequest)
	pbRequest.Data = []byte(nil)
	pbRequest.Uri = utils2.GetConfigStr("json_prefix") + jsonIdReq.JsonId + ".json"
	pbRequest.InitSupply = stringDecimalToBytes(req.Info[0].Amount)
	pbRequest.InitOwnerAddress = utils2.GetConfigStr("owner_address")

	respCreat1155, err := contractClient.Create(context.TODO(), pbRequest)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call Create failed", err)
		return
	}
	logger.Info("respCreat1155: ", respCreat1155)
	if respCreat1155.TransactionHash == "" {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call StoreMintOCardInfo failed", err)
		return
	}

	//store the relation between json id and group id into the mongodb
	jsonIdClient := pbProperty.NewOperatePropertyService("property", consulreg.MicroSer.Client())
	jsonId, err := jsonIdClient.RelationshipJsonIdAndGroupId(context.TODO(), jsonIdReq)
	if err != nil {
		response.SetError(utils2.RECODE_MICROERR, "bind data failed in RelationshipJsonIdAndGroupId", err)
		statusCode = http.StatusInternalServerError
		return
	} else if jsonId.Code != utils2.RECODE_OK {
		response.SetError(jsonId.Code, "bind data failed in MintCard", err)
		statusCode = http.StatusInternalServerError
		return
	}

	//store to mongodb
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	storeMintInfo := pbProperty.NewOperatePropertyService("property", consulreg.MicroSer.Client())
	ocardInfo.Uid = claims.Uid
	ocardInfo.TransactionHash = respCreat1155.TransactionHash
	ocardInfo.CardType = req.Info[0].Ocard
	ocardInfo.Amount = stringDecimalToBytes(req.Info[0].Amount)
	storeResult, err := storeMintInfo.StoreMintOCardInfo(context.TODO(), ocardInfo)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call StoreMintOCardInfo failed", err)
		return
	} else if storeResult.Code != utils2.RECODE_OK {
		statusCode = http.StatusInternalServerError
		response.SetError(storeResult.Code, "call StoreMintOCardInfo failed", err)
		return
	}
	statusCode = http.StatusOK
	response.NewSuccess()
}

// TransferCards godoc
// @Summary 批量转移用户 card
// @Description 把一批资产（ids）批量转移到指定的账户中。
// @ID TransferCards
// @tags 用户资产
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param TransferCards body property.TransferCardReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} property.TransferCardResponse 成功
// @Failure 4005 {object} property.TransferCardResponse "微服务可能挂了"
// @Failure 4004 {object} property.TransferCardResponse "输入参数错误"
// @Failure 4006 {object} property.TransferCardResponse "卡存入失败"
// @Failure 4301 {object} property.TransferCardResponse "用户余额不足"
// @Failure 4123 {object} property.TransferCardResponse "手续费不匹配"
// @Router /property/ocard/transfer-batch [POST]
func TransferCards(ctx *gin.Context) {
	var (
		statusCode int
		req        property.TransferCardReq
		response   property.TransferCardResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils2.RECODE_DATAERR, "bind data failed in TransferCards", err)
		statusCode = http.StatusBadRequest
		return
	}
	//Your code logic here
	//define var
	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	ids := make([][]byte, len(req.Ids))
	quantities := make([][]byte, len(req.Ids))
	NFTContractAddr := utils2.GetConfigStr("nft1155")

	//generate uris
	//uris := make([]string, len(req.Ids))

	var totalAmount float64
	for index, info := range req.Ids {
		//uris[index] = utils.GetConfigStr("json_prefix") + jsonIdAndGroupIds[index].JsonId + ".json"
		//amount[index] = stringDecimalToBytes(info.Amount)

		//check the ocard amount on ops enough to take away
		err := CheckOcardAmountOps(claims.Uid, info, req.Quantities[index])
		if err != nil {
			logger.Error(err)
			response.SetError(utils2.RECODE_INSUFFICIENT_FUND, "用户卡余额不足", err)
			statusCode = http.StatusBadRequest
			return
		}
		ids[index] = stringDecimalToBytes(info)
		quantities[index] = int64ToByte(req.Quantities[index])
		totalAmount += float64(req.Quantities[index])
	}

	//compare the gas fee by using frontend recived and backend recived
	pbGetGasfeeReq := &pbNft1155.TransferBatchIDReq{}
	pbGetGasfeeReq.NftContractAddr = NFTContractAddr
	pbGetGasfeeReq.ToAddress = req.ToAddress
	pbGetGasfeeReq.Ids = ids
	pbGetGasfeeReq.Quantities = quantities
	logger.Infof("getGasFeeReq=%+v  amount=%.18f", pbGetGasfeeReq, quantities)
	pbGetGasfeeResp, err := RpcNft1155Service.GetTransferBatchIDPrice(context.TODO(), pbGetGasfeeReq)
	if err != nil {
		response.SetError(utils2.RECODE_MICROERR, utils2.RecodeTest(utils2.RECODE_MICROERR), err)
		statusCode = http.StatusBadRequest
		return
	}

	ETH := GasToEth(pbGetGasfeeResp.GasLimit, new(big.Int).SetBytes(pbGetGasfeeResp.GasPrice))
	logger.Infof("eth is: %s, gasLimit=%d gasPrice=%s", ETH.Text('f', 18),
		pbGetGasfeeResp.GasLimit, new(big.Int).SetBytes(pbGetGasfeeResp.GasPrice).String())
	priceResponse, err := RpcSwapService.ETH2OPS(ctx, &pbSwap.MoneyRequest{Money: ETH.String()})
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call ETH2OPS failed", err)
		return
	}
	logger.Info("eth to ops price is: ", priceResponse.Money)
	priceResponseGasfee, err := stringDecimalToBigFloat(priceResponse.Money)
	if err != nil {
		response.SetError(utils2.RECODE_DATAERR, utils2.RecodeTest(utils2.RECODE_DATAERR), err)
		statusCode = http.StatusBadRequest
		return
	}

	rate := new(big.Float).SetFloat64(0.005)
	gasfeeFromClient, err := stringDecimalToBigFloat(req.Gasfee)
	if err != nil {
		response.SetError(utils2.RECODE_DATAERR, utils2.RecodeTest(utils2.RECODE_DATAERR), err)
		statusCode = http.StatusBadRequest
		return
	}

	if err = checkFloatGreatRate(gasfeeFromClient, priceResponseGasfee, rate); err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_WITHDRAW_GASFEE, "checkFloatGreatRate", err)
		return
	}

	//minus ocard and gasfee, the gasfee is opspoint
	gasfeeFromClientFloat64, _ := gasfeeFromClient.Float64()
	opspoint, _ := OperateOpspoint(claims.Uid, -gasfeeFromClientFloat64)
	if opspoint != utils2.RECODE_OK {
		statusCode = http.StatusInternalServerError
		response.SetError(opspoint, utils2.RecodeTest(opspoint), err)
		return
	}
	for index, value := range req.Ids {
		err := OpearteOcardAmountOps(claims.Uid, value, -req.Quantities[index])
		if err != nil {
			statusCode = http.StatusInternalServerError
			response.SetError(utils2.RECODE_MICROERR, utils2.RecodeTest(utils2.RECODE_MICROERR), err)
			return
		}
	}

	//call contract withdraw ocard
	pbRequest := new(pbNft1155.TransferBatchIDReq)
	pbRequest.NftContractAddr = NFTContractAddr
	pbRequest.ToAddress = req.ToAddress
	pbRequest.Ids = ids
	pbRequest.Quantities = quantities
	respCreat1155, err := RpcNft1155Service.TransferBatchID(context.TODO(), pbRequest)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call Create failed", err)
		return
	}
	logger.Info("respCreat1155: ", respCreat1155)
	if respCreat1155.TransactionHash == "" {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call TransferBatchID failed", err)
		// sent the notification to message service 转出卡失败
		err = utils2.SendTradeMsgFailed(respCreat1155.TransactionHash, req.ToAddress, "fromAddress", claims.Uid, totalAmount)
		if err != nil {
			logger.Error(err, "消息进入redis队列失败")
			return
		}
	} else {
		err = utils2.SendTradeMsgSuccess(respCreat1155.TransactionHash, req.ToAddress, "fromAddress", claims.Uid, totalAmount)
		if err != nil {
			logger.Error(err, "消息进入redis队列失败")
			return
		}
	}

	//cache transaction hash, uid, ocard, gasfee into mongo, maybe called withdraw_ocard_history
	for index, value := range req.Ids {
		remote, err := RpcPropertyService.StoreOCardFromChain(context.TODO(),
			&pbProperty.OCardsOnOps{BuyerUid: claims.Uid, GroupId: value, Amount: req.Quantities[index]})
		if err != nil {
			statusCode = http.StatusInternalServerError
			response.SetError(utils2.RECODE_MICROERR, utils2.RecodeTest(utils2.RECODE_MICROERR), err)
			return
		} else if remote.Code != utils2.RECODE_OK {
			statusCode = http.StatusInternalServerError
			response.SetError(utils2.RECODE_STOREDATA_FAILED, utils2.RecodeTest(utils2.RECODE_STOREDATA_FAILED), err)
			return
		}
	}
	//get the notification from contract to judge that the transaction is success or failed

	//success: send the message to the message service

	//failed: roll back the opspoint and return the transaction is failed
	statusCode = http.StatusOK
	response.NewSuccess()
	response.Data.TransactionHash = respCreat1155.GetTransactionHash()
}

// QueryTransaction godoc
// @Summary 查询 ocard 相关的交易信息
// @Description 返回 ocard 交易信息详情，注意：在交易还是 pending 的状态是没有数据的。
// @ID QueryTransaction
// @tags 用户资产
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param QueryTransaction body property.QueryTransactionReq true "传入交易后的 transaction_hash"
// @Produce  json
// @Success 0 {object} property.QueryTransactionResponse 成功
// @Failure 4005 {object} property.QueryTransactionResponse "微服务可能挂了"
// @Failure 4004 {object} property.QueryTransactionResponse "输入参数错误"
// @Router /property/ocard/query-transaction [POST]
func QueryTransaction(ctx *gin.Context) {
	var (
		statusCode int
		req        property.QueryTransactionReq
		response   property.QueryTransactionResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils2.RECODE_DATAERR, "bind data failed in QueryTransaction", err)
		statusCode = http.StatusBadRequest
		return
	}
	//Your code logic here
	pbReq := new(pbNft1155.GetTransactionByHashReq)
	pbReq.TransactionHash = req.TransactionHash
	pbResp, err := RpcNft1155Service.GetTransactionByHash(context.TODO(), pbReq)
	if err != nil {
		response.SetError(utils2.RECODE_MICROERR, "bind data failed in QueryTransaction", err)
		statusCode = http.StatusBadRequest
		return
	}

	response.Data.Data = pbResp.Data
	response.Data.ContractAddress = pbResp.ContractAddress
	response.Data.ContractFrom = pbResp.ContractFrom
	response.Data.ContractTo = pbResp.ContractTo
	response.Data.IsPending = pbResp.IsPending
	var allAmount uint64
	for i := range pbResp.IdAndAmount {
		a := intBytesToUint64(pbResp.IdAndAmount[i].Amount)
		response.Data.IdAndAmount = append(response.Data.IdAndAmount, property.IDAmount{
			ID:     intBytesToUint64(pbResp.IdAndAmount[i].Id),
			Amount: a,
		})
		allAmount += a
	}

	claims := ctx.MustGet("claims").(*myjwt.CustomClaims)
	//send notify message
	//if !pbResp.IsPending{
	err = utils2.SendTradeMsgPending(req.TransactionHash, pbResp.ContractTo,
		pbResp.ContractFrom, claims.Uid, float64(allAmount))
	if err != nil {
		logger.Error(err)
	}
	//}

	logger.Infof("%+v", pbResp)

	statusCode = http.StatusOK
	response.NewSuccess()
}

/*
// ChargeToServer godoc
// @Summary 用户充值一批卡到服务器。
// @Description 从用户钱包到服务器
// @ID ChargeToServer
// @tags 用户资产
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param ChargeToServer body property.ChargeToServerReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} property.ChargeToServerResponse 成功
// @Failure 4005 {object} property.ChargeToServerResponse "微服务可能挂了"
// @Failure 4004 {object} property.ChargeToServerResponse "输入参数错误"
// @Router /property/ocard/charge-to-server [POST]
func ChargeToServer(ctx *gin.Context){
	var(
		statusCode int
		req property.ChargeToServerReq
		response property.ChargeToServerResponse
		nft1155ManagerAddress = utils2.GetConfigStr("owner_address")
		claims            = ctx.MustGet("claims").(*myjwt.CustomClaims)
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils2.RECODE_DATAERR, "bind data failed in ChargeToServer", err )
		statusCode = http.StatusBadRequest
		return
	}
	reqGasfee, err := stringDecimalToBigFloat(req.GasFee)
	if err != nil {
		response.SetError(utils2.RECODE_DATAERR, "stringDecimalToBigFloat", err)
		statusCode = http.StatusBadRequest
		return
	}

	pbReq := &pbNft1155.TransferBatchIDReq{
		NftContractAddr: req.FromAddress,
		ToAddress:       nft1155ManagerAddress,
		Ids:             utils2.StringNumbersToBigIntBytes(req.Ids),
		Quantities:      utils2.StringNumbersToBigIntBytes(req.Quantities),
	}
	pbGetGasfeeResp, err := property.RpcNft1155Service.GetTransferBatchIDPrice(context.TODO(), pbReq)
	if err != nil{
		response.SetError(utils2.RECODE_MICROERR, "bind data failed in QueryTransactionCardFee", err )
		statusCode = http.StatusInternalServerError
		return
	}

	ETH := GasToEth(pbGetGasfeeResp.GasLimit, new(big.Int).SetBytes(pbGetGasfeeResp.GasPrice))
	logger.Infof("eth is: %s, gasLimit=%d gasPrice=%s", ETH.Text('f', 18),
		pbGetGasfeeResp.GasLimit, new(big.Int).SetBytes(pbGetGasfeeResp.GasPrice).String())
	priceResponse, err := property.RpcSwapService.ETH2OPS(ctx, &pbSwap.MoneyRequest{Money: ETH.String()})
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call ETH2OPS failed", err)
		return
	}
	priceNumber, err := stringDecimalToBigFloat(priceResponse.Money)
	if err != nil{
		response.SetError(utils.RECODE_DATAERR, utils.RecodeTest(utils.RECODE_DATAERR), err)
		statusCode = http.StatusBadRequest
		return
	}

	rate := new(big.Float).SetFloat64(0.005)
	if err = checkFloatGreatRate(reqGasfee, priceNumber, rate); err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_WITHDRAW_GASFEE, "checkFloatGreatRate", err)
		return
	}
	numberF64, _ := priceNumber.Float64()
	decOpsReq := &pbProperty.IncDecOpsPointReq{
		Op:       pbProperty.OpsIncDec_Dec,
		OpsPoint: numberF64,
		Uid:      claims.Uid,
	}
	decOpsResult, err := property.RpcPropertyService.IncDecOpsPoint(ctx, decOpsReq)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "checkFloatGreatRate", err)
		return
	}
	logger.Infof("user %s balance is: %f", claims.Uid, decOpsResult.OpsBalance)


	//Your code logic here
	//1. 查询用户的 ops 数量是否能负担本笔交易的费用, （需暴露转卡手续费接口）
	//2. 查询用户 ocard 数量是否能转移出去
	//3. 扣掉 OPS
	//4. 扣掉 ocard

	//3.1 交易失败回滚 ops
	//4.1 交易失败回滚 ocard

	statusCode = http.StatusOK
	response.NewSuccess()
}
*/

// QueryTransactionCardFee godoc
// @Summary 查询转卡所需要的手续费
// @Description 需要输入 from，to 地址，ids, amounts 列表，
// @ID QueryTransactionCardFee
// @tags 用户资产
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param QueryTransactionFee body property.QueryTransactionFeeReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} property.QueryTransactionFeeResponse 成功
// @Failure 4005 {object} property.QueryTransactionFeeResponse "微服务可能挂了"
// @Failure 4004 {object} property.QueryTransactionFeeResponse "输入参数错误"
// @Router /property/ocard/query-transaction-fee [POST]
func QueryTransactionCardFee(ctx *gin.Context) {
	var (
		statusCode int
		req        property.QueryTransactionFeeReq
		response   property.QueryTransactionFeeResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils2.RECODE_DATAERR, "bind data failed in QueryTransactionCardFee", err)
		statusCode = http.StatusBadRequest
		return
	}

	//Your code logic here
	pbReq := &pbNft1155.TransferBatchIDReq{
		NftContractAddr: req.NFTContractAddr,
		ToAddress:       req.ToAddress,
		Ids:             utils2.StringNumbersToBigIntBytes(req.Ids),
		Quantities:      utils2.StringNumbersToBigIntBytes(req.Quantities),
	}
	pbGetGasfeeResp, err := RpcNft1155Service.GetTransferBatchIDPrice(context.TODO(), pbReq)
	if err != nil {
		response.SetError(utils2.RECODE_MICROERR, "bind data failed in QueryTransactionCardFee", err)
		statusCode = http.StatusInternalServerError
		return
	}

	ETH := GasToEth(pbGetGasfeeResp.GasLimit, new(big.Int).SetBytes(pbGetGasfeeResp.GasPrice))
	logger.Infof("eth is: %s, gasLimit=%d gasPrice=%s", ETH.Text('f', 18),
		pbGetGasfeeResp.GasLimit, new(big.Int).SetBytes(pbGetGasfeeResp.GasPrice).String())
	priceResponse, err := RpcSwapService.ETH2OPS(ctx, &pbSwap.MoneyRequest{Money: ETH.String()})
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils2.RECODE_MICROERR, "call ETH2OPS failed", err)
		return
	}
	response.Data.Money = priceResponse.Money

	statusCode = http.StatusOK
	response.NewSuccess()
}
