package lp

import "ops/pkg/model"

type LpType int32
const(
	//0 代表 LP（OPS-USDT）; 1 代表（OPS-FLUX）
	OpsUsdt LpType = 0
	OpsFlux LpType = 1
)

// GetOpsUsdtApyReq for request
type GetOpsUsdtApyReq struct {
}

// GetOpsUsdtApyResponse for response
type GetOpsUsdtApyResponse struct{
	model.JSONResult
	Data ApyValue `json:"data"`
}

// GetOpsFluxApyReq for request
type GetOpsFluxApyReq struct {
}

// GetOpsFluxApyResponse for response
type GetOpsFluxApyResponse struct{
	model.JSONResult
	Data ApyValue `json:"data"`
}

type ApyValue struct{
	Apy string `json:"apy"`
}

// GetOpsPriceUsdtReq for request
type GetOpsPriceUsdtReq struct {
}

// GetOpsPriceUsdtResponse for response
type GetOpsPriceUsdtResponse struct{
	model.JSONResult
	Data OpsPrice `json:"data"`
}
type OpsPrice struct{
	Price string `json:"price"`
}

// GetMiningPoolWorthUsdtReq for request
type GetMiningPoolWorthUsdtReq struct {
	PoolType LpType `json:"pool_type"`
}

// GetMiningPoolWorthUsdtResponse for response
type GetMiningPoolWorthUsdtResponse struct{
	model.JSONResult
	Data GetMiningPoolWorthUsdtData `json:"data"`
}
type GetMiningPoolWorthUsdtData struct{
	Worth string  `json:"worth"`
	PoolType LpType `json:"pool_type"`
	PoolTypeStr string `json:"pool_type_str"`
}


// GetOpsPriceFluxReq for request
type GetOpsPriceFluxReq struct {
}

// GetOpsPriceFluxResponse for response
type GetOpsPriceFluxResponse struct{
	model.JSONResult
	Data OpsPrice `json:"data"`
}

// GetUserLPInfo for request
type GetUserLPInfo struct {
	PoolType LpType `json:"pool_type"`
	Address string `json:"address" binding:"required"`
}

// GetUserLPInfoResponse for response
type GetUserLPInfoResponse struct{
	model.JSONResult
	Data GetUserLPInfoData `json:"data"`
}
type GetUserLPInfoData struct{
	Balance string `json:"balance"`
	Amount string `json:"amount"`
	OpsEarned string `json:"ops_earned"`

	PoolType LpType `json:"pool_type"`
	PoolTypeStr string `json:"pool_type_str"`
}


