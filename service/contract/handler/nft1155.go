package handler

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"net/url"
	"ops/pkg/logger"
	"ops/pkg/utils"
	"ops/proto/nft1155"
	nft1155Model "ops/service/contract/ercpkg/nft1155"
	"ops/service/contract/ercpkg/nftholder"
)

type nft1155Handler struct {
	client *ethclient.Client
	rpcClient *rpc.Client
	contractInstance *nft1155Model.Nft1155
	contractAddress common.Address
	privateKey *ecdsa.PrivateKey

	holderInstance *nftholder.Nftholder
	holderContractAddr common.Address
}

func NewNft1155Handler(opts *Options)(*nft1155Handler, error){
	parsedURL, err := url.Parse(opts.BlockChainURLHttp)
	if err != nil{
		return nil, err
	}
	if parsedURL.Scheme == "" ||
		parsedURL.Scheme == "ws" ||
		parsedURL.Scheme == "wss" { //subscribe chain event
		logger.Info("use web socket connection, subscribing the chain event now.")
	}
	ch := new(nft1155Handler)

	ch.privateKey, err = crypto.HexToECDSA(opts.NFT1155PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("private Key is not corrected - %w", err)
	}
	ch.rpcClient, err = rpc.Dial(opts.BlockChainURLHttp)
	if err != nil{
		return nil, err
	}
	//TODO: ethClient has some bugs, it need to be replaced.
	ch.client = ethclient.NewClient(ch.rpcClient)

	ch.contractAddress = opts.NFT1155ContractAddr
	ch.contractInstance, err = nft1155Model.NewNft1155(ch.contractAddress, ch.client)
	if err != nil {
		return nil, err
	}

	ch.holderContractAddr = opts.HolderContractAddr
	ch.holderInstance, err = nftholder.NewNftholder(ch.holderContractAddr, ch.client)
	if err != nil{
		return nil, err
	}

	return ch, nil
}

func (n *nft1155Handler)retrieveTransactionOpts( nonce *big.Int)(opts *bind.TransactOpts, err error){
	chainID, err := n.client.NetworkID(context.Background())
	if err != nil {
		logger.Error(err)
		return
	}
	opts, err = bind.NewKeyedTransactorWithChainID(n.privateKey, chainID)
	if err != nil{
		logger.Error(err)
		return
	}
	if nonce != nil{
		opts.Nonce = nonce
	}else{
		var uint64Nonce uint64
		uint64Nonce, err = n.client.NonceAt(context.Background(), opts.From, nil)
		if err != nil{
			return
		}
		opts.Nonce = new(big.Int)
		opts.Nonce.SetUint64(uint64Nonce)
	}
	opts.Context = context.TODO()
	return
}

func (n *nft1155Handler)getWritingReceipt(ctx context.Context,
	method string, transaction *types.Transaction,
	response *pbNft1155.WritingMethodResponse)(err error){
	defer func() {
		logger.Infof("transactionHash=%s", transaction.Hash())
	}()
	if transaction.Hash().String() == ""{
		return errors.New("cannot get the transaction hash")
	}
	var receipt types.Receipt
	err = n.rpcClient.CallContext(ctx, &receipt,
		"eth_getTransactionReceipt", transaction.Hash())
	if err != nil {
		return err
	}
	response.State = receipt.Status

	return nil
}

func (n nft1155Handler) Name(ctx context.Context,
	request *pbNft1155.NameRequest,
	response *pbNft1155.NameResponse) error {
	logger.Info("calling Name")
	var err error
	response.Name, err = n.contractInstance.Name(nil)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (n nft1155Handler) Symbol(ctx context.Context,
	request *pbNft1155.SymbolRequest,
	response *pbNft1155.SymbolResponse) error {
	logger.Info("calling Symbol")
	var err error
	response.Symbol, err = n.contractInstance.Symbol(nil)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (n nft1155Handler) Owner(ctx context.Context,
	request *pbNft1155.OwnerRequest,
	response *pbNft1155.OwnerResponse) error {
	defer func(){
		logger.Infof("calling Owner %+v %+v", request, response)
	}()
	address, err := n.contractInstance.Owner(nil)
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Infof("%s\n", address.String())
	response.OwnerAddress = address.String()
	return nil
}

func (n nft1155Handler) Uri(ctx context.Context,
	request *pbNft1155.UriRequest,
	response *pbNft1155.UriResponse) error {
	logger.Info("calling Uri")

	bi := utils.Bytes2bigInt(request.Id)
	uri, err := n.contractInstance.Uri(nil, bi)
	if err != nil{
		logger.Error(err)
		return err
	}
	response.Uri = uri
	return nil
}

func (n nft1155Handler) TokenSupply(ctx context.Context,
	request *pbNft1155.TokenSupplyRequest,
	response *pbNft1155.TokenSupplyResponse) error {
	logger.Info("calling TokenSupply")
	bi := utils.Bytes2bigInt(request.Id)
	resp, err := n.contractInstance.TokenSupply(nil, bi)
	if err != nil{
		return err
	}
	response.Amount = resp.Bytes()
	return nil
}

func (n nft1155Handler) BalanceOf(ctx context.Context,
	request *pbNft1155.BalanceOfRequest,
	response *pbNft1155.BalanceOfResponse) error {
	logger.Info("calling BalanceOf")
	address := common.HexToAddress(string(request.OwnerAddress))
	bi := utils.Bytes2bigInt(request.Id)
	resp, err := n.contractInstance.BalanceOf(nil, address, bi)
	if err != nil{
		logger.Error(err)
		return err
	}
	logger.Info(resp.String())
	response.Balance = utils.BigInt2Bytes(resp)
	return nil
}

func (n nft1155Handler) BalanceOfBatch(ctx context.Context,
	request *pbNft1155.BalanceOfBatchRequest,
	response *pbNft1155.BalanceOfBatchResponse) error {
	logger.Info("calling BalanceOfBatch", request.String())
	var (
		addresses []common.Address
		ids []*big.Int
	)
	for i:=0; i<len(request.Owners); i++{
		addresses = append(addresses, common.HexToAddress(request.Owners[i].OwnerAddress))
		ids = append(ids, utils.Bytes2bigInt(request.Owners[i].Id))
	}
	resp, err := n.contractInstance.BalanceOfBatch(nil, addresses, ids)
	if err != nil{
		return err
	}
	logger.Info(resp)
	for i:=0; i<len(resp); i++{
		response.Balances = append(response.Balances, &pbNft1155.BalanceOfResponse{
			OwnerAddress: request.Owners[i].GetOwnerAddress(),
			Balance:      utils.BigInt2Bytes(resp[i]),
		})
	}
	return nil
}

func (n nft1155Handler) GetNextTokenID(ctx context.Context,
	request *pbNft1155.GetNextTokenIDRequest,
	response *pbNft1155.GetNextTokenIDResponse) error {
	defer func(){
		logger.Infof("calling GetNextTokenID, request=%+v, response=%+v",
			request, response)
	}()
	id, err := n.contractInstance.GetNextTokenID(nil)
	if err != nil{
		logger.Error(err)
		return err
	}
	response.Id = utils.BigInt2Bytes(id)
	return nil
}

func (n nft1155Handler) IsApprovedForAll(ctx context.Context,
	request *pbNft1155.IsApprovedForAllRequest,
	response *pbNft1155.IsApprovedForAllResponse) error {
	logger.Info("calling IsApprovedAll")

	account := common.HexToAddress(request.OwnerAddress)
	operator := common.HexToAddress(request.OperatorAddress)
	isOperator, err := n.contractInstance.IsApprovedForAll(nil, account, operator)
	if err != nil{
		logger.Error(err)
		return err
	}
	response.IsOperator = isOperator
	return nil
}

func (n nft1155Handler) SafeTransferFrom(ctx context.Context, request *pbNft1155.SafeTransferFromRequest, response *pbNft1155.WritingMethodResponse) error {
	panic("implement me")
}

func (n nft1155Handler) SafeBatchTransferFrom(ctx context.Context,
	request *pbNft1155.SafeBatchTransferFromRequest,
	response *pbNft1155.WritingMethodResponse) error {
	methodName := "SafeBatchTransferFrom"
	var err error
	defer func() {
		logger.Infof("calling SafeBatchTransferFrom. request=%+v response=%+v, err=%+v",
			request, response, err)
	}()
	ids := utils.ArrayBytes2BigInts(request.Ids)
	amounts := utils.ArrayBytes2BigInts(request.Amounts)

	opts, err := n.retrieveTransactionOpts(nil)
	if err != nil{
		return err
	}

	transaction, err := n.contractInstance.SafeBatchTransferFrom(opts,
		common.HexToAddress(request.FromAddress),
		common.HexToAddress(request.ToAddress),
		ids,
		amounts,
		request.Data)
	if err != nil{
		return err
	}
	return n.getWritingReceipt(ctx, methodName, transaction, response)
}

func (n nft1155Handler) SetApprovalForAll(ctx context.Context, request *pbNft1155.SetApprovalForAllRequest, response *pbNft1155.WritingMethodResponse) error {
	panic("implement me")
}

func (n nft1155Handler) SetBaseMetadataURI(ctx context.Context, request *pbNft1155.SetBaseMetadataURIRequest, response *pbNft1155.WritingMethodResponse) error {
	var (
		err error
		methodName = "setBaseMetadataURI"
	)
	defer func(){
		logger.Infof("calling %s request=%+v, response=%+v, %v", methodName, request, response, err)
	}()

	opts, err := n.retrieveTransactionOpts(nil)
	if err != nil{
		return err
	}
	newTransaction, err := n.contractInstance.SetBaseMetadataURI(opts,
		request.BaseMetadataURI)
	if err != nil{
		return err
	}
	return n.getWritingReceipt(ctx, methodName, newTransaction, response)
}

func (n nft1155Handler) Create(ctx context.Context,
	request *pbNft1155.CreateRequest,
	response *pbNft1155.WritingMethodResponse) error {
	var (
		err error
		methodName = "create"
	)
	defer func(){
		logger.Infof("calling %s request=%+v, response=%+v, %v",
			methodName, request, response, err)
	}()

	opts, err := n.retrieveTransactionOpts(nil)
	if err != nil{
		return err
	}

	initOwnerAddress := common.HexToAddress(string(request.InitOwnerAddress))
	initSupply := utils.Bytes2bigInt(request.InitSupply)
	uri := request.Uri
	data := request.Data
	newTransaction, err := n.contractInstance.Create(opts,
		initOwnerAddress, initSupply, uri, data)
	if err != nil{
		return err
	}
	response.TransactionHash = newTransaction.Hash().String()
	response.Method = methodName

	return n.getWritingReceipt(ctx, methodName, newTransaction, response)
}

func (n nft1155Handler) CreateBatch(ctx context.Context,
	request *pbNft1155.CreateBatchRequest,
	response *pbNft1155.WritingMethodResponse) error {
	var (
		err error
		methodName = "createBatch"
	)
	func(){
		logger.Infof("calling %s request=%+v, response=%+v, %v",
			methodName, request, response, err)
	}()
	opts, err := n.retrieveTransactionOpts(nil)
	if err != nil{
		return err
	}

	var newTransaction *types.Transaction
	newTransaction, err = n.contractInstance.CreateBatch(opts,
		common.HexToAddress(string(request.InitOwnerAddress)),
		utils.ArrayBytes2BigInts(request.Quantities), request.Uris, request.Data )
	if err != nil{
		logger.Error(err)
		return err
	}
	response.TransactionHash = newTransaction.Hash().String()
	response.Method = methodName

	return n.getWritingReceipt(ctx, methodName, newTransaction, response)
}

func (n nft1155Handler) TransferGovernorship(ctx context.Context,
	request *pbNft1155.TransferGovernorshipRequest,
	response *pbNft1155.WritingMethodResponse) error {
	var (
		err error
		methodName="transferOwnership"
	)
	defer func(){
		logger.Infof("calling %s request=%+v, response=%+v, %v",
			methodName, request, response, err)
	}()
	opts, err := n.retrieveTransactionOpts(nil)
	if err != nil{
		return err
	}
	newTransaction, err := n.contractInstance.TransferOwnership(
		opts,common.HexToAddress(request.NewGovernorAddress))
	if err != nil{
		return err
	}

	return n.getWritingReceipt(ctx, methodName, newTransaction, response)
}

func (n nft1155Handler) Mint(ctx context.Context, request *pbNft1155.MintRequest,
	response *pbNft1155.WritingMethodResponse) error {
	var (
		err error
		methodName = "mint"
	)
	defer func(){
		logger.Infof("calling %s request=%+v, response=%+v, %v",
			methodName, request, response, err)
	}()
	opts, err := n.retrieveTransactionOpts(nil)
	if err != nil{
		return err
	}

	newTransaction, err := n.contractInstance.Mint(opts,
		common.HexToAddress(string(request.AddressTo)),
		utils.Bytes2bigInt(request.Id),
		utils.Bytes2bigInt(request.Quantity),
		request.Data)
	if err != nil{
		return err
	}

	return n.getWritingReceipt(ctx, methodName, newTransaction, response)
}

func (n nft1155Handler) MintBatch(ctx context.Context,
	request *pbNft1155.MintBatchRequest,
	response *pbNft1155.WritingMethodResponse) error {
	var (
		err error
		methodName = "mintBatch"
	)
	defer func(){
		logger.Infof("calling %s request=%+v, response=%+v, %v",
			methodName, request, response, err)
	}()
	opts, err := n.retrieveTransactionOpts(nil)
	if err != nil{
		return err
	}

	newTransaction, err := n.contractInstance.MintBatch(
		opts,common.HexToAddress(request.AddressTo),
		utils.ArrayBytes2BigInts(request.Quantities),
		utils.ArrayBytes2BigInts(request.Ids),
		request.Data)
	if err != nil{
		return err
	}

	return n.getWritingReceipt(ctx, methodName, newTransaction, response)
}

func (n nft1155Handler) SetCreator(ctx context.Context,
	request *pbNft1155.SetCreatorRequest,
	response *pbNft1155.WritingMethodResponse) error {
	var (
		err error
		methodName = "setCreator"
	)
	defer func(){
		logger.Infof("calling %s request=%+v, response=%+v, %v",
			methodName, request, response, err)
	}()

	opts, err := n.retrieveTransactionOpts(nil)
	if err != nil{
		return err
	}

	newTransaction, err := n.contractInstance.SetCreator(
		opts,common.HexToAddress(request.AddressTo),
		utils.ArrayBytes2BigInts(request.Ids) )
	if err != nil{
		return err
	}

	return n.getWritingReceipt(ctx, methodName, newTransaction, response)
}

func (n nft1155Handler) SetIdURI(ctx context.Context,
	request *pbNft1155.SetIdURIRequest,
	response *pbNft1155.WritingMethodResponse) error {
	var (
		err error
		methodName = "setIdURI"
	)
	defer func(){
		logger.Infof("calling %s request=%+v, response=%+v, %v",
			methodName, request, response, err)
	}()

	opts, err := n.retrieveTransactionOpts(nil)
	if err != nil{
		return err
	}
	newTransaction, err := n.contractInstance.SetIdURI(opts,
		utils.Bytes2bigInt(request.Id), request.Uri)
	if err != nil{
		return err
	}

	return n.getWritingReceipt(ctx, methodName, newTransaction, response)
}

func (n nft1155Handler) SetIdURIs(ctx context.Context,
	request *pbNft1155.SetIdURIsRequest,
	response *pbNft1155.WritingMethodResponse) error {
	var (
		err error
		methodName = "setIdURIs"
	)
	defer func(){
		logger.Infof("calling %s request=%+v, response=%+v, %v",
			methodName, request, response, err)
	}()
	opts, err := n.retrieveTransactionOpts(nil)
	if err != nil{
		return err
	}
	newTransaction, err := n.contractInstance.SetIdURIs(opts, utils.ArrayBytes2BigInts(request.Ids),
		request.Uris)
	if err != nil{
		return err
	}

	return n.getWritingReceipt(ctx, methodName, newTransaction, response)
}


func (h *nft1155Handler) getContractGasFee(ctx context.Context,
	inputData []byte, contractAddr common.Address, response *pbNft1155.GasFeeResponse) error {
	opts, err := h.retrieveTransactionOpts(nil)
	if err != nil{
		return err
	}

	if code, err := h.client.PendingCodeAt(ctx, contractAddr); err != nil {
		return err
	} else if len(code) == 0 {
		return errors.New("PendingCodeAt return noCode")
	}

	gas, gasLimit, err := getGasFee(ctx, h.client,
		opts.From,
		contractAddr, inputData)
	if err != nil{
		return fmt.Errorf("execute getGastFee error: %v", err)
	}
	logger.Info("gas=", gas.String(), "gasLimit=", gasLimit)

	response.GasLimit = gasLimit
	response.GasPrice = gas.Bytes()
	return nil
}

func (n *nft1155Handler) GetCreateBatchPrice(ctx context.Context,
	request *pbNft1155.CreateBatchRequest,
	response *pbNft1155.GasFeeResponse) (err error) {
	defer func(){
		logger.Infof("calling GetGasFee, request=%+v, response=%+v, err=%v",
			request, response, err)
	}()

	initOwnerAddr := common.HexToAddress(request.InitOwnerAddress)
	inputData, err := nft1155ABI.Pack("createBatch", initOwnerAddr,
		utils.ArrayBytes2BigInts(request.Quantities),
		request.Uris,
		request.Data)
	if err != nil {
		return fmt.Errorf("execute pack error: %v", err)
	}

	return n.getContractGasFee(ctx, inputData, n.contractAddress, response)
}

func (n *nft1155Handler) TransferBatchID(ctx context.Context,
	req *pbNft1155.TransferBatchIDReq,
	response *pbNft1155.WritingMethodResponse) (err error){
	var (
		methodName = "transferBatchIdERC1155"
	)
	func(){
		logger.Infof("calling %s request=%+v, response=%+v, %v",
			methodName, req, response, err)
	}()
	opts, err := n.retrieveTransactionOpts(nil)
	if err != nil{
		return
	}
	var newTransaction *types.Transaction
	newTransaction, err = n.holderInstance.TransferBatchIdERC1155(opts,
		common.HexToAddress(req.NftContractAddr),
		common.HexToAddress(req.ToAddress),
		utils.ArrayBytes2BigInts(req.Ids),
		utils.ArrayBytes2BigInts(req.Quantities))
	if err != nil{
		logger.Error(err)
		return
	}
	logger.Infof("%+v, hash=%s err=%v",
		newTransaction, newTransaction.Hash().String(), err)
	response.TransactionHash = newTransaction.Hash().String()
	response.Method = methodName

	return n.getWritingReceipt(ctx, methodName, newTransaction, response)
}

func (n *nft1155Handler) GetTransferBatchIDPrice(ctx context.Context, req *pbNft1155.TransferBatchIDReq, response *pbNft1155.GasFeeResponse) (err error) {
	defer func(){
		logger.Infof("calling GetTransferBatchIDPrice, request=%+v, response=%+v, err=%v",
			req, response, err)
	}()
	methodName := "transferBatchIdERC1155"

	inputData, err := holderABI.Pack(methodName,
		common.HexToAddress(req.NftContractAddr),
		common.HexToAddress(req.ToAddress),
		utils.ArrayBytes2BigInts(req.Ids),
		utils.ArrayBytes2BigInts(req.Quantities))
	if err != nil {
		return fmt.Errorf("execute pack error: %v", err)
	}

	return n.getContractGasFee(ctx, inputData, n.holderContractAddr, response)
}

func (n *nft1155Handler) TransferERC20(ctx context.Context,
	request *pbNft1155.TransferERC20Request,
	response *pbNft1155.WritingMethodResponse) (err error) {
	var (
		methodName = "transferERC20"
	)
	func(){
		logger.Infof("calling %s request=%+v, response=%+v, %v",
			methodName, request, response, err)
	}()
	opts, err := n.retrieveTransactionOpts(nil)
	if err != nil{
		logger.Error(err)
		return
	}

	opsDecmals, err := opsDecimalsFunc()
	if err != nil{
		logger.Error(err)
		return err
	}
	//base := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(opsDecmals)), nil)
	//bigAmount := new(big.Int).Exp(bytes2bigInt(request.Amount), base, nil)
	var newTransaction *types.Transaction
	newTransaction, err = n.holderInstance.TransferERC20(opts,
		common.HexToAddress(request.TokenContract),
		common.HexToAddress(request.AddressTo),
		utils.FloatStringToBigInt(request.Amount, int(opsDecmals)),
		)
	if err != nil{
		logger.Error(err)
		return
	}
	logger.Infof("%+v, hash=%s err=%v",
		newTransaction, newTransaction.Hash().String(), err)
	response.TransactionHash = newTransaction.Hash().String()
	response.Method = methodName

	return n.getWritingReceipt(ctx, methodName, newTransaction, response)
}


func (n *nft1155Handler) GetTransferERC20Price(ctx context.Context,
	request *pbNft1155.TransferERC20Request,
	response *pbNft1155.GasFeeResponse)(err error){
	defer func(){
		logger.Infof("calling GetGasFee, request=%+v, response=%+v, err=%v",
			request, response, err)
	}()

	methodName := "transferERC20"

	decimals, err := opsDecimalsFunc()
	if err != nil{
		return err
	}

	inputData, err := holderABI.Pack(methodName,
		common.HexToAddress(request.TokenContract),
		common.HexToAddress(request.AddressTo),
		utils.FloatStringToBigInt(request.Amount, int(decimals)),
	)
	if err != nil {
		return fmt.Errorf("execute pack error: %v", err)
	}

	return n.getContractGasFee(ctx, inputData, n.holderContractAddr, response)
}

func (n *nft1155Handler) GetTransactionByHash(ctx context.Context, req *pbNft1155.GetTransactionByHashReq, response *pbNft1155.GetTransactionByHashResponse) error {
	var err error
	defer func(){
		logger.Infof("calling GetTransactionByHash, request=%+v, response=%+v, err=%v",
			req, response, err)
	}()
	tx, isPending, err := n.client.TransactionByHash(ctx,
		common.HexToHash(req.TransactionHash) )
	if err != nil{
		return err
	}

	if tx.To().String() == globalOptions.NFT1155ContractAddr.String(){
		output := make(map[string]interface{})
		m, er := nft1155ABI.MethodById(tx.Data()[:4])
		if er != nil {
			logger.Error(er)
			return er
		}
		if er = m.Inputs.UnpackIntoMap(output, tx.Data()[4:]); er != nil {
			logger.Error(er)
			return er
		}
		logger.Infof("%+v", output)
		amounts := output["amounts"].([]*big.Int)
		ids := output["ids"].([]*big.Int)
		data := output["data"].([]byte)
		fromAddr := output["from"].(common.Address)
		toAddr := output["to"].(common.Address)

		response.ContractAddress = globalOptions.NFT1155ContractAddr.String()
		response.ContractFrom = fromAddr.String()
		response.ContractTo = toAddr.String()
		for i := range ids {
			response.IdAndAmount = append(response.IdAndAmount, &pbNft1155.IdAmount{
				Id:     ids[i].Bytes(),
				Amount: amounts[i].Bytes(),
			})
		}
		response.Data = data
	}else if tx.To().String() == globalOptions.HolderContractAddr.String() {
		output := make(map[string]interface{})
		m, er := holderABI.MethodById(tx.Data()[:4])
		if er != nil {
			logger.Error(er)
			return er
		}
		if er = m.Inputs.UnpackIntoMap(output, tx.Data()[4:]); er != nil {
			logger.Error(er)
			return er
		}
		logger.Infof("%+v", output)
		amounts := output["_quantities"].([]*big.Int)
		ids := output["_ids"].([]*big.Int)
		fromAddr := output["nftContract"].(common.Address)
		toAddr := output["to"].(common.Address)

		response.ContractAddress = globalOptions.HolderContractAddr.String()
		response.ContractFrom = fromAddr.String()
		response.ContractTo = toAddr.String()
		for i := range ids {
			response.IdAndAmount = append(response.IdAndAmount, &pbNft1155.IdAmount{
				Id:     ids[i].Bytes(),
				Amount: amounts[i].Bytes(),
			})
		}
	}


	response.Type =     uint32(tx.Type())
	response.Nonce =    tx.Nonce()
	response.GasPrice = tx.GasPrice().Bytes()
	response.Gas =      tx.Gas()
	response.EthValue= tx.Value().Bytes()
	response.EthTo=       tx.To().String()
	response.Hash=     tx.Hash().String()
	response.IsPending = isPending

	if isPending == false{ //if the transaction is not pending, we can get the receipt exactly
		var r *types.Receipt
		r , err = n.client.TransactionReceipt(ctx, tx.Hash())
		if err != nil{
			logger.Error(err)
			response.Status = 0
		}else{
			response.Status = uint32(r.Status)
			logger.Error(err)
		}
	}
	logger.Info(response)
	return nil
}
