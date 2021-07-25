package handler

import (
	"math/big"
	"testing"
)

func TestCalculateApy(t *testing.T){
	/*Input:
		opsDailyNew=73008.000000000000000000,
		dailyPer=0.600000000000000000,
		totalSupply=411013153864639002599,
		r0=750099961160080630174,
		r1=229247755314043901275,
		lpAmount=216999000000000000000,
		opsPrice=0.305622939853905136,
		r0Price=0.305622939853905136,
		r1Price=1.000000000000000000,
		decimals=18

	  Output:
		20186.611094275186743930
	 */
	opsDailyNew := new(big.Rat).SetInt64(73008)
	dailyPer := big.NewRat(6, 10)
	totalSupply, _:=new(big.Int).SetString("411013153864639002599", 10)
	r0, _:=new(big.Int).SetString("750099961160080630174", 10)
	r1, _ := new(big.Int).SetString("229247755314043901275", 10)
	lpAmount, _ :=new(big.Int).SetString("216999000000000000000",10)

	opsPrice, _ :=new(big.Rat).SetString("0.305622939853905136")
	r0Price, _ :=new(big.Rat).SetString("0.305622939853905136")
	r1Price  := big.NewRat(1,1)
	decimals := 18
	output, err := CalculateApy(opsDailyNew, dailyPer,
		totalSupply,  r0, r1, lpAmount,
		opsPrice, r0Price, r1Price,
		decimals)
	if err != nil{
		t.Error()
	}
	result := output.FloatString(decimals)
	t.Log(result)
	if result != "20186.611094275186732648"{
		t.Error("computing error")
	}
}

func newBigInt(str string)*big.Int{
	bi, _:=new(big.Int).SetString(str, 10)
	return bi
}
func newBigRat(str string)*big.Rat{
	bi, _:=new(big.Rat).SetString(str)
	return bi
}
func TestCalculatePoolWorth(t *testing.T) {
	var(
		r0 = newBigInt("750099961160080630174")
		r1 = newBigInt("229247755314043901275")
		r0price = newBigRat("0.305622939853905136")
		r1price = newBigRat("1")
		decimals = 18
		want= newBigRat("458.495510628087802294")
	)
	result := CalculatePoolWorth(r0, r1, r0price, r1price, decimals)
	t.Logf("%s == %s", result.FloatString(decimals), want.FloatString(decimals))
	t.Log(result.Cmp(want))
	if result.FloatString(decimals) != want.FloatString(decimals) {
		t.Error("CalculatePoolWorth error.")
	}
}