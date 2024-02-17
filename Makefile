ENV := GOARCH=wasm GOOS=js

build:
	$(ENV) go build -o _site/web/app.wasm
	cp styles.css _site/web/
	cp icon.png _site/web/
	cp icon.svg _site/web/
	go run ./

vim:
	$(ENV) nvim

.PHONY: serve
serve:
	$(ENV) go build -o web/app.wasm
	cp styles.css web/
	cp icon.png web/
	cp icon.svg web/
	go run . --serve

test:
	go test ./...
