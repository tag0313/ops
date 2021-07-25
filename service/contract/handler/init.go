package handler

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	nft1155BC "ops/service/contract/ercpkg/nft1155"
	holderBC "ops/service/contract/ercpkg/nftholder"
	opsBC "ops/service/contract/ercpkg/ops"
	opslpBC "ops/service/contract/ercpkg/opslp"
	contractSwap "ops/service/contract/ercpkg/swap"
	"reflect"
	"strings"
)

var (
	opsABI, _ = abi.JSON(strings.NewReader(opsBC.OpsABI))
	nft1155ABI, _ = abi.JSON(strings.NewReader(nft1155BC.Nft1155ABI))
	holderABI, _ = abi.JSON(strings.NewReader(holderBC.NftholderABI))

	nft1155Ins      *nft1155BC.Nft1155
	holderIns       *holderBC.Nftholder
	opsIns          *opsBC.Ops
	lpOpsUsdtIns    *opslpBC.Opslp
	lpOpsFluxIns    *opslpBC.Opslp
	lpEthUsdtIns    *opslpBC.Opslp
	swapOps2usdtIns *contractSwap.Swap
	swapEth2usdtIns *contractSwap.Swap
	swapOps2FluxIns *contractSwap.Swap

	ethereumClient *ethclient.Client
	globalOptions *Options
)

func InitContract(opts *Options) error{
	var err error
	ethereumClient, err = ethclient.Dial(opts.BlockChainURLWss)
	if err != nil{
		return err
	}

	if opsIns, err = opsBC.NewOps(opts.OpsContractAddr, ethereumClient); err != nil{
		return err
	}

	if nft1155Ins, err = nft1155BC.NewNft1155(opts.NFT1155ContractAddr, ethereumClient); err != nil{
		return err
	}

	if holderIns, err = holderBC.NewNftholder(opts.HolderContractAddr, ethereumClient); err != nil{
		return err
	}

	if lpOpsUsdtIns, err = opslpBC.NewOpslp(opts.OpsLpUsdtContractAddr, ethereumClient); err != nil{
		return err
	}

	swapEth2usdtIns, err = contractSwap.NewSwap(opts.SwapEthUsdtContractAddr, ethereumClient)
	if err !=nil{
		return err
	}

	swapOps2usdtIns, err = contractSwap.NewSwap(opts.SwapOpsUsdtContractAddr, ethereumClient)
	if err !=nil{
		return err
	}
	swapOps2FluxIns, err = contractSwap.NewSwap(opts.SwapOpsFluxContractAddr, ethereumClient)
	if err != nil{
		return err
	}

	lpOpsFluxIns, err = opslpBC.NewOpslp(opts.OpsLpFluxContractAddr, ethereumClient)
	if err != nil{
		return err
	}

	lpEthUsdtIns, err = opslpBC.NewOpslp(opts.EthLpUsdtContractAddr, ethereumClient)
	if err != nil{
		return err
	}

	globalOptions = opts
	return nil
}

type Options struct {
	BlockChainURLHttp string
	BlockChainURLWss string

	LogPath string
	ListenAddr string
	MicroName string
	Version string

	NFT1155Manager common.Address
	NFT1155PrivateKey string
	DonatePrivateKey string

	OpsManager              common.Address
	OpsContractAddr         common.Address
	NFT1155ContractAddr     common.Address
	HolderContractAddr      common.Address
	SwapOpsUsdtContractAddr common.Address
	SwapOpsFluxContractAddr common.Address
	SwapEthUsdtContractAddr common.Address

	OpsLpUsdtContractAddr   common.Address
	OpsLpFluxContractAddr   common.Address
	EthLpUsdtContractAddr   common.Address
}

func NewOptions(viperIns *viper.Viper) *Options {
	opt := new(Options)
	checkError := func(e error){
		if e != nil{
			panic(e)
		}
	}

	checkError(viperIns.UnmarshalKey("log_path", &opt.LogPath))
	checkError(viperIns.UnmarshalKey("chainServer.http", &opt.BlockChainURLHttp))
	checkError(viperIns.UnmarshalKey("chainServer.wss", &opt.BlockChainURLWss))

	checkError(viperIns.UnmarshalKey("micro.addr", &opt.ListenAddr))
	checkError(viperIns.UnmarshalKey("micro.name", &opt.MicroName))
	checkError(viperIns.UnmarshalKey("micro.version", &opt.Version))
	checkError(viperIns.UnmarshalKey("key.nft1155_private", &opt.NFT1155PrivateKey))
	checkError(viperIns.UnmarshalKey("key.donate_private", &opt.DonatePrivateKey))

	addrDecodeHook := viper.DecodeHook(AddressStringHookFunc())
	checkError(viperIns.UnmarshalKey("key.nft1155_manager", &opt.NFT1155Manager, addrDecodeHook))
	checkError(viperIns.UnmarshalKey("key.erc20_manager", &opt.OpsManager, addrDecodeHook))

	checkError(viperIns.UnmarshalKey("contract_addr.ops", &opt.OpsContractAddr, addrDecodeHook))
	checkError(viperIns.UnmarshalKey("contract_addr.nft1155", &opt.NFT1155ContractAddr, addrDecodeHook))
	checkError(viperIns.UnmarshalKey("contract_addr.holder", &opt.HolderContractAddr, addrDecodeHook))
	checkError(viperIns.UnmarshalKey("contract_addr.swap_ops_usdt", &opt.SwapOpsUsdtContractAddr, addrDecodeHook))
	checkError(viperIns.UnmarshalKey("contract_addr.swap_eth_usdt", &opt.SwapEthUsdtContractAddr, addrDecodeHook))
	checkError(viperIns.UnmarshalKey("contract_addr.swap_ops_flux", &opt.SwapOpsFluxContractAddr, addrDecodeHook))
	checkError(viperIns.UnmarshalKey("contract_addr.ops_lp_usdt", &opt.OpsLpUsdtContractAddr, addrDecodeHook))
	checkError(viperIns.UnmarshalKey("contract_addr.eth_lp_usdt", &opt.EthLpUsdtContractAddr, addrDecodeHook))
	checkError(viperIns.UnmarshalKey("contract_addr.ops_lp_flux", &opt.OpsLpFluxContractAddr, addrDecodeHook))
	return opt
}

func AddressStringHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		// Check that the data is string
		if f.Kind() != reflect.String {
			return data, nil
		}

		// Check that the target type is our custom type
		if t != reflect.TypeOf(common.Address{}) {
			return data, nil
		}

		strAddr := data.(string)
		if strings.TrimSpace(strAddr) == ""{
			return nil, fmt.Errorf("filed %v is not a address", f)
		}
		if !common.IsHexAddress(strAddr){
			return nil, fmt.Errorf("filed %v is not a hex string", f)
		}
		// Return the parsed value
		addr := common.HexToAddress(strings.TrimSpace(strAddr))
		return addr, nil
	}
}

