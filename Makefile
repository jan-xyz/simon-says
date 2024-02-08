build:
	go run ./

vim:
	GOOS=js GOARCH=wasm nvim

.PHONY: serve
serve:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	cp styles.css web/
	go run . --serve
