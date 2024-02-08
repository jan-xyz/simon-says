build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	cp styles.css web/

vim:
	GOOS=js GOARCH=wasm nvim

.PHONY: serve
serve: build
	go run .
