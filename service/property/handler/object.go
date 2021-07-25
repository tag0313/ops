package handler

import (
	"ops/pkg/db"
)

const(
	propertyDB = "property"
)


//MintedOCard  mapping for minted_card
type MintedOCard struct {
	*db.ObjectID
	GroupId         string `json:"group_id,omitempty" bson:"group_id,omitempty"`
	CardType        string `json:"card_type,omitempty" bson:"card_type,omitempty"`
	MintDate        string `json:"mint_date,omitempty" bson:"mint_date,omitempty"`
	Uid             string `json:"uid,omitempty" bson:"uid,omitempty"`
	Amount          uint64 `json:"amount,omitempty" bson:"amount,omitempty"`
	TransactionHash string `json:"transaction_hash,omitempty" bson:"transaction_hash,omitempty"`
	JsonId          string `json:"json_id,omitempty" bson:"json_id,omitempty"`
	Sold            int64  `json:"sold,omitempty" bson:"sold,omitempty"`
}

func (m MintedOCard) DbName() string {
	return propertyDB
}

func (m MintedOCard) CollectionName() string {
	return "minted_ocard"
}


//BuyerInfoOps mapping for buyer_info_ops
type BuyerInfoOps struct {
	*db.ObjectID
	BuyerUid     string  `json:"buyer_uid,omitempty" bson:"buyer_uid,omitempty"`
	GroupId      string  `json:"group_id,omitempty" bson:"group_id,omitempty"`
	CardType     string  `json:"card_type,omitempty" bson:"card_type,omitempty"`
	MintDate     string  `json:"mint_date,omitempty" bson:"mint_date,omitempty"`
	SellerUid    string  `json:"seller_uid,omitempty" bson:"seller_uid,omitempty"`
	Amount       uint64  `json:"amount,omitempty" bson:"amount,omitempty"`
	PurchaseTime string  `json:"purchase_time,omitempty" bson:"purchase_time,omitempty"`
	UnitPrice    float64 `json:"unit_price,omitempty" bson:"unit_price,omitempty"`
}

func (b BuyerInfoOps) DbName() string {
	return propertyDB
}

func (b BuyerInfoOps) CollectionName() string {
	return "buyer_info_ops"
}


//BuyerInfoChain mapping for buyer_info_chain
type BuyerInfoChain struct {
	*db.ObjectID
	GroupId         string `json:"group_id" bson:"group_id"`
	Amount          uint64 `json:"amount" bson:"amount"`
	OwnerID         string `json:"owner_id" bson:"owner_id"`
	CreatorID		string `bson:"creator_id" bson:"creator_id"`
	CardType        string `json:"card_type" bson:"card_type"`
}

func (b BuyerInfoChain) DbName() string {
	return propertyDB
}

func (b BuyerInfoChain) CollectionName() string {
	return "buyer_info_chain"
}

type WithdrawOcardHistory struct {
	*db.ObjectID
	GroupID string `json:"group_id"`
	Amount  uint64 `json:"amount"`
	TransactionHash string `json:"transaction_hash"`
	FromAddress     string `json:"from_address"`
	ToAddress       string `json:"to_address"`
	OperatorAddress string `json:"operator_address"`
	BuyerID         string `json:"buyer_id"`
	TransactionTime int64 `json:"transaction_time"`
}

func (w WithdrawOcardHistory) DbName() string {
	return propertyDB
}

func (w WithdrawOcardHistory) CollectionName() string {
	return "withdraw_ocard_history"
}