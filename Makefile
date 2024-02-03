clean:
	go clean
	rm -f wasm_exec.js


wasm_exec.js:
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" .

build: wasm_exec.js
	GOOS=js GOARCH=wasm go build -o main.wasm

.PHONY: serve
serve: build wasm_exec.js
	go run ./serve/main.go
