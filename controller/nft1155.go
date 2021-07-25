package controller

import (
	"context"
	"github.com/asim/go-micro/v3/logger"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"ops/pkg/model/consulreg"
	webModel "ops/pkg/model/contract"
	"ops/pkg/utils"
	pbNft1155 "ops/proto/nft1155"
	pbSwap "ops/proto/swap"
)

// NFT1155Info godoc
// @Summary 查询 nft1155 的 name, symbol, token_supply, owner
// @Description NFT1155Info 查询合约的基本信息
// @ID ContractNFT1155Info
// @tags NFT1155_Reading
// @Accept json
// @Produce  json
// @Success 0 {object} webModel.NFT1155Info "返回 NFT1155 合约的基本信息"
// @Failure 4005 {object} webModel.NFT1155Info "微服务可能挂了"
// @Router /contract/nft1155/info [POST]
func NFT1155Info(ctx *gin.Context) {
	var (
		response   webModel.NFT1155Info
		statusCode int
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	microClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	nameResponse, err := microClient.Name(context.TODO(), new(pbNft1155.NameRequest))
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "call Name error", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data.Name = nameResponse.GetName()

	ownerResponse, err := microClient.Owner(context.TODO(), new(pbNft1155.OwnerRequest))
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "call Owner error", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data.Owner = ownerResponse.OwnerAddress

	tokenSupplyResponse, err := microClient.TokenSupply(context.TODO(), new(pbNft1155.TokenSupplyRequest))
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "call TokenSupply error", err)
		statusCode = http.StatusInternalServerError
		return
	}
	bi := new(big.Int)
	bi.SetBytes(tokenSupplyResponse.Amount)
	response.Data.TokenSupply = bi.String()

	symbolResponse, err := microClient.Symbol(context.TODO(), new(pbNft1155.SymbolRequest))
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "call Symbol error", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data.Symbol = symbolResponse.Symbol
	statusCode = http.StatusOK
	response.NewSuccess()
}

// NFT1155NextTokenID godoc
// @Summary 查询 nft1155 的 next token
// @Description NFT1155Info 获取下一个资产类别 ID
// @ID NFT1155NextTokenID
// @tags NFT1155_Reading
// @Accept json
// @Produce  json
// @Success 0 {object} webModel.NFT1155NextTokenID "返回 NFT1155 下一个资产类别 ID，类型为十进制 int 字符串"
// @Failure 4005 {object} webModel.NFT1155NextTokenID "微服务可能挂了"
// @Router /contract/nft1155/get-next-token-id [POST]
func NFT1155NextTokenID(ctx *gin.Context) {
	var (
		response   webModel.NFT1155NextTokenID
		statusCode int
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	microClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	resp, err := microClient.GetNextTokenID(context.TODO(), new(pbNft1155.GetNextTokenIDRequest))
	if err != nil {
		response.SetError(utils.RECODE_MICROERR, "call Name error", err)
		statusCode = http.StatusInternalServerError
		return
	}
	response.Data.ID = intBytesToString(resp.Id)
	statusCode = http.StatusOK
	response.NewSuccess()
}

// NFT1155Balance godoc
// @Summary 查询 nft1155 的某个钱包的资产 id 余额
// @Description NFT1155Balance 查询 nft1155 的钱包地址的余额
// @ID NFT1155Balance
// @tags NFT1155_Reading
// @Accept json
// @Param idAddress body webModel.NFT1155BalanceReq.id true "NFT 资产 id 和地址"
// @Produce  json
// @Success 0 {object} webModel.NFT1155Balance.code "返回查询地址的余额, Balance 为十进制 int 字符串"
// @Failure 4005 {object} webModel.NFT1155Balance.code "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155Balance.code "输入参数错误"
// @Router /contract/nft1155/balance-of [POST]
func NFT1155Balance(ctx *gin.Context) {
	var (
		statusCode int
		response   webModel.NFT1155Balance
		request    webModel.NFT1155BalanceReq
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&request); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in BalanceOf", err)
		statusCode = http.StatusBadRequest
		return
	}
	logger.Info(request)

	microClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(pbNft1155.BalanceOfRequest)
	pbRequest.Id = stringDecimalToBytes(request.ID)
	pbRequest.OwnerAddress = request.OwnerAddress
	resp, err := microClient.BalanceOf(context.TODO(), pbRequest)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call balance of failed", err)
		return
	}
	response.Data.Amount = intBytesToString(resp.Balance)
	response.Data.OwnerAddress = request.OwnerAddress
	response.Data.ID = request.ID
	statusCode = http.StatusOK
	response.NewSuccess()
}

// NFT1155BalanceBatch godoc
// @Summary 批量查询 nft1155 的多个钱包地址的余额
// @Description NFT1155BalanceBatch 批量查询
// @ID NFT1155BalanceBatch
// @tags NFT1155_Reading
// @Accept json
// @Param idAddressArray body []webModel.NFT1155BalanceReq true "NFT 资产 id 和地址"
// @Produce  json
// @Success 0 {object} webModel.NFT1155BalanceBatchResponse "返回多个地址的余额, Balance 为十进制 int 字符串"
// @Failure 4005 {object} webModel.NFT1155BalanceBatchResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155BalanceBatchResponse "输入参数错误"
// @Router /contract/nft1155/balance-of-batch [POST]
func NFT1155BalanceBatch(ctx *gin.Context) {
	var (
		statusCode int
		response   webModel.NFT1155BalanceBatchResponse
		request    []webModel.NFT1155BalanceReq
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&request); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in BalanceOfBatch", err)
		statusCode = http.StatusBadRequest
		return
	}

	microClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(pbNft1155.BalanceOfBatchRequest)
	for i := 0; i < len(request); i++ {
		pbRequest.Owners = append(pbRequest.Owners, &pbNft1155.BalanceOfRequest{
			OwnerAddress: request[i].OwnerAddress,
			Id:           stringDecimalToBytes(request[i].ID),
		})
	}
	resp, err := microClient.BalanceOfBatch(context.TODO(), pbRequest)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call balance of failed", err)
		return
	}
	for i := 0; i < len(request); i++ {
		response.Data = append(response.Data, webModel.NFTBalanceType{
			NFT1155IDType: webModel.NFT1155IDType{ID: request[i].ID},
			OwnerAddress:  string(resp.Balances[i].OwnerAddress),
			Amount:        intBytesToString(resp.Balances[i].Balance),
		})
	}
	statusCode = http.StatusOK
	response.NewSuccess()
}

// NFT1155IsApprovedForAll godoc
// @Summary 查询用户 _owner 是否给 _operator地址操作 nft 资产的权限
// @Description NFT1155IsApprovedForAll 权限查询
// @ID NFT1155IsApprovedForAll
// @tags NFT1155_Reading
// @Accept json
// @Param operatorOwner body webModel.NFT1155ApprovedReq.operator true "传入 operator 和 owner 地址"
// @Produce  json
// @Success 0 {object} webModel.NFT1155ApprovedResponse "返回查询地址的余额, Balance 为十进制 int 字符串"
// @Failure 4005 {object} webModel.NFT1155ApprovedResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155ApprovedResponse "输入参数错误"
// @Router /contract/nft1155/is-approved-for-all [POST]
func NFT1155IsApprovedForAll(ctx *gin.Context) {
	var (
		statusCode int
		req        webModel.NFT1155ApprovedReq
		response   webModel.NFT1155ApprovedResponse
	)
	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in NFT1155IsApprovedForAll", err)
		statusCode = http.StatusBadRequest
		return
	}

	microClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(pbNft1155.IsApprovedForAllRequest)
	pbRequest.OwnerAddress = req.OwnerAddress
	pbRequest.OperatorAddress = req.OperatorAddress
	resp, err := microClient.IsApprovedForAll(context.TODO(), pbRequest)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call balance of failed", err)
		return
	}
	response.Data.IsApproved = resp.IsOperator
	statusCode = http.StatusOK
	response.NewSuccess()
}

// NFT1155URI godoc
// @Summary 返回资产类别 id 对应 NFT Uri 地址
// @Description NFT1155URI 返回资产类别 id 对应 NFT Uri 地址
// @ID NFT1155URI
// @tags NFT1155_Reading
// @Accept json
// @Param id body webModel.NFT1155URIRequest true "传入 NFT id"
// @Produce  json
// @Success 0 {object} webModel.NFT1155URIResponse "返回查询的 Uri 地址"
// @Failure 4005 {object} webModel.NFT1155URIResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155URIResponse "输入参数错误"
// @Router /contract/nft1155/uri [POST]
func NFT1155URI(ctx *gin.Context) {
	var (
		statusCode int
		req        webModel.NFT1155URIRequest
		response   webModel.NFT1155URIResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in NFT1155URI", err)
		statusCode = http.StatusBadRequest
		return
	}

	microClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(pbNft1155.UriRequest)
	pbRequest.Id = stringDecimalToBytes(req.ID)
	resp, err := microClient.Uri(context.TODO(), pbRequest)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call Uri failed", err)
		return
	}
	response.Data.Uri = resp.Uri
	statusCode = http.StatusOK
	response.NewSuccess()
}

// NFT1155Create godoc
// @Summary 创建一类新资产
// @Description 创建一类新资产，_initialOwner：资产拥有者，_initialSupply 资产数量，_data额外数量,可为"0x00"下同,_uri 资产对应的url（为"" 即默认规则）
// @ID NFT1155Create
// @tags NFT1155_Writing
// @Accept json
// @Param newProperty body webModel.NFT1155CreateReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} webModel.NFT1155CreateResponse 成功
// @Failure 4005 {object} webModel.NFT1155CreateResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155CreateResponse "输入参数错误"
// @Router /contract/nft1155/create [POST]
func NFT1155Create(ctx *gin.Context) {
	var (
		statusCode int
		req        webModel.NFT1155CreateReq
		response   webModel.NFT1155CreateResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in NFT1155Create", err)
		statusCode = http.StatusBadRequest
		return
	}

	microClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(pbNft1155.CreateRequest)
	pbRequest.Data = []byte(req.Data)
	pbRequest.Uri = req.Uri
	pbRequest.InitSupply = stringDecimalToBytes(req.InitSupply)
	pbRequest.InitOwnerAddress = req.InitOwner

	resp, err := microClient.Create(context.TODO(), pbRequest)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call Create failed", err)
		return
	}
	logger.Info(resp)
	//response.Data.Uri = resp.Uri
	statusCode = http.StatusOK
	response.NewSuccess()
}

// NFT1155CreateBatch godoc
// @Summary 批量创建新资产，_initialOwner：资产拥有者，_quantities 资产数量数组 ,_uris 资产对应的url（为"" 即默认规则）
// @Description 批量创建新资产，_initialOwner：资产拥有者，_quantities 资产数量数组 ,_uris 资产对应的url（为"" 即默认规则）
// @ID NFT1155CreateBatch
// @tags NFT1155_Writing
// @Accept json
// @Param newProperty body webModel.NFT1155CreateBatchReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} webModel.NFT1155CreateBatchResponse 成功
// @Failure 4005 {object} webModel.NFT1155CreateBatchResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155CreateBatchResponse "输入参数错误"
// @Router /contract/nft1155/create [POST]
func NFT1155CreateBatch(ctx *gin.Context) {
	var (
		statusCode int
		req        webModel.NFT1155CreateBatchReq
		response   webModel.NFT1155CreateBatchResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in NFT1155CreateBatch", err)
		statusCode = http.StatusBadRequest
		return
	}

	microClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(pbNft1155.CreateBatchRequest)
	pbRequest.Data = []byte(req.Data)
	pbRequest.Uris = req.Uris
	for _, q := range req.Quantities {
		pbRequest.Quantities = append(pbRequest.Quantities, stringDecimalToBytes(q))
	}
	pbRequest.InitOwnerAddress = req.InitOwner

	resp, err := microClient.CreateBatch(ctx, pbRequest)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call NFT1155CreateBatch failed", err)
		return
	}
	logger.Info(resp)
	statusCode = http.StatusOK
	response.NewSuccess()
}

// NFT1155CreateBatchPrice godoc
// @Summary 计算批量创建资产所需要的手续费（单位：OPS）
// @Description 返回批量创建资产所需要的手续费。（单位：OPS）
// @ID NFT1155CreateBatchPrice
// @tags NFT1155_Reading
// @Accept json
// @Param token header string true "header中的token自带了uid"
// @Param NFT1155CreateBatchPrice body webModel.NFT1155CreateBatchPriceReq true "和create-batch创建资产所需要的参数一样"
// @Produce  json
// @Success 0 {object} webModel.NFT1155CreateBatchPriceResponse 成功
// @Failure 4005 {object} webModel.NFT1155CreateBatchPriceResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155CreateBatchPriceResponse "输入参数错误"
// @Router /contract/nft1155/create-batch-price [POST]
func NFT1155CreateBatchPrice(ctx *gin.Context) {
	var (
		statusCode int
		req        webModel.NFT1155CreateBatchPriceReq
		response   webModel.NFT1155CreateBatchPriceResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in NFT1155CreateBatchPrice", err)
		statusCode = http.StatusBadRequest
		return
	}
	//Your code logic here
	microClient := pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(pbNft1155.CreateBatchRequest)
	pbRequest.Data = []byte(req.Data)
	for index, uri := range req.Uris {
		req.Uris[index] = utils.GetConfigStr("json_prefix") + uri + ".json"
	}
	pbRequest.Uris = req.Uris
	for _, q := range req.Quantities {
		pbRequest.Quantities = append(pbRequest.Quantities, stringDecimalToBytes(q))
	}
	pbRequest.InitOwnerAddress = req.InitOwner
	resp, err := microClient.GetCreateBatchPrice(ctx, pbRequest)
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call GetCreateBatchPrice failed", err)
		return
	}

	ETH := GasToEth(resp.GasLimit, new(big.Int).SetBytes(resp.GasPrice))
	logger.Infof("eth is: %s, gasLimit=%d gasPrice=%s", ETH.Text('f', 18),
		resp.GasLimit, new(big.Int).SetBytes(resp.GasPrice).String())

	swapClient := pbSwap.NewSwapService("contract", consulreg.MicroSer.Client())
	priceResponse, err := swapClient.ETH2OPS(ctx, &pbSwap.MoneyRequest{Money: ETH.String()})
	if err != nil {
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call GetCreateBatchPrice failed", err)
		return
	}
	response.Data.Money = priceResponse.Money

	statusCode = http.StatusOK
	response.NewSuccess()
}

/*
// NFT1155SetBaseMetaURI godoc
// @Summary 设置统一 URI 前缀
// @Description 设置统一 URI 前缀
// @ID SetBaseMetaURI
// @tags NFT1155_Writing
// @Accept json
// @Param id body webModel.NFT1155BaseMetaURIReq true "传入前缀"
// @Produce  json
// @Success 0 {object} webModel.NFT1155BaseMetaURIResponse 成功
// @Failure 4005 {object} webModel.NFT1155BaseMetaURIResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155BaseMetaURIResponse "输入参数错误"
// @Router /contract/nft1155/set-base-meta-uri [PUT]
func NFT1155SetBaseMetaURI(ctx *gin.Context){
	var(
		statusCode int
		req webModel.NFT1155BaseMetaURIReq
		response webModel.NFT1155BaseMetaURIResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in NFT1155SetBaseMetaURI", err )
		statusCode = http.StatusBadRequest
		return
	}

	microClient := nft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(nft1155.SetBaseMetadataURIRequest)
	pbRequest.BaseMetadataURI = req.Prefix
	resp, err := microClient.SetBaseMetadataURI(context.TODO(), pbRequest)
	if err != nil{
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call SetBaseMetadataURI failed", err);
		return
	}
	logger.Info(resp)
	//response.Data.Uri = resp.Uri
	statusCode = http.StatusOK
	response.NewSuccess()
}


// NFT1155Create godoc
// @Summary 创建一类新资产
// @Description 创建一类新资产，_initialOwner：资产拥有者，_initialSupply 资产数量，_data额外数量,可为"0x00"下同,_uri 资产对应的url（为"" 即默认规则）
// @ID NFT1155Create
// @tags NFT1155_Writing
// @Accept json
// @Param newProperty body webModel.NFT1155CreateReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} webModel.NFT1155CreateResponse 成功
// @Failure 4005 {object} webModel.NFT1155CreateResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155CreateResponse "输入参数错误"
// @Router /contract/nft1155/create [POST]
func NFT1155Create(ctx *gin.Context){
	var(
		statusCode int
		req webModel.NFT1155CreateReq
		response webModel.NFT1155CreateResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in NFT1155Create", err )
		statusCode = http.StatusBadRequest
		return
	}

	microClient := nft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(nft1155.CreateRequest)
	pbRequest.Data = []byte(req.Data)
	pbRequest.Uri = req.Uri
	pbRequest.InitSupply = stringDecimalToBytes(req.InitSupply)
	pbRequest.InitOwnerAddress = []byte(req.InitOwner)

	resp, err := microClient.Create(context.TODO(), pbRequest)
	if err != nil{
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call Create failed", err);
		return
	}
	logger.Info(resp)
	//response.Data.Uri = resp.Uri
	statusCode = http.StatusOK
	response.NewSuccess()
}




// NFT1155TransferGovernorship godoc
// @Summary 转移管理者权限到newGovernor地址
// @Description 转移管理者权限到newGovernor地址
// @ID NFT1155TransferGovernorship
// @tags NFT1155_Writing
// @Accept json
// @Param newProperty body webModel.NFT1155TransferGovernorShipReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} webModel.NFT1155TransferGovernorShipResponse 成功
// @Failure 4005 {object} webModel.NFT1155TransferGovernorShipResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155TransferGovernorShipResponse "输入参数错误"
// @Router /contract/nft1155/transferGovernorship [PUT]
func NFT1155TransferGovernorship(ctx *gin.Context){
	var(
		statusCode int
		req webModel.NFT1155TransferGovernorShipReq
		response webModel.NFT1155TransferGovernorShipResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in NFT1155TransferGovernorship", err )
		statusCode = http.StatusBadRequest
		return
	}

	microClient := nft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(nft1155.TransferGovernorshipRequest)
	pbRequest.NewGovernorAddress = []byte(req.NewGovernorAddress)

	resp, err := microClient.TransferGovernorship(ctx, pbRequest)
	if err != nil{
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call NFT1155TransferGovernorship failed", err);
		return
	}
	logger.Info(resp)
	statusCode = http.StatusOK
	response.NewSuccess()
}


// NFT1155Mint godoc
// @Summary 资产类别为_id(资产类别需要事先创建) 增发_quantity数量给_to
// @Description 资产类别为_id(资产类别需要事先创建) 增发_quantity数量给_to
// @ID NFT1155Mint
// @tags NFT1155_Writing
// @Accept json
// @Param newProperty body webModel.NFT1155MintReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} webModel.NFT1155MintResponse 成功
// @Failure 4005 {object} webModel.NFT1155MintResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155MintResponse "输入参数错误"
// @Router /contract/nft1155/mint [POST]
func NFT1155Mint(ctx *gin.Context){
	var(
		statusCode int
		req webModel.NFT1155MintReq
		response webModel.NFT1155MintResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in NFT1155Mint", err )
		statusCode = http.StatusBadRequest
		return
	}

	microClient := nft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(nft1155.MintRequest)
	pbRequest.Id = stringDecimalToBytes(req.ID)
	pbRequest.AddressTo = []byte(req.AddressTo)
	pbRequest.Quantity = stringDecimalToBytes(req.Quantity)
	pbRequest.Data = []byte(req.Data)

	resp, err := microClient.Mint(ctx, pbRequest)
	if err != nil{
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call NFT1155Mint failed", err);
		return
	}
	logger.Info(resp)
	statusCode = http.StatusOK
	response.NewSuccess()
}

// NFT1155MintBatch godoc
// @Summary 资产类别为_id(资产类别需要事先创建) 增发_quantity数量给_to
// @Description 资产类别为_id(资产类别需要事先创建) 增发_quantity数量给_to
// @ID NFT1155MintBatch
// @tags NFT1155_Writing
// @Accept json
// @Param newProperty body webModel.NFT1155MintBatchReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} webModel.NFT1155MintBatchResponse 成功
// @Failure 4005 {object} webModel.NFT1155MintBatchResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155MintBatchResponse "输入参数错误"
// @Router /contract/nft1155/mint-batch [POST]
func NFT1155MintBatch(ctx *gin.Context){
	var(
		statusCode int
		req webModel.NFT1155MintBatchReq
		response webModel.NFT1155MintBatchResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in NFT1155MintBatch", err )
		statusCode = http.StatusBadRequest
		return
	}

	microClient := nft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(nft1155.MintBatchRequest)
	for i := range req.Ids {
		pbRequest.Ids = append(pbRequest.Ids, stringDecimalToBytes(req.Ids[i]))
		pbRequest.Quantities = append(pbRequest.Ids, stringDecimalToBytes(req.Quantities[i]))
	}
	pbRequest.Data = []byte(req.Data)

	resp, err := microClient.MintBatch(ctx, pbRequest)
	if err != nil{
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call NFT1155MintBatch failed", err);
		return
	}
	logger.Info(resp)
	statusCode = http.StatusOK
	response.NewSuccess()

}

// NFT1155SetCreator godoc
// @Summary 将增发资产类别为_id数组的权限转移给_to
// @Description 将增发资产类别为_id数组的权限转移给_to
// @ID NFT1155SetCreator
// @tags NFT1155_Writing
// @Accept json
// @Param creator body webModel.NFT1155SetCreatorReq true "传入相关参数"
// @Produce  json
// @Success 0 {object} webModel.NFT1155SetCreatorResponse 成功
// @Failure 4005 {object} webModel.NFT1155SetCreatorResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155SetCreatorResponse "输入参数错误"
// @Router /contract/nft1155/set-creator [PUT]
func NFT1155SetCreator(ctx *gin.Context){
	var(
		statusCode int
		req webModel.NFT1155SetCreatorReq
		response webModel.NFT1155SetCreatorResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in NFT1155SetCreator", err )
		statusCode = http.StatusBadRequest
		return
	}

	microClient := nft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(nft1155.SetCreatorRequest)
	for _, id := range req.Ids {
		pbRequest.Ids = append(pbRequest.Ids, stringDecimalToBytes(id))
	}
	pbRequest.AddressTo = []byte(req.AddressTo)

	resp, err := microClient.SetCreator(ctx, pbRequest)
	if err != nil{
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call NFT1155SetCreator failed", err);
		return
	}
	logger.Info(resp)
	statusCode = http.StatusOK
	response.NewSuccess()
}

// NFT1155SetIdURI godoc
// @Summary 设置_id资产类型的url
// @Description 设置_id资产类型的url
// @ID NFT1155SetIdURI
// @tags NFT1155_Writing
// @Accept json
// @Param creator body webModel.NFT1155SetIDUri true "传入相关参数"
// @Produce  json
// @Success 0 {object} webModel.NFT1155SetIDUriResponse 成功
// @Failure 4005 {object} webModel.NFT1155SetIDUriResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155SetIDUriResponse "输入参数错误"
// @Router /contract/nft1155/set-id-uri [PUT]
func NFT1155SetIdURI(ctx *gin.Context){
	var(
		statusCode int
		req webModel.NFT1155SetIDUri
		response webModel.NFT1155SetIDUriResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in NFT1155SetIdURI", err )
		statusCode = http.StatusBadRequest
		return
	}

	microClient := nft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(nft1155.SetIdURIRequest)
	pbRequest.Id = stringDecimalToBytes(req.ID)
	pbRequest.Uri = req.Uri

	resp, err := microClient.SetIdURI(ctx, pbRequest)
	if err != nil{
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call NFT1155SetIdURI failed", err);
		return
	}
	logger.Info(resp)
	statusCode = http.StatusOK
	response.NewSuccess()
}

// NFT1155SetIdURIBatch godoc
// @Summary 批量设置_id资产类型的url
// @Description 批量设置_id资产类型的url
// @ID NFT1155SetIdURIBatch
// @tags NFT1155_Writing
// @Accept json
// @Param creator body []webModel.NFT1155SetIDUri true "传入相关参数"
// @Produce  json
// @Success 0 {object} webModel.NFT1155SetIDUriResponse 成功
// @Failure 4005 {object} webModel.NFT1155SetIDUriBatchResponse "微服务可能挂了"
// @Failure 4004 {object} webModel.NFT1155SetIDUriBatchResponse "输入参数错误"
// @Router /contract/nft1155/set-id-uri-batch [PUT]
func NFT1155SetIdURIBatch(ctx *gin.Context){
	var(
		statusCode int
		req []webModel.NFT1155SetIDUri
		response webModel.NFT1155SetIDUriResponse
	)

	defer func() {
		responseHTTP(ctx, statusCode, &response)
	}()

	if err := ctx.ShouldBind(&req); err != nil {
		response.SetError(utils.RECODE_DATAERR, "bind data failed in NFT1155SetIdURIBatch", err )
		statusCode = http.StatusBadRequest
		return
	}

	microClient := nft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	pbRequest := new(nft1155.SetIdURIsRequest)
	for _, r := range req {
		pbRequest.Ids = append(pbRequest.Ids, stringDecimalToBytes(r.ID))
		pbRequest.Uris = append(pbRequest.Uris, r.Uri)
	}

	resp, err := microClient.SetIdURIs(ctx, pbRequest)
	if err != nil{
		statusCode = http.StatusInternalServerError
		response.SetError(utils.RECODE_MICROERR, "call NFT1155SetIdURIBatch failed", err);
		return
	}
	logger.Info(resp)
	statusCode = http.StatusOK
	response.NewSuccess()
}

*/
