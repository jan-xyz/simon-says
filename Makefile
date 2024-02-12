build:
	GOARCH=wasm GOOS=js go build -o _site/web/app.wasm
	cp styles.css _site/web/
	cp icon.png _site/web/
	go run ./

vim:
	GOOS=js GOARCH=wasm nvim

.PHONY: serve
serve:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	cp styles.css web/
	cp icon.png web/
	go run . --serve
