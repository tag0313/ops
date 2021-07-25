package model

import (
	"errors"
	"math/big"
)

type AmountDataType struct {
	Amount string `json:"amount" binding:"required" example:"1.0000"`
}

func AmountToFloatString(dataType AmountDataType)(string, error){
	f := new(big.Float)
	a, b := f.SetString(dataType.Amount)
	if !b {
		return "", errors.New("invalid float string")
	}
	return a.String(),nil
}

type TransactionData struct{
	FromAddress string `json:"from_address"`
	ToAddress string `json:"to_address"`
	AmountDataType
	State int `json:"state"`
}
