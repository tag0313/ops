package handler

import (
	"context"
	"crypto/ecdsa"
	"github.com/asim/go-micro/v3/client"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

var(
	pbClientG  client.Client        = nil
)

var opsDecimalsFunc = func()(uint8,error){
	return opsIns.Decimals(nil)
}

func SetPBClient(pbClient client.Client) {
	if pbClientG == nil {
		pbClientG = pbClient
	}
}

func retrieveTransactionOpts(privateKey *ecdsa.PrivateKey, client *ethclient.Client, nonce *big.Int) (opts *bind.TransactOpts, err error) {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return
	}
	opts, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return
	}
	if nonce != nil {
		opts.Nonce = nonce
	} else {
		var uint64Nonce uint64
		uint64Nonce, err = client.NonceAt(context.Background(), opts.From, nil)
		if err != nil {
			return
		}
		opts.Nonce = new(big.Int)
		opts.Nonce.SetUint64(uint64Nonce)
	}

	return
}

func getGasFee(ctx context.Context, ethClient *ethclient.Client, from,
	to common.Address, inputData []byte) (gasPrice *big.Int, gas uint64, err error) {
	gasPrice, err = ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return
	}
	var toAddress *common.Address
	if to.Hex() == "" {
		toAddress = nil //the destination contract (nil for contract creation)
	} else {
		toAddress = &to
	}
	msg := ethereum.CallMsg{
		From:     from,
		To:       toAddress,
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     inputData,
	}
	gas, err = ethClient.EstimateGas(ctx, msg)
	if err != nil {
		return
	}
	return
}

func getOpsTransactionFee(ctx context.Context,ethClient *ethclient.Client,
	fromAddress, toAddress common.Address, amount *big.Int) (gasPrice *big.Int, gas uint64, err error){
	inputData, err := opsABI.Pack("transfer", toAddress,
		amount)
	if err != nil {
		return
	}
	return getGasFee(ctx, ethClient, fromAddress, toAddress, inputData)
}