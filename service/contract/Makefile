DATETIME = `date '+%Y%m%d_%H%M'`
.PHONY: proto build run test build_ops_erc20 docker clean pack_code

build:
	go build -o contract main.go

run:
	go run main.go

test:
	go test -v ./...


build_ops_erc20:
	mkdir -p ercpkg/ops ercpkg/nft1155 ercpkg/nftholder ercpkg/swap ercpkg/erc20basic ercpkg/opslp
	abigen --abi=ops/OpsErc20.abi --bin=ops/OpsErc20.bin --pkg=ops --out=ercpkg/ops/ops.go
	abigen --abi=ops/OpsNft1155.abi --bin=ops/OpsNft1155.bin --pkg=nft1155 --out=ercpkg/nft1155/nft.go
	abigen --abi=ops/OpsHolder.abi --bin=ops/OpsHolder.bin --pkg=nftholder --out=ercpkg/nftholder/holder.go
	abigen --abi=ops/UniswapV2Pair.abi --bin=ops/UniswapV2Pair.bin --pkg=swap --out=ercpkg/swap/UniswapV2Pair.go
	abigen --abi=ops/Erc20Basic.abi --pkg=erc20basic --out=ercpkg/erc20basic/erc20basic.go
	abigen --abi=ops/OpsLpRewards.abi --pkg=opslp --out=ercpkg/opslp/OpsLpRewards.go

docker: linux_build
	./push_image.sh contract-${DATETIME}

linux_build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o contract main.go

pack_code: clean
	tar -czvf contract-${DATETIME}.tar.gz *

clean:
	@rm -f contract contract*.tar.gz
