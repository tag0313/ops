package contract

import "ops/pkg/model"

// NFT1155BaseMetaURIReq set-base-meta-uri
type NFT1155BaseMetaURIReq struct {
	Prefix string `json:"prefix" binding:"required"`
}
type NFT1155BaseMetaURIResponse struct {
	model.JSONResult
}

// NFT1155CreateReq Create
type NFT1155CreateReq struct {
	InitOwner  string `json:"init_owner" binding:"required"`
	InitSupply string `json:"init_supply" binding:"required"`
	Uri        string `json:"uri"`
	Data       string `json:"data"`
}
type NFT1155CreateResponse struct {
	model.JSONResult
	Data NFT1155CreateData `json:"data"`
}
type NFT1155CreateData struct {
}

// NFT1155CreateBatchReq create-batch
type NFT1155CreateBatchReq struct {
	InitOwner  string   `json:"init_owner" binding:"required"`
	Quantities []string `json:"quantities" binding:"required"`
	Uris       []string `json:"uris"`
	Data       string   `json:"data"`
}
type NFT1155CreateBatchResponse struct {
	model.JSONResult
	Data NFT1155CreateBatchData `json:"data"`
}
type NFT1155CreateBatchData struct {
}

// NFT1155TransferGovernorShipReq transfer-governor-ship
type NFT1155TransferGovernorShipReq struct {
	NewGovernorAddress string `json:"new_governor_address" binding:"required"`
}
type NFT1155TransferGovernorShipResponse struct {
	model.JSONResult
	Data NFT1155TransferGovernorShipData `json:"data"`
}
type NFT1155TransferGovernorShipData struct {
}

// NFT1155MintReq mint
type NFT1155MintReq struct {
	NFT1155IDType
	AddressTo string `json:"address_to" binding:"required"`
	Quantity  string `json:"quantity" binding:"required"`
	Data      string `json:"data" binding:"required"` //????? is really required?
}
type NFT1155MintResponse struct {
	model.JSONResult
	Data NFT1155MintData `json:"data"`
}
type NFT1155MintData struct {
}

// NFT1155MintBatchReq MintBatch
type NFT1155MintBatchReq struct {
	AddressTo  string   `json:"address_to" binding:"required"`
	Ids        []string `json:"ids" binding:"required"`
	Quantities []string `json:"quantities" binding:"required"`
	Data       string   `json:"data" binding:"required"` //????? is really required?
}
type NFT1155MintBatchResponse struct {
	model.JSONResult
	Data NFT1155MintData `json:"data"`
}
type NFT1155MintBatchData struct {
}

// NFT1155SetCreatorReq set-creator
type NFT1155SetCreatorReq struct {
	AddressTo string   `json:"address_to" binding:"required"`
	Ids       []string `json:"ids" binding:"required"`
}
type NFT1155SetCreatorResponse struct {
	model.JSONResult
	Data NFT1155SetCreatorData `json:"data"`
}
type NFT1155SetCreatorData struct {
}

// NFT1155SetIDUri set-id-uri and set-id-uri batch
type NFT1155SetIDUri struct {
	NFT1155IDType
	Uri string `json:"uri" binding:"required"`
}
type NFT1155SetIDUriResponse struct {
	model.JSONResult
	Data NFT1155SetIDUriData `json:"data"`
}
type NFT1155SetIDUriData struct {
}

type NFT1155SetIDUriBatchResponse struct {
	model.JSONResult
	Data []NFT1155SetIDUriData `json:"data"`
}
