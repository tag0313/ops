package controller

import (
	"context"
	"errors"
	"ops/pkg/logger"
	"ops/pkg/model/consulreg"
	"ops/pkg/utils"
	pbContract "ops/proto/contract"
	"ops/proto/nft1155"
	"ops/proto/property"
	"ops/proto/swap"
)

var (
	RpcPropertyService pbProperty.OperatePropertyService
	RpcNft1155Service  pbNft1155.NFT1155Service
	RpcSwapService     pbSwap.SwapService
	RpcOpsService      pbContract.ContractService
	RpcLpService       pbContract.LpService
)

func InitRpcCaller() {
	RpcPropertyService = pbProperty.NewOperatePropertyService("property", consulreg.MicroSer.Client())
	RpcNft1155Service = pbNft1155.NewNFT1155Service("contract", consulreg.MicroSer.Client())
	RpcSwapService = pbSwap.NewSwapService("contract", consulreg.MicroSer.Client())
	RpcOpsService = pbContract.NewContractService("contract", consulreg.MicroSer.Client())
	RpcLpService = pbContract.NewLpService("contract", consulreg.MicroSer.Client())
}

func GetOpsPoint(uid string, currentTotal float64) (string, error) {
	getOpspoint, err := RpcPropertyService.CheckOpsPoint(context.TODO(), &pbProperty.OpsPoint{Uid: uid})
	if err != nil { //4005
		return utils.RECODE_MICROERR, err
	} else if getOpspoint.Uid == utils.RECODE_DATAERR { //4004
		return utils.RECODE_DATAERR, err
	} else if getOpspoint.OpsPoint-currentTotal < 0 { //4301
		return utils.RECODE_INSUFFICIENT_FUND, err
	}
	return utils.RECODE_OK, nil
}

func OperateOpspoint(uid string, currentTotal float64) (string, error) {
	minusOpspoint, err := RpcPropertyService.OperateOpsPoint(context.TODO(), &pbProperty.OpsPoint{Uid: uid, OpsPoint: currentTotal})
	if err != nil { //4005
		return utils.RECODE_MICROERR, nil
	} else if minusOpspoint.Code == utils.RECODE_STOREDATA_FAILED { //4006
		return utils.RECODE_STOREDATA_FAILED, nil
	}
	return utils.RECODE_OK, nil
}

func CheckOcardAmountOps(buyerId, groupId string, quantities int64) error {
	ops, err := RpcPropertyService.CheckOCardAmountOps(context.TODO(), &pbProperty.OCardsOnOps{GroupId: groupId, BuyerUid: buyerId})
	if err != nil {
		return err
	} else if ops.Amount-quantities < 0 {
		logger.Infof("opsAmount is=%+v, parameterAmount is=%+v", ops.Amount, quantities)
		return errors.New("卡数量不足")
	}
	return nil
}

func OpearteOcardAmountOps(buyerId, groupId string, amount int64) error {
	remote, err := RpcPropertyService.OperateOCardAmountOps(context.TODO(), &pbProperty.OCardsOnOps{BuyerUid: buyerId, GroupId: groupId, Amount: amount})
	if err != nil {
		logger.Error(err)
		return err
	} else if remote.Code == utils.RECODE_STOREDATA_FAILED {
		return errors.New("卡操作失败")
	}
	logger.Infof("opsAmount is=%+v, parameterAmount is=%+v", amount, remote.Code)
	return nil
}
