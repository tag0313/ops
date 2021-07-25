package handler

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"ops/pkg/logger"
	"ops/pkg/utils"
	"ops/proto/property"
	"reflect"
)

var (
	isSubscribed   = false
	logsChan       chan types.Log
	eventsMap      map[string]*eventInfo
	propertyClient pbProperty.OperatePropertyService = nil
)

type handlerFunc func(eventInfo *eventInfo, log types.Log) error

func initEvents() {
	events := []*eventInfo{
		&eventInfo{
			MethodName: "TransferSingle",
			Abi:        &nft1155ABI,
			handler:    singleCreateEvent,
		},
		&eventInfo{
			MethodName: "TransferBatch",
			Abi:        &nft1155ABI,
			handler:    batchCreateEvent,
		},
		&eventInfo{
			MethodName: "Transfer",
			Abi:        &opsABI,
			handler:    transferEvent,
		},
	}
	eventsMap = make(map[string]*eventInfo)
	for _, e := range events {
		method, ok := e.Abi.Events[e.MethodName]
		if !ok {
			panic(fmt.Errorf("cannot find the method: %s", e.MethodName))
		}
		e.topicHash = crypto.Keccak256Hash([]byte(method.Sig)).Hex()
		fmt.Printf("%s-%s", e.MethodName, e.topicHash)
		eventsMap[e.topicHash] = e
	}
}

type eventInfo struct {
	MethodName string
	topicHash  string
	Abi        *abi.ABI
	handler    handlerFunc
}

func propertyInstance() pbProperty.OperatePropertyService {
	if propertyClient == nil {
		propertyClient = pbProperty.NewOperatePropertyService("property", pbClientG)
	}
	return propertyClient
}

func SubscribeChainEvent() (err error) {
	if isSubscribed {
		return
	}
	initEvents()

	isSubscribed = true
	client, err := rpc.Dial(globalOptions.BlockChainURLWss)
	if err != nil {
		return err
	}

	logsChan = make(chan types.Log, 1024)

	go resultQueue()
	go subscribeLoop(client)
	return nil
}

func subscribeLoop(client *rpc.Client) {
	ethClientIns := ethclient.NewClient(client)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			globalOptions.OpsContractAddr,
			globalOptions.NFT1155ContractAddr,
		},
	}
	chainLogs := make(chan types.Log)
	sub, err := ethClientIns.SubscribeFilterLogs(context.Background(), query, chainLogs)
	if err != nil {
		logger.Fatal(err)
		return
	}

	for {
		select {
		case er := <-sub.Err():
			logger.Fatal(er)
		case vLog := <-chainLogs:
			eventTopic := vLog.Topics[0].Hex()
			if e, ok := eventsMap[eventTopic]; ok {
				logger.Infof("receive event=%s, TopicHash=%s",
					e.MethodName, e.topicHash)
				//parse the result may take too much time
				logsChan <- vLog
			}
		}
	}

	logger.Fatal("this function would never return, something might wrong")
	return
}

func resultQueue() {
	for {
		select {
		case vlog := <-logsChan:
			e := eventsMap[vlog.Topics[0].Hex()]
			logger.Infof("call event handler: %s", e.MethodName)
			if err := e.handler(e, vlog); err != nil {
				logger.Error(err)
			}
		}
	}
	logger.Fatal("this function would never return, something might wrong")
	return
}

func singleCreateEvent(info *eventInfo, vLog types.Log) error {
	operator := common.BytesToAddress(vLog.Topics[1].Bytes())
	from := common.BytesToAddress(vLog.Topics[2].Bytes())
	to := common.BytesToAddress(vLog.Topics[3].Bytes())
	logger.Infof("operator=%s, from=%s, to=%s", operator, from, to)
	output := make(map[string]interface{})
	err := info.Abi.UnpackIntoMap(output, info.MethodName, vLog.Data)
	if err != nil {
		return err
	}
	id := output["id"].(*big.Int)
	txHash := vLog.TxHash.String()

	logger.Info("receive id: ", id.String())
	pbRequest := new(pbProperty.OCard)
	pbRequest.GroupId = id.String()
	pbRequest.TransactionHash = txHash
	_, err = propertyInstance().StoreMintOCardInfo(context.TODO(), pbRequest)
	if err != nil {
		return err
	}
	return nil
}

func batchCreateEvent(info *eventInfo, vLog types.Log) error {
	operator := common.BytesToAddress(vLog.Topics[1].Bytes())
	from := common.BytesToAddress(vLog.Topics[2].Bytes())
	to := common.BytesToAddress(vLog.Topics[3].Bytes())
	logger.Infof("receive holder event operator=%s, from=%s, to=%s", operator, from, to)
	output := make(map[string]interface{})
	err := info.Abi.UnpackIntoMap(output, info.MethodName, vLog.Data)
	if err != nil {
		return err
	}
	var ids = output["ids"]
	var values = output["values"]
	txHash := vLog.TxHash.String()
	s := reflect.ValueOf(ids)
	var bigIds [][]byte
	for i := 0; i < s.Len(); i++ {
		id := s.Index(i).Interface().(*big.Int)
		bigIds = append(bigIds, id.Bytes())
		logger.Infof("new ids[%d] = %s", i, id.String())
	}
	vs := reflect.ValueOf(values)
	var supplies [][]byte
	for i := 0; i < vs.Len(); i++ {
		v := vs.Index(i).Interface().(*big.Int)
		supplies = append(supplies, v.Bytes())
		logger.Infof("new supplies[%d] = %s", i, v.String())
	}
	pbRequest := &pbProperty.MintOCardsSuccessInfo{
		GroupIds:        bigIds,
		Amounts:         supplies,
		TransactionHash: txHash,
		ToAddress: to.String(),
		FromAddress: from.String(),
		OperatorAddress: operator.String(),
	}
	_, err = propertyInstance().StoreMintOCardsSuccessInfo(context.TODO(), pbRequest)
	return err
}

func transferEvent(info *eventInfo, vLog types.Log) error {
	output := make(map[string]interface{})
	err := info.Abi.UnpackIntoMap(output, info.MethodName, vLog.Data)
	if err != nil {
		return err
	}
	fromAddress := common.BytesToAddress(vLog.Topics[1].Bytes())
	toAddress := common.BytesToAddress(vLog.Topics[2].Bytes())
	value := output["value"].(*big.Int)

	logger.Infof("%+v fromAddress=%s toAddress=%s, txHash=%s",
		output["value"], fromAddress.String(), toAddress.String(), vLog.TxHash.String())

	decimals, err := opsDecimalsFunc()
	if err != nil {
		logger.Error(err)
		return err
	}

	pbRequest := &pbProperty.TransferERC20Request{
		FromAddress:     fromAddress.String(),
		ToAddress:       toAddress.String(),
		Amount:          utils.BigIntToFloatString(value, int(decimals)),
		TransactionHash: vLog.TxHash.String(),
	}
	_, err = propertyInstance().StoreTransferERC20History(context.TODO(), pbRequest)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
