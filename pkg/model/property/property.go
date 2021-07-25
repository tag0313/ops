package property

import (
	"encoding/json"
	"ops/pkg/model"
	"ops/pkg/utils"
	"ops/proto/property"
)

type OpsPointData struct {
	Uid      string  `json:"uid"`
	OpsPoint float64 `json:"ops_point"`
}

type OpsPointResp struct {
	model.JSONResult
	Data *OpsPointData `json:"data"`
}

func (o *OpsPointResp) NewSuccess(data *pbProperty.OpsPoint) *OpsPointResp {
	o.Code = utils.RECODE_OK
	o.Message = utils.RecodeTest(utils.RECODE_OK)
	o.Success = true
	marshal, _ := json.Marshal(&data)
	realData := &OpsPointData{}
	_ = json.Unmarshal(marshal, &realData)
	o.Data = realData
	return o
}

type OCardResp struct {
	GroupId         string `json:"group_id"`
	CardType        string `json:"card_type"`
	MintDate        string `json:"mint_date"`
	Uid             string `json:"uid"`
	Amount          int64  `json:"amount"`
	TransactionHash string `json:"transaction_hash"`
	JsonId          string `json:"json_id"`
	Sold            int64  `json:"sold"`
}

type OCard struct {
	model.JSONResult
	Data []*OCardResp `json:"data"`
}

func (c *OCard) NewSuccess(data []*pbProperty.OCardOnMongo) *OCard {
	c.Code = utils.RECODE_OK
	c.Message = utils.RecodeTest(utils.RECODE_OK)
	c.Success = true
	marshal, _ := json.Marshal(&data)
	var realData []*OCardResp
	_ = json.Unmarshal(marshal, &realData)
	c.Data = realData
	return c
}

type MintCardReq struct {
	Info   []*Info `json:"info"`
	Gasfee string  `json:"gasfee"`
}

type Info struct {
	Ocard  string `json:"ocard"`
	Amount string `json:"amount"`
}

// MintCardResponse for response
type MintCardResponse struct {
	model.JSONResult
	Data *pbProperty.MintCard `json:"data,omitempty"`
}

type OCardsOnOpsResp struct {
	BuyerUid     string  `json:"buyer_uid"`
	GroupId      string  `json:"group_id"`
	CardType     string  `json:"card_type"`
	MintDate     string  `json:"mint_date"`
	SellerUid    string  `json:"seller_uid"`
	Amount       int64   `json:"amount"`
	PurchaseTime string  `json:"purchase_time"`
	UnitPrice    float64 `json:"unit_price"`
}

type OCardOnOps struct {
	model.JSONResult
	Data []*OCardsOnOpsResp `json:"data"`
}

func (c *OCardOnOps) NewSuccess(data []*pbProperty.OCardsOnOps) *OCardOnOps {
	c.Code = utils.RECODE_OK
	c.Message = utils.RecodeTest(utils.RECODE_OK)
	c.Success = true
	marshal, _ := json.Marshal(&data)
	var realData []*OCardsOnOpsResp
	_ = json.Unmarshal(marshal, &realData)
	c.Data = realData
	return c
}

// TransferCardReq for request
type TransferCardReq struct {
	//NFTContractAddr string   `json:"nft_contract_addr"`
	ToAddress  string   `json:"to_address"`
	Ids        []string `json:"ids"`
	Quantities []int64  `json:"quantities"`
	Gasfee     string   `json:"gasfee"`
}

// TransferCardResponse for response
type TransferCardResponse struct {
	model.JSONResult
	Data TransferCardData `json:"data"`
}
type TransferCardData struct {
	TransactionHash string `json:"transaction_hash"`
}

// WithDrawOpspointReq for request
type WithdrawOpspointReq struct {
	Opspoint float64 `json:"opspoint" binding:"required"`
	Gasfee   float64 `json:"gasfee" binding:"required"`
}

// WithdrawOpsResponse for response
type WithdrawOpsResponse struct {
	model.JSONResult
}

// QueryTransactionReq for request
type QueryTransactionReq struct {
	TransactionHash string `json:"transaction_hash" binding:"required"`
}

// QueryTransactionResponse for response
type QueryTransactionResponse struct {
	model.JSONResult
	Data TransactionCardResponseData `json:"data"`
}

type IDAmount struct {
	ID     uint64 `json:"id"`
	Amount uint64 `json:"amount"`
}

type TransactionCardResponseData struct {
	ContractFrom    string     `json:"contract_from"`
	ContractAddress string     `json:"contract_address"`
	ContractTo      string     `json:"contract_to"`
	Data            []byte     `json:"data"`
	IdAndAmount     []IDAmount `json:"id_and_amount"`
	IsPending       bool       `json:"is_pending"`
}

// ChargeToServerReq for request
type ChargeToServerReq struct {
	FromAddress string   `json:"from_address"`
	Ids         []string `json:"ids"`
	Quantities  []string `json:"quantities"`
	GasFee      string   `json:"gas_fee"`
}

// ChargeToServerResponse for response
type ChargeToServerResponse struct {
	model.JSONResult
	Data ChargeToServerData `json:"data"`
}
type ChargeToServerData struct {
}

// QueryTransactionFeeReq for request
type QueryTransactionFeeReq struct {
	NFTContractAddr string   `json:"nft_contract_addr"`
	ToAddress       string   `json:"to_address"`
	Ids             []string `json:"ids"`
	Quantities      []string `json:"quantities"`
}

// QueryTransactionFeeResponse for response
type QueryTransactionFeeResponse struct {
	model.JSONResult
	Data QueryTransactionFeeData `json:"data"`
}
type QueryTransactionFeeData struct {
	Money string `json:"money"`
}
