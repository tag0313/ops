package handler

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"ops/pkg/logger"
	"ops/pkg/utils"
	"ops/proto/ethereum"
)


type ethHandler struct {
	client *ethclient.Client
}

func NewEthHandler(opts *Options)(*ethHandler, error){
	var(
		err error
		h = new(ethHandler)
	)
	h.client, err = ethclient.Dial(opts.BlockChainURLHttp)
	if err != nil{
		return nil, err
	}
	return h, err
}


func (h ethHandler) BalanceETH(ctx context.Context, request *pbEthereum.ETHBalanceRequest, response *pbEthereum.ETHBalanceResponse) error {
	var err error
	defer func(){
		logger.Infof("calling BalanceETH request=%+v,  response=%+v,  err=%v", request, response, err)
	}()

	address := common.HexToAddress(request.Address)
	resp, err := h.client.BalanceAt(ctx, address, nil)
	if err != nil{
		return err
	}
	response.Balance = utils.BigInt2Bytes(resp)
	response.Address = request.Address
	response.Decimals = 18
	return nil
}

func sendEthTransaction(ctx context.Context,
	ethRpcClient *rpc.Client,
	privateKey *ecdsa.PrivateKey,
	toAddr common.Address,
	amount *big.Rat)(transactionHash string,nonce uint64, err error){
	chainID, err := ethereumClient.NetworkID(context.Background())
	if err != nil {
		return
	}
	gasPrice, err := ethereumClient.SuggestGasPrice(ctx)
	if err != nil{
		return
	}

	from :=crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err = ethereumClient.NonceAt(context.Background(), from, nil)
	if err != nil {
		return
	}

	pow := utils.BigIntPow10(18)
	ethAmount := new(big.Rat).Mul(amount, new(big.Rat).SetInt(pow))
	tx := types.NewTx(&types.AccessListTx{
		ChainID:  chainID,
		Nonce:    nonce,
		To:       &toAddr,
		Value:    ethAmount.Num(),
		Gas:      21000,
		GasPrice: gasPrice,
		Data:     nil, //if it's not a contract, fill it as nil
	})
	signer := types.NewEIP2930Signer(chainID)
	signedTx, err := types.SignTx(tx,  signer, privateKey)
	if err!=nil{
		return
	}
	data, err := signedTx.MarshalBinary()
	if err != nil {
		return
	}

	var result interface{}
	err = ethRpcClient.CallContext(ctx, &result,"eth_sendRawTransaction",
		hexutil.Encode(data))
	if err != nil{
		logger.Error(err)
		return
	}
	transactionHash = result.(string)
	return
}