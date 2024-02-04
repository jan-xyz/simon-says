clean:
	go clean
	rm -rf public


build: wasm_exec.js
	GOOS=js GOARCH=wasm go build -o main.wasm

bundle:
	mkdir -p public
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" ./public/
	cp ./main.wasm ./public/
	cp ./index.html ./public/
	cp ./favicon.ico ./public/

vim:
	GOOS=js GOARCH=wasm nvim

.PHONY: serve
serve: build bundle
	go run ./serve/main.go -dir public
