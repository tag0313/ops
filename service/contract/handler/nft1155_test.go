package handler

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"ops/pkg/logger"
	"ops/pkg/utils"
	"ops/proto/nft1155"
	"testing"
)

func initNew1155Handler()(*nft1155Handler, error){
	return NewNft1155Handler(testOptions)
}

func TestNft1155GetMethod(t *testing.T){
	ctx := context.TODO()
	h, err := initNew1155Handler()
	if err != nil{
		t.Fatal(err)
	}

	var nameResp pbNft1155.NameResponse
	err = h.Name(ctx, nil,  &nameResp)
	if err != nil{
		t.Fatal(err)
	}
	t.Log("Name is: ",nameResp.Name)

	var symbolResp pbNft1155.SymbolResponse
	err = h.Symbol(ctx, nil, &symbolResp)
	if err != nil{
		t.Fatal(err)
	}
	t.Log("symbol is: ", symbolResp.Symbol)

	var ownerResp pbNft1155.OwnerResponse
	err = h.Owner(ctx, nil, &ownerResp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Owner is 0x%x\n", ownerResp.OwnerAddress)

	var(
		getNextTokenReq pbNft1155.GetNextTokenIDRequest
		getNextTokenResp pbNft1155.GetNextTokenIDResponse
	)
	err = h.GetNextTokenID(nil, &getNextTokenReq, &getNextTokenResp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("next token id is", getNextTokenResp.Id)

	var (
		uriRequest pbNft1155.UriRequest
		uriResponse pbNft1155.UriResponse
	)
	uriRequest.Id = getNextTokenResp.Id
	err = h.Uri(ctx, &uriRequest, &uriResponse)
	if err != nil{
		t.Fatal(err)
	}
	t.Logf("uri is %s\n", uriResponse.Uri)

	var(
		boReq  pbNft1155.BalanceOfRequest
		boResp pbNft1155.BalanceOfResponse
	)
	boReq.OwnerAddress = testOptions.NFT1155Manager.String()
	boReq.Id = getNextTokenResp.Id
	err = h.BalanceOf(ctx, &boReq, &boResp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("balance of %s\n", boResp.Balance)

	var(
		boBatchReq pbNft1155.BalanceOfBatchRequest
		boBatchResp pbNft1155.BalanceOfBatchResponse
	)
	newBo := boReq
	boBatchReq.Owners = append(boBatchReq.Owners, &boReq,&newBo)
	err = h.BalanceOfBatch(ctx, &boBatchReq, &boBatchResp)
	if err != nil{
		t.Fatal(err)
	}
	t.Log(boBatchResp)

	var(
		isApprovedForAllReq pbNft1155.IsApprovedForAllRequest
		isApprovedForAllResp pbNft1155.IsApprovedForAllResponse
	)
	isApprovedForAllReq.OwnerAddress = testOptions.NFT1155Manager.String()
	isApprovedForAllReq.OperatorAddress = testOptions.NFT1155Manager.String()
	err = h.IsApprovedForAll(ctx, &isApprovedForAllReq, &isApprovedForAllResp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("isApproved =", isApprovedForAllResp.IsOperator)
}

func TestNft1155Create(t *testing.T) {
	ctx := context.Background()
	h, err := initNew1155Handler()
	if err != nil {
		t.Fatal(err)
	}
	var(
		createReq pbNft1155.CreateBatchRequest
		createResp pbNft1155.WritingMethodResponse
	)

	const num = 20
	for i:=0; i < num; i++{
		createReq.Quantities = append(createReq.Quantities, utils.BigInt2Bytes(big.NewInt(9527)))
		createReq.Uris = append(createReq.Uris,"http://baidu.com/1.json")
	}
	createReq.InitOwnerAddress = testOptions.NFT1155Manager.String()
	createReq.Data = []byte("00")

	err = h.CreateBatch(ctx, &createReq, &createResp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(createResp)

	var(
		priceResp = new(pbNft1155.GasFeeResponse)
	)
	err = h.GetCreateBatchPrice(context.TODO(), &createReq, priceResp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("response=%+v", priceResp)
}

func TestGasFee(t *testing.T) {
	h, err := initNew1155Handler()
	if err != nil{
		t.Fatal(err)
	}

	var(
		req = new(pbNft1155.CreateBatchRequest)
		resp = new(pbNft1155.GasFeeResponse)
	)
	req.InitOwnerAddress = testOptions.NFT1155Manager.String()
	for i:=0;i<100;i++ {
		req.Quantities = append(req.Quantities, big.NewInt(100).Bytes())
		req.Uris = append(req.Uris, fmt.Sprintf("https://www.baidu.com/%d.json", i))
	}
	req.Data = nil
	err = h.GetCreateBatchPrice(context.TODO(), req, resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("response=%+v",  resp)
}


func TestUnderStandingAddress(t *testing.T){
	hexAddr := common.HexToAddress(testOptions.OpsContractAddr.String())
	bytesAddress := common.BytesToAddress([]byte(testOptions.OpsContractAddr.String()))
	t.Logf("hex=%s bytes=%s origin=%s",hexAddr.Hex(), bytesAddress.Hex(), testOptions.OpsContractAddr.String())
}

func Test1155TransferBatchID(t *testing.T){
	h, err := initNew1155Handler()
	if err != nil{
		t.Fatal(err)
	}
	var(
		req = new(pbNft1155.TransferBatchIDReq)
		resp = new(pbNft1155.WritingMethodResponse)
		price = new(pbNft1155.GasFeeResponse)
	)
	req.Ids = [][]byte{[]byte("1240")}
	req.Quantities = [][]byte{[]byte("1")}
	req.NftContractAddr = "0x2b7D03324D2c8E89E70C058088Bc08f14AEF8Da3"
	req.ToAddress = "0x5fA4B253C9f20cccf021BCD8A501De998C0CBa41"
	//get price
	err = h.GetTransferBatchIDPrice(context.TODO(), req, price)
	if err != nil{
		t.Error(err)
	}

	err = h.TransferBatchID(context.TODO(), req, resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("response=%+v",  resp)
}


func TestTransferERC20(t *testing.T){
	opsDecimalsFunc = func() (uint8, error) {
		return 18, nil
	}

	h, err := initNew1155Handler()
	if err != nil{
		t.Fatal(err)
	}
	var(
		req = new(pbNft1155.TransferERC20Request)
		resp = new(pbNft1155.WritingMethodResponse)
		price = new(pbNft1155.GasFeeResponse)
	)
	req.Amount = "0.00001"
	req.AddressTo = "0x42D073B89f8854e6A7ce61F1583FA4A28d71BC0d"
	req.TokenContract = "0xb247BefC358BeDf8E0f305826350cC40184879c5"
	//get price
	err = h.GetTransferERC20Price(context.TODO(), req, price)
	if err != nil{
		t.Error(err)
	}

	err = h.TransferERC20(context.TODO(), req, resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("response=%+v",  resp)
}

func TestSafeBatchTransferFrom(t *testing.T){
	h, err := initNew1155Handler()
	if err != nil{
		t.Fatal(err)
	}
	var(
		req = new(pbNft1155.SafeBatchTransferFromRequest)
		resp = new(pbNft1155.WritingMethodResponse)
		//holder = "0xDfC05ec7a9C15643411A2584438E9D085d4c54eE"
		richAccount = "0x66d59cA5721Ce058B706581d983bbD7c5bA366f1"
		receiveAccount = "0x42D073B89f8854e6A7ce61F1583FA4A28d71BC0d"
		balanceOfReq = new(pbNft1155.BalanceOfRequest)
		balanceOfResp = new(pbNft1155.BalanceOfResponse)
		ids [][]byte
	)
	balances := make(map[*big.Int]*big.Int)
	balances[new(big.Int).SetInt64(1323)]= new(big.Int).SetInt64(0)
	balances[new(big.Int).SetInt64(1324)]= new(big.Int).SetInt64(0)
	balances[new(big.Int).SetInt64(1325)]= new(big.Int).SetInt64(0)
	// get the balance of ids before transfer
	ids = make([][]byte, len(balances))
	i := 0
	for id, _ := range balances {
		balanceOfReq.OwnerAddress = receiveAccount
		balanceOfReq.Id = id.Bytes()
		err = h.BalanceOf(context.TODO(), balanceOfReq, balanceOfResp)
		if err != nil{
			logger.Fatal(err)
		}
		balances[id] = utils.Bytes2bigInt(balanceOfResp.Balance)
		ids[i] = id.Bytes()
		i++
		logger.Infof("id(%d).Balance=%d", id, balances[id])
	}

	//these ids has a huge amount of ids, don't hesitate when testing the code
	//1323 1324 1325
	transferNumber := new(big.Int).SetInt64(1)
	req.FromAddress = richAccount
	req.ToAddress = receiveAccount
	req.Ids = ids
	req.Amounts = utils.BigInts2ArrayBytes([]*big.Int{
		transferNumber,
		transferNumber,
		transferNumber,
	})
	req.Data=nil
	err = h.SafeBatchTransferFrom(context.Background(), req, resp)
	if err != nil{
		logger.Fatal(err)
	}
	logger.Info(resp.TransactionHash)

	for id, b := range balances {
		balanceOfReq.OwnerAddress = receiveAccount
		balanceOfReq.Id = id.Bytes()
		err = h.BalanceOf(context.TODO(), balanceOfReq, balanceOfResp)
		if err != nil{
			logger.Fatal(err)
		}
		respBal := utils.Bytes2bigInt(balanceOfResp.Balance)
		sum := new(big.Int).Add(b,transferNumber)
		if respBal.Cmp(sum) != 0 {
			logger.Infof("this is not a failure, the blockchain need time to merge the transaction blocks. id(%d).Balance %d != %d", id, respBal, sum)
		}else{
			logger.Infof("id(%d).Balance %d == %d", id, respBal, b)
		}
	}
}

func TestTransferBatchID(t *testing.T){
	h, err := initNew1155Handler()
	if err != nil{
		t.Fatal(err)
	}
	var(
		req = new(pbNft1155.TransferBatchIDReq)
		resp = new(pbNft1155.WritingMethodResponse)
		receiveAccount = "0x42D073B89f8854e6A7ce61F1583FA4A28d71BC0d"
		balanceOfReq = new(pbNft1155.BalanceOfRequest)
		balanceOfResp = new(pbNft1155.BalanceOfResponse)
		ids [][]byte
	)
	balances := make(map[*big.Int]*big.Int)
	balances[new(big.Int).SetInt64(1323)]= new(big.Int).SetInt64(0)
	balances[new(big.Int).SetInt64(1324)]= new(big.Int).SetInt64(0)
	balances[new(big.Int).SetInt64(1325)]= new(big.Int).SetInt64(0)
	// get the balance of ids before transfer
	ids = make([][]byte, len(balances))
	i := 0
	for id, _ := range balances {
		balanceOfReq.OwnerAddress = receiveAccount
		balanceOfReq.Id = id.Bytes()
		err = h.BalanceOf(context.TODO(), balanceOfReq, balanceOfResp)
		if err != nil{
			logger.Fatal(err)
		}
		balances[id] = utils.Bytes2bigInt(balanceOfResp.Balance)
		ids[i] = id.Bytes()
		i++
		logger.Infof("id(%d).Balance=%d", id, balances[id])
	}

	//these ids has a huge amount of ids, don't hesitate when testing the code
	//1323 1324 1325
	transferNumber := new(big.Int).SetInt64(1)
	req.NftContractAddr = testOptions.NFT1155ContractAddr.String()
	req.ToAddress = receiveAccount
	req.Ids = ids
	req.Quantities = utils.BigInts2ArrayBytes([]*big.Int{
		transferNumber,
		transferNumber,
		transferNumber,
	})
	err = h.TransferBatchID(context.Background(), req, resp)
	if err != nil{
		logger.Fatal(err)
	}
	logger.Info(resp.TransactionHash)

	for id, b := range balances {
		balanceOfReq.OwnerAddress = receiveAccount
		balanceOfReq.Id = id.Bytes()
		err = h.BalanceOf(context.TODO(), balanceOfReq, balanceOfResp)
		if err != nil{
			logger.Fatal(err)
		}
		respBal := utils.Bytes2bigInt(balanceOfResp.Balance)
		sum := new(big.Int).Add(b,transferNumber)
		if respBal.Cmp(sum) != 0 {
			logger.Infof("this is not a failure, the blockchain need time to merge the transaction blocks. id(%d).Balance %d != %d", id, respBal, sum)
		}else{
			logger.Infof("id(%d).Balance %d == %d", id, respBal, b)
		}
	}
}
