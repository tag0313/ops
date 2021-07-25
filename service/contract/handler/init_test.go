package handler

import (
	"fmt"
	"github.com/spf13/viper"
	"ops/pkg/logger"
	"os"
	"strings"
	"testing"
)

var(
configStr =`
log_path: ./log
micro:
  addr: localhost
  name: contract
  version: 1.0

key:
  nft1155_manager: 0x66d59cA5721Ce058B706581d983bbD7c5bA366f1
  nft1155_private: 83f14453333d28879b008258d9a24bf0409fc69699275198845edb2ae0fc00d6
  erc20_manager:  0x66d59cA5721Ce058B706581d983bbD7c5bA366f1
  donate_private: fc015d610cbfd6b087f36cffca6427a0425a2e39fcdc546160b655434a01c854

chainServer:
  http: https://rinkeby.infura.io/v3/aacefa715bbb4a23b6ad98cf61cec2bc
  wss: wss://rinkeby.infura.io/ws/v3/aacefa715bbb4a23b6ad98cf61cec2bc

contract_addr:
  ops: 0xb247BefC358BeDf8E0f305826350cC40184879c5
  nft1155: 0x2b7D03324D2c8E89E70C058088Bc08f14AEF8Da3
  holder: 0xDfC05ec7a9C15643411A2584438E9D085d4c54eE
  swap_ops_usdt: 0xe8788cde4b81587386230b36093b96cbdf7f6a0a
  swap_ops_flux: 0x11f3908f579e3aee9abeb76683ae1c90a91594df
  swap_eth_usdt: 0x5308a481B2b65F6086083D2268acb73AADC757E0
  ops_lp_usdt: 0x87f991e87eCcF53DD65108aACd2aCaD6A30C0fC2
  eth_lp_usdt: 0x5308a481B2b65F6086083D2268acb73AADC757E0
  ops_lp_flux: 0x11f3908f579e3aee9abeb76683ae1c90a91594df
`
	testOptions *Options
)

func TestMain(m *testing.M) {
	fmt.Println("running test main")
	logger.InitDefault(&logger.Options{
		AbsLogDir: "",
		Debug: true,
	})
	defer logger.Sync()
	logger.Info("test main running before")
	v := viper.New()
	v.SetConfigName("conf")
	v.SetConfigType("yaml")
	if err := v.ReadConfig(strings.NewReader(configStr)); err != nil{
		logger.Fatal(err)
	}
	testOptions = NewOptions(v)
	if err := InitContract(testOptions); err != nil{
		logger.Fatal(err)
	}
	code := m.Run()

	logger.Info("test main running after")

	os.Exit(code)
}
