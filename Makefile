build:
	GOARCH=wasm GOOS=js go build -o _site/app.wasm
	cp styles.css _site/

vim:
	GOOS=js GOARCH=wasm nvim

.PHONY: serve
serve: build
	go run .
