package contract

import "ops/pkg/model"

type NFT1155Info struct {
	model.JSONResult
	Data NFT1155InfoData `json:"data"`
}

type NFT1155InfoData struct {
	Name        string `json:"name"`
	TokenSupply string `json:"token_supply"` //uint256
	Owner       string `json:"owner"`
	Symbol      string `json:"symbol"`
}

type NFT1155NextTokenID struct {
	model.JSONResult
	Data NFT1155IDType `json:"data"`
}

type NFT1155IDType struct {
	ID string `json:"id" example:"1234567890" binding:"required"` //base is 10
}

type NFT1155Balance struct {
	model.JSONResult
	Data NFTBalanceType `json:"data"`
}

type NFT1155BalanceBatchResponse struct {
	model.JSONResult
	Data []NFTBalanceType `json: "data"`
}

type NFTBalanceType struct {
	NFT1155IDType
	OwnerAddress string `json:"owner_address"`
	Amount       string `json:"amount" example:"1234567890"`
}

type NFT1155BalanceReq struct {
	OwnerAddress string `json:"owner_address" binding:"required"`
	NFT1155IDType
}

type NFT1155ApprovedReq struct {
	OwnerAddress    string `json:"owner_address" binding:"required"`
	OperatorAddress string `json:"operator_address" binding:"required"`
}

type NFT1155ApprovedResponse struct {
	model.JSONResult
	Data approved `json:"data"`
}

type approved struct {
	IsApproved bool `json:"is_approved"`
}

type NFT1155URIRequest struct {
	NFT1155IDType
}

type NFT1155URIResponse struct {
	model.JSONResult
	Data uri `json:"data"`
}

type uri struct {
	Uri string `json:"uri"`
}

// NFT1155CreateBatchPriceReq for request
type NFT1155CreateBatchPriceReq struct {
	NFT1155CreateBatchReq
}

// NFT1155CreateBatchPriceResponse for response
type NFT1155CreateBatchPriceResponse struct {
	model.JSONResult
	Data NFT1155CreateBatchPriceData `json:"data"`
}
type NFT1155CreateBatchPriceData struct {
	Money string `json:"money"`
}
