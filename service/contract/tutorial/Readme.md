# Project document

## How to deploy a smart contract
Smart contract is a kind of programming language, it is executed
by Ethererum virtual machine. I think this is the simplest way to understand
what a smart contract is. But if you want to know it intuitively, I strongly recommend 
you to deploy a simple contract to see how it works. In this article, I will describe the smart contract deploying
step by step. 

## Install the [Ganache](https://github.com/trufflesuite/ganache-cli)
Ganache is a testing block chain running on local machine.
You can use it to test the money transaction, viewing block information,
deploying and executing smart contract. Because of the complexity of 
nodejs development, don't install the Ganache through `npm`, replace it for docker.
```bash
docker pull trufflesuite/ganache-cli:latest
docker run --detach --publish 8545:8545 trufflesuite/ganache-cli:latest
```
Now, you have a block chain running on the localhost:8545.
The initial state of the network has some foundation block, running
`docker logs container_id` will show the private key and balance. And
you can write a client to connect the block chain and check the balance through public address.
```go
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
    //input your the address that output by Ganache  
	account := common.HexToAddress("0x24C524276d6f9a9a9B2d92591806d8059592b008")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance is:", balance) // 25893180161173005034
```
To see more useful code, viewing the `transfer` file.

### Deploy a simple smart contract
We have a simple contract `store.sol`, it implemented two simple method - `set` and `get`
. We will deploy it into the Ganache. In this case, we will 
use `deploy.go` to deploy our contract to Ganache, `contract_load.go`
to load the contract and execute the `set` method, `contract_get.go` to get the value that we stored.

#### building ABI, BIN and go package.
```bash
solc --bin store.sol -o store.bin
solc --abi store.sol -o store.abi
abigen --bin=store.bin/Store.bin  --abi=store.abi/Store.abi --pkg=store --out=store.go
```

### deploying the contract
To deploy a smart contract into block chain, we must have a wallet with
enough balance. But in our testing environment, money is not a problem. We
can copy a private key from `Ganache`. And the generated file `store.go` has
the `DeployStore` method, the remaining work is just call `DeployStore`. 
Create a Ethererum client, paid the gas fee, the contract will be deployed.
```go
    client, err := ethclient.Dial("http://localhost:8545/")
	if err != nil {
		log.Fatal(err)
	}

    auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

    address, tx, instance, err := store.DeployStore(auth, client)
    //....
```

### Calling the set and get method 
The set method is the same as the deploy, you must pay the gas fee to call it. But the get method
doesn't need to pay the gas.
```go
    auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

    //call set
	address, tx, instance, err := store.DeployStore(auth, client)
    
    //call get
    tx, err := instance.Get(nil)
	if err != nil{
		log.Fatal(err)
	}
	log.Println(tx)
```


