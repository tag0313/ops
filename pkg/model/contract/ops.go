package contract

import (
	"errors"
	"math/big"
	"ops/pkg/model"
)

type OpsInfo struct {
	model.JSONResult
	Data InfoData `json:"data"`
}

type InfoData struct {
	Name        string `json:"name"`
	TotalSupply string `json:"total_supply"`
	Owner       string `json:"owner"`
	Decimals    uint32 `json:"decimals"`
	Symbol      string `json:"symbol"`
}

type AmountDataType struct {
	Amount string `json:"amount" binding:"required" example:"1.0000"`
}

func AmountToFloatString(dataType AmountDataType)(string, error){
	f := new(big.Float)
	a, b := f.SetString(dataType.Amount)
	if(!b){
		return "", errors.New("invalid float string")
	}
	return a.String(),nil
}

// OpsBalanceReq for request
type OpsBalanceReq struct {
	Address string `json:"address" binding:"required"`
}

// OpsBalanceResponse for response
type OpsBalanceResponse struct{
	model.JSONResult
	Data OpsBalanceData `json:"data"`
}
type OpsBalanceData struct{
	BalanceData
}

// SubscribeOPSRechargeReq for request
type SubscribeOPSRechargeReq struct {
	TransactionHash string `json:"transaction_hash" binding:"required"`
}
type SubscribeOPSRechargeResponse struct{
	model.JSONResult
	Data model.TransactionData `json:"data"`
}

// GetGasFeeReq for request
type GetGasFeeReq struct {
	FromAddress string `json:"from_address"`
	ToAddress string `json:"to_address"`
	model.AmountDataType
}

// GetGasFeeResponse for response
type GetGasFeeResponse struct{
	model.JSONResult
	Data GetGasFeeData `json:"data"`
}
type GetGasFeeData struct{
	Money string `json:"money"`
}

// WithdrawFeeReq for request
type WithdrawFeeReq struct {
	ToAddress string `json:"to_address" binding:"required"`
	model.AmountDataType
}

// WithdrawOpsReq for request
type WithdrawOpsReq struct {
	TokenAddress string `json:"token_address"`
	ToAddress string `json:"to_address" binding:"required"`
	model.AmountDataType
	GasFee string `json:"gas_fee" binding:"required"`
}

// WithdrawOpsResponse for response
type WithdrawOpsResponse struct{
	model.JSONResult
	Data WithdrawOpsData `json:"data"`
}

type WithdrawOpsData struct{
	TransactionHash string `json:"transaction_hash"`
}