package contract

import (
	"ops/pkg/model"
)

// EthBalanceReq for request
type EthBalanceReq struct {
	Address string `json:"address" binding:"required"`
}

// EthBalanceResponse for response
type EthBalanceResponse struct {
	model.JSONResult
	Data EthBalanceData `json:"data"`
}
type EthBalanceData struct {
	BalanceData
}

type BalanceData struct {
	Balance  string `json:"balance" example:"1"`
	Decimals uint32 `json:"decimals" example:"2"`
	Address  string `json:"address" example:"0x53080E7e6d529e1040fC8Db4dc503705C31D384A"`
}
