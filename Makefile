build-wasi:
	GOOS=wasip1 GOARCH=wasm go build -o main.wasm main.go

build-wasm:
	GOOS=js GOARCH=wasm go build -o main.wasm
