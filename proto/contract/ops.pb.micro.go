// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/contract/ops.proto

package pbContract

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Contract service

func NewContractEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Contract service

type ContractService interface {
	//reading methods
	BalanceOf(ctx context.Context, in *BalanceRequest, opts ...client.CallOption) (*BalanceResponse, error)
	TotalSupply(ctx context.Context, in *TotalSupplyRequest, opts ...client.CallOption) (*TotalSupplyResponse, error)
	Owner(ctx context.Context, in *OwnerRequest, opts ...client.CallOption) (*OwnerResponse, error)
	Decimals(ctx context.Context, in *DecimalsRequest, opts ...client.CallOption) (*DecimalsResponse, error)
	Name(ctx context.Context, in *NameRequest, opts ...client.CallOption) (*NameResponse, error)
	Symbol(ctx context.Context, in *SymbolRequest, opts ...client.CallOption) (*SymbolResponse, error)
	GetGasFee(ctx context.Context, in *GetGasFeeRequest, opts ...client.CallOption) (*GetGasFeeResponse, error)
	GetTransactionByHash(ctx context.Context, in *GetTransactionByHashRequest, opts ...client.CallOption) (*GetTransactionByHashResponse, error)
	//writing methods
	Transfer(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferResponse, error)
	Approve(ctx context.Context, in *ApproveRequest, opts ...client.CallOption) (*ApproveResponse, error)
	TransferFrom(ctx context.Context, in *TransferFromRequest, opts ...client.CallOption) (*TransferFromResponse, error)
	TestError(ctx context.Context, in *Empty, opts ...client.CallOption) (*Empty, error)
	DonateCoin(ctx context.Context, in *DonateCoinRequest, opts ...client.CallOption) (*DonateCoinResponse, error)
}

type contractService struct {
	c    client.Client
	name string
}

func NewContractService(name string, c client.Client) ContractService {
	return &contractService{
		c:    c,
		name: name,
	}
}

func (c *contractService) BalanceOf(ctx context.Context, in *BalanceRequest, opts ...client.CallOption) (*BalanceResponse, error) {
	req := c.c.NewRequest(c.name, "Contract.BalanceOf", in)
	out := new(BalanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractService) TotalSupply(ctx context.Context, in *TotalSupplyRequest, opts ...client.CallOption) (*TotalSupplyResponse, error) {
	req := c.c.NewRequest(c.name, "Contract.TotalSupply", in)
	out := new(TotalSupplyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractService) Owner(ctx context.Context, in *OwnerRequest, opts ...client.CallOption) (*OwnerResponse, error) {
	req := c.c.NewRequest(c.name, "Contract.Owner", in)
	out := new(OwnerResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractService) Decimals(ctx context.Context, in *DecimalsRequest, opts ...client.CallOption) (*DecimalsResponse, error) {
	req := c.c.NewRequest(c.name, "Contract.Decimals", in)
	out := new(DecimalsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractService) Name(ctx context.Context, in *NameRequest, opts ...client.CallOption) (*NameResponse, error) {
	req := c.c.NewRequest(c.name, "Contract.Name", in)
	out := new(NameResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractService) Symbol(ctx context.Context, in *SymbolRequest, opts ...client.CallOption) (*SymbolResponse, error) {
	req := c.c.NewRequest(c.name, "Contract.Symbol", in)
	out := new(SymbolResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractService) GetGasFee(ctx context.Context, in *GetGasFeeRequest, opts ...client.CallOption) (*GetGasFeeResponse, error) {
	req := c.c.NewRequest(c.name, "Contract.GetGasFee", in)
	out := new(GetGasFeeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractService) GetTransactionByHash(ctx context.Context, in *GetTransactionByHashRequest, opts ...client.CallOption) (*GetTransactionByHashResponse, error) {
	req := c.c.NewRequest(c.name, "Contract.GetTransactionByHash", in)
	out := new(GetTransactionByHashResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractService) Transfer(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferResponse, error) {
	req := c.c.NewRequest(c.name, "Contract.Transfer", in)
	out := new(TransferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractService) Approve(ctx context.Context, in *ApproveRequest, opts ...client.CallOption) (*ApproveResponse, error) {
	req := c.c.NewRequest(c.name, "Contract.Approve", in)
	out := new(ApproveResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractService) TransferFrom(ctx context.Context, in *TransferFromRequest, opts ...client.CallOption) (*TransferFromResponse, error) {
	req := c.c.NewRequest(c.name, "Contract.TransferFrom", in)
	out := new(TransferFromResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractService) TestError(ctx context.Context, in *Empty, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "Contract.TestError", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contractService) DonateCoin(ctx context.Context, in *DonateCoinRequest, opts ...client.CallOption) (*DonateCoinResponse, error) {
	req := c.c.NewRequest(c.name, "Contract.DonateCoin", in)
	out := new(DonateCoinResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Contract service

type ContractHandler interface {
	//reading methods
	BalanceOf(context.Context, *BalanceRequest, *BalanceResponse) error
	TotalSupply(context.Context, *TotalSupplyRequest, *TotalSupplyResponse) error
	Owner(context.Context, *OwnerRequest, *OwnerResponse) error
	Decimals(context.Context, *DecimalsRequest, *DecimalsResponse) error
	Name(context.Context, *NameRequest, *NameResponse) error
	Symbol(context.Context, *SymbolRequest, *SymbolResponse) error
	GetGasFee(context.Context, *GetGasFeeRequest, *GetGasFeeResponse) error
	GetTransactionByHash(context.Context, *GetTransactionByHashRequest, *GetTransactionByHashResponse) error
	//writing methods
	Transfer(context.Context, *TransferRequest, *TransferResponse) error
	Approve(context.Context, *ApproveRequest, *ApproveResponse) error
	TransferFrom(context.Context, *TransferFromRequest, *TransferFromResponse) error
	TestError(context.Context, *Empty, *Empty) error
	DonateCoin(context.Context, *DonateCoinRequest, *DonateCoinResponse) error
}

func RegisterContractHandler(s server.Server, hdlr ContractHandler, opts ...server.HandlerOption) error {
	type contract interface {
		BalanceOf(ctx context.Context, in *BalanceRequest, out *BalanceResponse) error
		TotalSupply(ctx context.Context, in *TotalSupplyRequest, out *TotalSupplyResponse) error
		Owner(ctx context.Context, in *OwnerRequest, out *OwnerResponse) error
		Decimals(ctx context.Context, in *DecimalsRequest, out *DecimalsResponse) error
		Name(ctx context.Context, in *NameRequest, out *NameResponse) error
		Symbol(ctx context.Context, in *SymbolRequest, out *SymbolResponse) error
		GetGasFee(ctx context.Context, in *GetGasFeeRequest, out *GetGasFeeResponse) error
		GetTransactionByHash(ctx context.Context, in *GetTransactionByHashRequest, out *GetTransactionByHashResponse) error
		Transfer(ctx context.Context, in *TransferRequest, out *TransferResponse) error
		Approve(ctx context.Context, in *ApproveRequest, out *ApproveResponse) error
		TransferFrom(ctx context.Context, in *TransferFromRequest, out *TransferFromResponse) error
		TestError(ctx context.Context, in *Empty, out *Empty) error
		DonateCoin(ctx context.Context, in *DonateCoinRequest, out *DonateCoinResponse) error
	}
	type Contract struct {
		contract
	}
	h := &contractHandler{hdlr}
	return s.Handle(s.NewHandler(&Contract{h}, opts...))
}

type contractHandler struct {
	ContractHandler
}

func (h *contractHandler) BalanceOf(ctx context.Context, in *BalanceRequest, out *BalanceResponse) error {
	return h.ContractHandler.BalanceOf(ctx, in, out)
}

func (h *contractHandler) TotalSupply(ctx context.Context, in *TotalSupplyRequest, out *TotalSupplyResponse) error {
	return h.ContractHandler.TotalSupply(ctx, in, out)
}

func (h *contractHandler) Owner(ctx context.Context, in *OwnerRequest, out *OwnerResponse) error {
	return h.ContractHandler.Owner(ctx, in, out)
}

func (h *contractHandler) Decimals(ctx context.Context, in *DecimalsRequest, out *DecimalsResponse) error {
	return h.ContractHandler.Decimals(ctx, in, out)
}

func (h *contractHandler) Name(ctx context.Context, in *NameRequest, out *NameResponse) error {
	return h.ContractHandler.Name(ctx, in, out)
}

func (h *contractHandler) Symbol(ctx context.Context, in *SymbolRequest, out *SymbolResponse) error {
	return h.ContractHandler.Symbol(ctx, in, out)
}

func (h *contractHandler) GetGasFee(ctx context.Context, in *GetGasFeeRequest, out *GetGasFeeResponse) error {
	return h.ContractHandler.GetGasFee(ctx, in, out)
}

func (h *contractHandler) GetTransactionByHash(ctx context.Context, in *GetTransactionByHashRequest, out *GetTransactionByHashResponse) error {
	return h.ContractHandler.GetTransactionByHash(ctx, in, out)
}

func (h *contractHandler) Transfer(ctx context.Context, in *TransferRequest, out *TransferResponse) error {
	return h.ContractHandler.Transfer(ctx, in, out)
}

func (h *contractHandler) Approve(ctx context.Context, in *ApproveRequest, out *ApproveResponse) error {
	return h.ContractHandler.Approve(ctx, in, out)
}

func (h *contractHandler) TransferFrom(ctx context.Context, in *TransferFromRequest, out *TransferFromResponse) error {
	return h.ContractHandler.TransferFrom(ctx, in, out)
}

func (h *contractHandler) TestError(ctx context.Context, in *Empty, out *Empty) error {
	return h.ContractHandler.TestError(ctx, in, out)
}

func (h *contractHandler) DonateCoin(ctx context.Context, in *DonateCoinRequest, out *DonateCoinResponse) error {
	return h.ContractHandler.DonateCoin(ctx, in, out)
}

// Api Endpoints for Lp service

func NewLpEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Lp service

type LpService interface {
	GetOpsUsdtApy(ctx context.Context, in *GetOpsUsdtApyRequest, opts ...client.CallOption) (*GetApyResponse, error)
	GetOpsFluxAPY(ctx context.Context, in *GetOpsFluxApyRequest, opts ...client.CallOption) (*GetApyResponse, error)
	GetEthUsdtApy(ctx context.Context, in *GetEthUsdtApyRequest, opts ...client.CallOption) (*GetApyResponse, error)
	GetOpsPriceUsdt(ctx context.Context, in *GetOpsPriceRequest, opts ...client.CallOption) (*WorthResponse, error)
	GetOpsPriceFlux(ctx context.Context, in *GetOpsPriceRequest, opts ...client.CallOption) (*WorthResponse, error)
	GetPoolWorth(ctx context.Context, in *GetPoolWorthRequest, opts ...client.CallOption) (*WorthResponse, error)
	GetUserLpBalance(ctx context.Context, in *UserLpRequest, opts ...client.CallOption) (*BalanceResponse, error)
	GetUserLpAmount(ctx context.Context, in *UserLpRequest, opts ...client.CallOption) (*LpAmountResponse, error)
	GetOpsReward(ctx context.Context, in *UserLpRequest, opts ...client.CallOption) (*GetOpsRewardsResponse, error)
}

type lpService struct {
	c    client.Client
	name string
}

func NewLpService(name string, c client.Client) LpService {
	return &lpService{
		c:    c,
		name: name,
	}
}

func (c *lpService) GetOpsUsdtApy(ctx context.Context, in *GetOpsUsdtApyRequest, opts ...client.CallOption) (*GetApyResponse, error) {
	req := c.c.NewRequest(c.name, "Lp.GetOpsUsdtApy", in)
	out := new(GetApyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lpService) GetOpsFluxAPY(ctx context.Context, in *GetOpsFluxApyRequest, opts ...client.CallOption) (*GetApyResponse, error) {
	req := c.c.NewRequest(c.name, "Lp.GetOpsFluxAPY", in)
	out := new(GetApyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lpService) GetEthUsdtApy(ctx context.Context, in *GetEthUsdtApyRequest, opts ...client.CallOption) (*GetApyResponse, error) {
	req := c.c.NewRequest(c.name, "Lp.GetEthUsdtApy", in)
	out := new(GetApyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lpService) GetOpsPriceUsdt(ctx context.Context, in *GetOpsPriceRequest, opts ...client.CallOption) (*WorthResponse, error) {
	req := c.c.NewRequest(c.name, "Lp.GetOpsPriceUsdt", in)
	out := new(WorthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lpService) GetOpsPriceFlux(ctx context.Context, in *GetOpsPriceRequest, opts ...client.CallOption) (*WorthResponse, error) {
	req := c.c.NewRequest(c.name, "Lp.GetOpsPriceFlux", in)
	out := new(WorthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lpService) GetPoolWorth(ctx context.Context, in *GetPoolWorthRequest, opts ...client.CallOption) (*WorthResponse, error) {
	req := c.c.NewRequest(c.name, "Lp.GetPoolWorth", in)
	out := new(WorthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lpService) GetUserLpBalance(ctx context.Context, in *UserLpRequest, opts ...client.CallOption) (*BalanceResponse, error) {
	req := c.c.NewRequest(c.name, "Lp.GetUserLpBalance", in)
	out := new(BalanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lpService) GetUserLpAmount(ctx context.Context, in *UserLpRequest, opts ...client.CallOption) (*LpAmountResponse, error) {
	req := c.c.NewRequest(c.name, "Lp.GetUserLpAmount", in)
	out := new(LpAmountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lpService) GetOpsReward(ctx context.Context, in *UserLpRequest, opts ...client.CallOption) (*GetOpsRewardsResponse, error) {
	req := c.c.NewRequest(c.name, "Lp.GetOpsReward", in)
	out := new(GetOpsRewardsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Lp service

type LpHandler interface {
	GetOpsUsdtApy(context.Context, *GetOpsUsdtApyRequest, *GetApyResponse) error
	GetOpsFluxAPY(context.Context, *GetOpsFluxApyRequest, *GetApyResponse) error
	GetEthUsdtApy(context.Context, *GetEthUsdtApyRequest, *GetApyResponse) error
	GetOpsPriceUsdt(context.Context, *GetOpsPriceRequest, *WorthResponse) error
	GetOpsPriceFlux(context.Context, *GetOpsPriceRequest, *WorthResponse) error
	GetPoolWorth(context.Context, *GetPoolWorthRequest, *WorthResponse) error
	GetUserLpBalance(context.Context, *UserLpRequest, *BalanceResponse) error
	GetUserLpAmount(context.Context, *UserLpRequest, *LpAmountResponse) error
	GetOpsReward(context.Context, *UserLpRequest, *GetOpsRewardsResponse) error
}

func RegisterLpHandler(s server.Server, hdlr LpHandler, opts ...server.HandlerOption) error {
	type lp interface {
		GetOpsUsdtApy(ctx context.Context, in *GetOpsUsdtApyRequest, out *GetApyResponse) error
		GetOpsFluxAPY(ctx context.Context, in *GetOpsFluxApyRequest, out *GetApyResponse) error
		GetEthUsdtApy(ctx context.Context, in *GetEthUsdtApyRequest, out *GetApyResponse) error
		GetOpsPriceUsdt(ctx context.Context, in *GetOpsPriceRequest, out *WorthResponse) error
		GetOpsPriceFlux(ctx context.Context, in *GetOpsPriceRequest, out *WorthResponse) error
		GetPoolWorth(ctx context.Context, in *GetPoolWorthRequest, out *WorthResponse) error
		GetUserLpBalance(ctx context.Context, in *UserLpRequest, out *BalanceResponse) error
		GetUserLpAmount(ctx context.Context, in *UserLpRequest, out *LpAmountResponse) error
		GetOpsReward(ctx context.Context, in *UserLpRequest, out *GetOpsRewardsResponse) error
	}
	type Lp struct {
		lp
	}
	h := &lpHandler{hdlr}
	return s.Handle(s.NewHandler(&Lp{h}, opts...))
}

type lpHandler struct {
	LpHandler
}

func (h *lpHandler) GetOpsUsdtApy(ctx context.Context, in *GetOpsUsdtApyRequest, out *GetApyResponse) error {
	return h.LpHandler.GetOpsUsdtApy(ctx, in, out)
}

func (h *lpHandler) GetOpsFluxAPY(ctx context.Context, in *GetOpsFluxApyRequest, out *GetApyResponse) error {
	return h.LpHandler.GetOpsFluxAPY(ctx, in, out)
}

func (h *lpHandler) GetEthUsdtApy(ctx context.Context, in *GetEthUsdtApyRequest, out *GetApyResponse) error {
	return h.LpHandler.GetEthUsdtApy(ctx, in, out)
}

func (h *lpHandler) GetOpsPriceUsdt(ctx context.Context, in *GetOpsPriceRequest, out *WorthResponse) error {
	return h.LpHandler.GetOpsPriceUsdt(ctx, in, out)
}

func (h *lpHandler) GetOpsPriceFlux(ctx context.Context, in *GetOpsPriceRequest, out *WorthResponse) error {
	return h.LpHandler.GetOpsPriceFlux(ctx, in, out)
}

func (h *lpHandler) GetPoolWorth(ctx context.Context, in *GetPoolWorthRequest, out *WorthResponse) error {
	return h.LpHandler.GetPoolWorth(ctx, in, out)
}

func (h *lpHandler) GetUserLpBalance(ctx context.Context, in *UserLpRequest, out *BalanceResponse) error {
	return h.LpHandler.GetUserLpBalance(ctx, in, out)
}

func (h *lpHandler) GetUserLpAmount(ctx context.Context, in *UserLpRequest, out *LpAmountResponse) error {
	return h.LpHandler.GetUserLpAmount(ctx, in, out)
}

func (h *lpHandler) GetOpsReward(ctx context.Context, in *UserLpRequest, out *GetOpsRewardsResponse) error {
	return h.LpHandler.GetOpsReward(ctx, in, out)
}
