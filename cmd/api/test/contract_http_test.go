package test


/*
import (
	"bytes"
	"encoding/json"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"ops/web/model/consulreg"
	contractModel "ops/web/model/contract"
	"os"
	"testing"
)

func setupMicro() {
	reg := consul.NewRegistry(
		registry.Addrs("localhost"))
	consulreg.MicroSer = micro.NewService(
		micro.Registry(reg),
	)
	gin.SetMode(gin.TestMode)
	logger.Init(logger.WithOutput(os.Stdout))
}

var (
	_router *gin.Engine = nil
)

func requestTestingServer(path string, data interface{}) *httptest.ResponseRecorder {
	if _router == nil {
		_router = main2.ginRouter(nil)
	}
	const (
		apiVer1Prefix = "/api/v1.0/contract"
	)
	wr := httptest.NewRecorder()
	var req *http.Request
	if data != nil {
		buf, _ := json.Marshal(data)
		req, _ = http.NewRequest("POST", apiVer1Prefix+path, bytes.NewBuffer(buf))
	} else {
		req, _ = http.NewRequest("POST", apiVer1Prefix+path, nil)
	}
	_router.ServeHTTP(wr, req)
	return wr

}

//You must starting a contract backend pb service for this interface
//func TestContractCrc20Reading(t *testing.T) {
//	setupMicro()
//
//	httpResp := requestTestingServer("/info", nil)
//	if httpResp.Code != http.StatusOK {
//		t.Errorf("call info error, http status is %d\n", httpResp.Code)
//	}
//	t.Log(httpResp.Body)
//
//	//test balance
//	m := contractModel.BalanceRequest{
//		Address: "0x66d59cA5721Ce058B706581d983bbD7c5bA366f1",
//	}
//	httpResp = requestTestingServer("/balance", &m)
//	if httpResp.Code != http.StatusOK{
//		t.Error("call balance failed")
//	}
//	t.Log(httpResp.Body)
//
//	//check nil address
//	m.Address = ""
//	httpResp = requestTestingServer("/balance", &m)
//	if httpResp.Code != http.StatusBadRequest{
//		t.Error("balance interface doesn't check the input parameters")
//	}
//}

func TestNFT1155Reading(t *testing.T) {
	setupMicro()
	//test for these methods
	//r5.POST("/nft1155/info", controller.NFT1155Info)
	//r5.POST("/nft1155/balance-of", controller.NFT1155Balance)
	//r5.POST("/nft1155/balance-of-batch", controller.NFT1155BalanceBatch)
	//r5.POST("/nft1155/uri", controller.NFT1155URI)
	//r5.POST("/nft1155/get-next-token-id", controller.NFT1155NextTokenID )
	//r5.POST("/nft1155/is-approved-for-all", controller.NFT1155IsApprovedForAll)

	httpResp := requestTestingServer("/nft1155/info", nil)
	if httpResp.Code != http.StatusOK {
		t.Error("calling info failed")
	}

	balanceOfParams := contractModel.NFT1155BalanceReq{
		OwnerAddress:  "0x66d59ca5721ce058b706581d983bbd7c5ba366f1",
		NFT1155IDType: contractModel.NFT1155IDType{ID: "1"},
	}
	httpResp = requestTestingServer("/nft1155/balance-of", balanceOfParams)
	if httpResp.Code != http.StatusOK {
		t.Error("call balance-of failed")
	}
	var response contractModel.NFT1155Balance
	json.NewDecoder(httpResp.Body).Decode(&response)
	if response.Data.OwnerAddress != balanceOfParams.OwnerAddress {
		t.Error("check owner address failed")
	}

	httpResp = requestTestingServer("/nft1155/balance-of", nil)
	if httpResp.Code != http.StatusBadRequest {
		t.Error("the balance-of input parameters must be checked")
	}

	httpResp = requestTestingServer("/nft1155/balance-of-batch", nil)
	if httpResp.Code != http.StatusBadRequest {
		t.Error("the balance-of-batch input parameters must be checked")
	}
}
 */