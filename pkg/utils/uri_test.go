package utils

import "testing"

func TestParsedNFT1155URI(t *testing.T) {
	exmapleURI := "https://ocard.opsnft.net/ocards/oX92N-1624120822941.json"
	id, err := ParseNFT1155URI(exmapleURI)
	if err != nil{
		t.Fatal(err)
	}
	t.Log(id, err)
	if id != "oX92N-1624120822941"{
		t.Error("the parser does not work correctly.")
	}
}