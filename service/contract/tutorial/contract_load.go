package tutorial

/*
import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/WhoopsDev/OPS-Server-Dev/contract/store"
)

func main() {
	client, err := ethclient.Dial("http://localhost:8545/")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x6B0ac57fb52a2C561A8f6ccD9f75ec3171520DC3")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}


	privateKey, err := crypto.HexToECDSA("b5c510fff5761282586809aceeb7704a01c4350f1224064bfee3afccda4604cc")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	fmt.Println("contract is loaded")
	tx, err := instance.Set(auth, big.NewInt(23333333))
	if err != nil{
		log.Fatal(err)
	}
	log.Println(tx)
}


 */