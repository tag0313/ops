package controller

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	"github.com/gin-gonic/gin"
	"math"
	"math/big"
	"net/http"
	baseModel "ops/pkg/model"
	"reflect"
)

//copy from https://github.com/ethereum/go-ethereum/blob/d8ff53dfb8a516f47db37dbc7fd7ad18a1e8a125/params/denomination.go
const (
	Wei   = 1
	GWei  = 1e9
	Ether = 1e18
)

func ByteToFloat64(bytes []byte) float64 {
	if bytes != nil {
		bits := binary.LittleEndian.Uint64(bytes)
		return math.Float64frombits(bits)
	}
	return 0
}

func intBytesToString(bytes []byte) string {
	bi := new(big.Int)
	bi.SetBytes(bytes)
	return bi.String()
}

func int64ToByte(int int64) []byte {
	bi := new(big.Int)
	bi.SetInt64(int)
	return bi.Bytes()
}

func stringDecimalToBytes(decimal string) []byte {
	bi := new(big.Int)
	bi.SetString(decimal, 10)
	return bi.Bytes()
}

func stringDecimalToBigFloat(decimal string) (*big.Float, error) {
	f := new(big.Float)
	a, b := f.SetString(decimal)
	if !b {
		return nil, errors.New("invalid float string")
	}
	return a, nil
}

func intBytesToInt(bytes []byte) int64 {
	bi := new(big.Int)
	bi.SetBytes(bytes)
	return bi.Int64()
}
func intBytesToUint64(bytes []byte) uint64 {
	bi := new(big.Int)
	bi.SetBytes(bytes)
	return bi.Uint64()
}

func checkFloatGreatRate(client, server, rate *big.Float) error {
	x := new(big.Float)
	xsub := x.Sub(client, server)
	xabs := x.Abs(xsub)

	rateMul := rate.Mul(client, rate)

	if xabs.Cmp(rateMul) == 1 {
		return fmt.Errorf("abs(c(%f)-s(%f)) = %f > %f,  rateNumber=%f", client, server, xabs, rateMul, rate)
	}
	return nil
}

func GasToEth(gas uint64, gasPrice *big.Int) *big.Float {
	bg := new(big.Int).SetUint64(gas)
	total := new(big.Int).Mul(bg, gasPrice)
	//convert wei to eth
	wei, _ := new(big.Float).SetString(total.String())
	return new(big.Float).Quo(wei, big.NewFloat(Ether))
}

func abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

type jsonResult interface {
	GetMessage() string
	NewError(recode string) *baseModel.JSONResult
	NewSuccess() *baseModel.JSONResult
}

func responseHTTP(ctx *gin.Context, statusCode int, data jsonResult) {
	if statusCode != http.StatusOK {
		logger.Error(data.GetMessage())
	}
	logger.Infof("http response code is %d, response Type is: %s, response data is: %+v",
		statusCode, reflect.TypeOf(data).String(), data)
	ctx.JSON(statusCode, &data)
}
