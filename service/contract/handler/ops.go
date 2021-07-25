package handler

import (
	"context"
	"crypto/ecdsa"
	"errors"
	microErr "github.com/asim/go-micro/v3/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"net/url"
	"ops/pkg/logger"
	"ops/pkg/utils"
	"ops/proto/contract"
	"ops/service/contract/ercpkg/ops"
)

type opsHandler struct {
	client           *ethclient.Client
	rpcClient        *rpc.Client
	contractInstance *ops.Ops
	abi              abi.ABI
	adminAddress     string
	contractAddress  string

	donateEcdsaPrivate *ecdsa.PrivateKey
}

func NewHandler(opt *Options) (*opsHandler, error) {
	parsedURL, err := url.Parse(opt.BlockChainURLHttp)
	if err != nil {
		return nil, err
	}
	if parsedURL.Scheme == "" {
		return nil, errors.New("invalid url")
	}

	ch := new(opsHandler)
	ch.rpcClient, err = rpc.Dial(opt.BlockChainURLHttp)
	if err != nil {
		return nil, err
	}
	ch.client = ethclient.NewClient(ch.rpcClient)

	ch.contractAddress = opt.OpsContractAddr.String()
	ch.adminAddress = opt.OpsManager.String()
	ch.contractInstance, err = ops.NewOps(opt.OpsContractAddr, ch.client)
	if err != nil {
		return nil, err
	}

	ch.donateEcdsaPrivate, err = crypto.HexToECDSA(opt.DonatePrivateKey)
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (ch *opsHandler) Transfer(ctx context.Context, request *pbContract.TransferRequest, response *pbContract.TransferResponse) error {
	panic("implement me")
}

func (ch *opsHandler) Approve(ctx context.Context, request *pbContract.ApproveRequest, response *pbContract.ApproveResponse) error {
	panic("implement me")
}

func (ch *opsHandler) TransferFrom(ctx context.Context, request *pbContract.TransferFromRequest, response *pbContract.TransferFromResponse) error {
	panic("implement me")
}

func (ch *opsHandler) BalanceOf(ctx context.Context, req *pbContract.BalanceRequest, rep *pbContract.BalanceResponse) error {
	var err error
	defer func() {
		logger.Infof("calling BalanceOf, request=%+v, response=%+v, err=%v", req, rep, err)
	}()
	blockChainAddress := common.HexToAddress(req.Address)
	balance, err := ch.contractInstance.BalanceOf(nil, blockChainAddress)
	if err != nil {
		return err
	}

	rep.Balance = utils.BigInt2Bytes(balance)
	rep.Address = req.Address

	decimals, err := ch.contractInstance.Decimals(nil)
	if err != nil {
		return err
	}
	rep.Decimals = uint32(decimals)
	return nil
}

func (ch *opsHandler) TotalSupply(ctx context.Context, request *pbContract.TotalSupplyRequest, response *pbContract.TotalSupplyResponse) error {
	var err error
	defer func() {
		logger.Infof("calling BalanceOf, request=%+v, response=%+v, err=%v", request, response, err)
	}()
	supply, err := ch.contractInstance.TotalSupply(nil)
	if err != nil {
		return err
	}
	response.TotalSupply = supply.String()
	return nil
}

func (ch *opsHandler) Owner(ctx context.Context, request *pbContract.OwnerRequest, response *pbContract.OwnerResponse) error {
	var err error
	defer func() {
		logger.Infof("calling BalanceOf, request=%+v, response=%+v, err=%v", request, response, err)
	}()
	address, err := ch.contractInstance.Owner(nil)
	if err != nil {
		return err
	}
	response.Address = address.String()
	return nil
}

func (ch *opsHandler) Decimals(ctx context.Context, request *pbContract.DecimalsRequest, response *pbContract.DecimalsResponse) error {
	var err error
	defer func() {
		logger.Infof("calling BalanceOf, request=%+v, response=%+v, err=%v", request, response, err)
	}()
	decimal, err := ch.contractInstance.Decimals(nil)
	if err != nil {
		return err
	}
	response.Decimals = uint32(decimal)
	return nil
}

func (ch *opsHandler) Name(ctx context.Context, request *pbContract.NameRequest, response *pbContract.NameResponse) error {
	var err error
	defer func() {
		logger.Infof("calling BalanceOf, request=%+v, response=%+v, err=%v", request, response, err)
	}()
	name, err := ch.contractInstance.Name(nil)
	if err != nil {
		logger.Error(err)
		return err
	}
	response.Name = name
	logger.Info("Name is", response.Name)
	return nil
}

func (ch *opsHandler) Symbol(ctx context.Context, request *pbContract.SymbolRequest, response *pbContract.SymbolResponse) error {
	var err error
	defer func() {
		logger.Infof("calling BalanceOf, request=%+v, response=%+v, err=%v", request, response, err)
	}()
	symbol, err := ch.contractInstance.Symbol(nil)
	if err != nil {
		return err
	}
	response.Symbol = symbol
	return nil
}

func (h *opsHandler) GetGasFee(ctx context.Context, request *pbContract.GetGasFeeRequest, response *pbContract.GetGasFeeResponse) error {
	var err error
	defer func() {
		logger.Infof("calling GetGasFee, request=%+v, response=%+v, err=%v", request, response, err)
	}()

	if len(request.FromAddress) == 0 && len(request.ToAddress) == 0 {
		return errors.New("fromAddress or toAddress cannot be empty both")
	}
	var (
		from string
		to   string
	)
	if len(request.FromAddress) == 0 {
		from = h.adminAddress
	} else {
		from = string(request.FromAddress)
	}

	if len(request.ToAddress) == 0 {
		to = h.adminAddress
	} else {
		to = string(request.ToAddress)
	}
	amount := utils.Bytes2bigInt(request.Amount)
	toAddress := common.HexToAddress(to)
	input, err := opsABI.Pack("transfer", &toAddress, amount)
	if err != nil {
		return err
	}
	gas, gasPrice, err := getGasFee(ctx, h.client, common.HexToAddress(from),
		toAddress, input)
	if err != nil {
		return err
	}
	logger.Infof("gas=%s  gasPrice=%d", gas.String(), gasPrice)

	response.GasPrice = gas.Bytes()
	response.GasLimit = gasPrice
	return nil
}

func (h *opsHandler) GetTransactionByHash(ctx context.Context,
	request *pbContract.GetTransactionByHashRequest,
	response *pbContract.GetTransactionByHashResponse) error {
	var err error
	defer func() {
		logger.Infof("calling GetTransactionByHash, request=%+v, response=%+v, err=%v",
			request, response, err)
	}()
	tx, isPending, err := h.client.TransactionByHash(ctx,
		common.HexToHash(request.TransactionHash))
	if err != nil {
		return err
	}
	logger.Infof("%+v", tx)
	bytes, _ := tx.MarshalJSON()
	logger.Infof("%s", string(bytes))

	logger.Info(tx.To().String(), " ", globalOptions.HolderContractAddr.String())
	if tx.To().String() == globalOptions.HolderContractAddr.String() { //提现
		var decimals uint8
		decimals, err = opsDecimalsFunc()
		if err != nil {
			return err
		}
		output := make(map[string]interface{})
		m, er := holderABI.MethodById(tx.Data()[:4])
		if er != nil {
			logger.Error(er)
			return er
		}
		if er = m.Inputs.UnpackIntoMap(output, tx.Data()[4:]); er != nil {
			logger.Error(er)
			return er
		}
		value := output["amount"].(*big.Int)
		response.ContractAmount = utils.BigIntToFloatString(value, int(decimals))
		toAddress := output["to"].(common.Address)
		tokenAddress := output["tokenContract"].(common.Address)
		response.ContractTo = toAddress.String()
		response.ContractAddress = tokenAddress.String()
	} else if tx.To().String() == globalOptions.OpsContractAddr.String() { //充值
		output := make(map[string]interface{})
		m, er := opsABI.MethodById(tx.Data()[:4])
		if er != nil {
			logger.Error(er)
			return er
		}
		if er = m.Inputs.UnpackIntoMap(output, tx.Data()[4:]); er != nil {
			logger.Error(er)
			return er
		}

		var decimals uint8
		decimals, err = opsDecimalsFunc()
		if err != nil {
			return err
		}
		value := output["amount"].(*big.Int)
		response.ContractAmount = utils.BigIntToFloatString(value, int(decimals))
		toAddress := output["recipient"].(common.Address)
		response.ContractTo = toAddress.String()
	} else if tx.To().String() == globalOptions.NFT1155ContractAddr.String() {
		logger.Infof("==========================================")
		output := make(map[string]interface{})
		m, er := nft1155ABI.MethodById(tx.Data()[:4])
		if er != nil {
			logger.Error(er)
			return er
		}
		if er = m.Inputs.UnpackIntoMap(output, tx.Data()[4:]); er != nil {
			logger.Error(er)
			return er
		}
		logger.Infof("%+v", output)

	}

	response.Type = uint32(tx.Type())
	response.Nonce = tx.Nonce()
	response.GasPrice = tx.GasPrice().Bytes()
	response.Gas = tx.Gas()
	response.EthValue = tx.Value().Bytes()
	response.To = tx.To().String()
	response.Hash = tx.Hash().String()
	response.IsPending = isPending

	if isPending == false { //if the transaction is not pending, we can get the receipt exactly
		var r *types.Receipt
		r, err = h.client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			logger.Error(err)
			response.Status = 0
		} else {
			response.Status = uint32(r.Status)
			logger.Error(err)
		}
	}
	logger.Info(response)
	return nil
}

func (ch *opsHandler) TestError(ctx context.Context, empty *pbContract.Empty, empty2 *pbContract.Empty) error {
	logger.Infof("calling test error")
	defer ctx.Done()
	return microErr.BadRequest("401", " bad fucking %s", "fucking")
}

func (ch *opsHandler) DonateCoin(ctx context.Context,
	request *pbContract.DonateCoinRequest,
	response *pbContract.DonateCoinResponse) (err error) {
	defer func() {
		logger.Infof("calling DonateCoin, request=%+v, resposne=%+v, err=%v",
			request, response, err,
		)
	}()
	opsUsdtIns, err := ops.NewOps(globalOptions.SwapOpsUsdtContractAddr, ethereumClient)
	if err != nil {
		return err
	}

	opsFluxIns, err := ops.NewOps(globalOptions.SwapOpsFluxContractAddr, ethereumClient)
	if err != nil {
		return err
	}
	toAddr := common.HexToAddress(request.UserWalletAddress)

	//donate eth
	donateAmount := big.NewRat(2, 100) //0.02
	ethTransactionHash, nonce, err := sendEthTransaction(ctx,
		ch.rpcClient,
		ch.donateEcdsaPrivate,
		toAddr,
		donateAmount)
	if err != nil {
		return
	}
	response.Coins = append(response.Coins, &pbContract.DonateCoin{
		CoinType:        "eth",
		CoinAmount:      donateAmount.FloatString(18),
		TransactionHash: ethTransactionHash,
	})

	//donate ops-usdt lp
	//nonce+1
	opts, err := retrieveTransactionOpts(ch.donateEcdsaPrivate, ch.client, big.NewInt(int64(nonce+1)))
	if err != nil {
		return
	}

	decimals, err := getDecimals(globalOptions.OpsContractAddr)
	if err != nil {
		return
	}
	pow := utils.BigIntPow10(int(decimals))
	baseAmount := big.NewRat(5, 1)
	opsUsdtAmount := new(big.Rat).Mul(baseAmount, new(big.Rat).SetInt(pow))
	opts.GasPrice, opts.GasLimit, err = getOpsTransactionFee(ctx,
		ch.client, opts.From, toAddr, opsUsdtAmount.Num())
	if err != nil {
		return err
	}
	opts.GasLimit = opts.GasLimit * 3
	opsUsdtTransaction, err := opsUsdtIns.Transfer(opts, toAddr, opsUsdtAmount.Num())
	if err != nil {
		logger.Error("ops-usdt transfer failed, err is: ", err)
		return
	}
	response.Coins = append(response.Coins, &pbContract.DonateCoin{
		CoinType:        "opsUsdtLP",
		CoinAmount:      baseAmount.FloatString(18),
		TransactionHash: opsUsdtTransaction.Hash().String(),
	})

	//donate ops-flux lp
	baseAmount = big.NewRat(5, 1)
	opsFluxAmount := new(big.Rat).Mul(baseAmount, new(big.Rat).SetInt(pow))

	opts, err = retrieveTransactionOpts(ch.donateEcdsaPrivate,
		ch.client, new(big.Int).Add(opts.Nonce, big.NewInt(1))) //nonce+1
	if err != nil {
		return
	}
	opts.GasPrice, opts.GasLimit, err = getOpsTransactionFee(ctx,
		ch.client, opts.From, toAddr, opsFluxAmount.Num())
	if err != nil {
		return err
	}
	opts.GasLimit = opts.GasLimit * 3
	opsFluxTransaction, err := opsFluxIns.Transfer(opts, toAddr, opsFluxAmount.Num())
	if err != nil {
		logger.Error("ops-flux transfer failed, err is: ", err)
		return
	}
	response.Coins = append(response.Coins, &pbContract.DonateCoin{
		CoinType:        "opsFluxLP",
		CoinAmount:      baseAmount.FloatString(18),
		TransactionHash: opsFluxTransaction.Hash().String(),
	})

	return nil
}
