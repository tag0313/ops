package tutorial

/*
import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"

	"contract/tutorial/store"
)

func main() {
	log.SetFlags(log.Lshortfile|log.Ldate)
	client, err := ethclient.Dial("http://localhost:8545/")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x6B0ac57fb52a2C561A8f6ccD9f75ec3171520DC3")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	tx, err := instance.Get(nil)
	if err != nil{
		log.Fatal(err)
	}
	log.Println(tx)
}


 */