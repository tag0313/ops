
store_build:
	solc --bin store.sol -o store.bin
	solc --abi store.sol -o store.abi
	abigen --bin=store.bin/Store.bin  --abi=store.abi/Store.abi --pkg=store --out=store.go
	rm store.go

store_clean:
	rm -rf store.abi store.bin store.go
