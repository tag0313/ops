package main

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/debug/trace"
	"ops/pkg/daemon"
	"ops/pkg/version"

	//"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"ops/pkg/logger"
	"ops/pkg/utils"
	"ops/proto/contract"
	"ops/proto/ethereum"
	"ops/proto/nft1155"
	"ops/proto/swap"
	"ops/service/contract/handler"
)

var(
)
type serverTracer struct {
	defaultTracer trace.Tracer
}

func (s *serverTracer) Start(ctx context.Context, name string) (context.Context, *trace.Span) {
	defCtx, defSpan := s.defaultTracer.Start(ctx, name)
	logger.Info("context", defCtx, "span", defSpan)

	return trace.ToContext(defCtx, defSpan.Trace, defSpan.Id), defSpan
}

// Finish the trace
func (s *serverTracer) Finish(span *trace.Span) error {
	logger.Info("finish........")
	return s.defaultTracer.Finish(span)
}

// Read the traces
func (s *serverTracer) Read(opts ...trace.ReadOption) ([]*trace.Span, error) {
	return s.defaultTracer.Read(opts...)
}

func main() {
	flag := daemon.NewCmdFlags()
	if err := flag.Parse(); err != nil{
		panic(err)
	}
	viper, err := utils.LoadConfigViper(flag.ConfigFile)
	if err != nil{
		panic(err)
	}

	loggerOpt := logger.NewOpts(utils.GetConfigStr("log_path"))
	loggerOpt.Debug = true
	logger.InitDefault(loggerOpt)
	defer logger.Sync()
	logger.Info(version.Long)

	contractOptions := handler.NewOptions(viper)
	st := &serverTracer{trace.DefaultTracer}
	consulReg := consul.NewRegistry(
		registry.Addrs(contractOptions.ListenAddr))
	srv := micro.NewService(
		micro.Name(contractOptions.MicroName),
		micro.Registry(consulReg),
		micro.Version(contractOptions.Version),
		micro.Tracer(st),
	)
	handler.SetPBClient(srv.Client())

	err = handler.InitContract(contractOptions)
	if err != nil {
		logger.Fatal(err)
	}

	if err = handler.SubscribeChainEvent(); err != nil {
		logger.Fatal(err)
	}

	contractHandler, err := handler.NewHandler(contractOptions)
	if err != nil {
		logger.Fatal(err)
	}

	err = pbContract.RegisterContractHandler(srv.Server(), contractHandler)
	if err != nil {
		return
	}

	//Create service
	nft1155Handler, err := handler.NewNft1155Handler( contractOptions)
	if err != nil {
		logger.Fatal(err)
	}
	err = pbNft1155.RegisterNFT1155Handler(srv.Server(), nft1155Handler)
	if err != nil {
		logger.Fatal(err)
	}

	ethHandler, err := handler.NewEthHandler(contractOptions)
	if err != nil {
		logger.Fatal(err)
	}
	err = pbEthereum.RegisterEthereumHandler(srv.Server(), ethHandler)
	if err != nil {
		logger.Fatal(err)
	}

	swapHandler, err := handler.NewSwapHandler(contractOptions)
	if err != nil {
		logger.Fatal(err)
	}
	if err = pbSwap.RegisterSwapHandler(srv.Server(), swapHandler); err != nil {
		logger.Fatal(err)
	}

	lpHandler, err := handler.NewLP(contractOptions)
	if err != nil{
		logger.Fatal(err)
		return
	}
	if err = pbContract.RegisterLpHandler(srv.Server(), lpHandler); err != nil{
		logger.Fatal(err)
		return
	}

	// Run service
	if err = srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
