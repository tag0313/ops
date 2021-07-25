package utils

import (
	"fmt"
	"github.com/asim/go-micro/v3/logger"
	"net/url"
	"strings"
)

//ParseNFT1155URI return the id inside the nft1155 uri
//example uri: https://ocard.opsnft.net/ocards/oX92N-1624120822941.json
func ParseNFT1155URI(uri string)(id string, err error){
	var parsedURL *url.URL
	parsedURL, err = url.Parse(uri)
	if err != nil {
		return "", err
	}
	path := parsedURL.Path
	seps := strings.Split(path, "/")
	logger.Info(seps)
	for i := range seps {
		var (
			strPart    string
			numberPart uint64
		)
		_, err = fmt.Sscanf(seps[i], "%5s-%d", &strPart, &numberPart)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(seps[i], id, numberPart)
		if numberPart != 0 && strPart != ""{
			return fmt.Sprintf("%s-%d", strPart, numberPart), nil
		}
	}

	return "", nil
}