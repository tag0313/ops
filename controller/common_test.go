package controller

import (
	"math/big"
	"testing"
)

func TestGasToEth(t *testing.T) {
	f := GasToEth(131776, big.NewInt(1000000000))
	t.Log(f.Text('f', 18))
	if f.Cmp(big.NewFloat(0.000131776)) != 0{
		t.Error("GasToEth computed error")
	}
	for _, n := range []int{1,5,10,20,50,100} {
		t.Logf("%d is %d",n, 131776*n)
		f = GasToEth(uint64(n*131776), big.NewInt(1000000000))
		t.Log(f.Text('f', 18))
	}
}
