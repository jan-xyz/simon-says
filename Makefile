ENV := GOARCH=wasm GOOS=js

clean:
	rm -rf _site/ web/

build:
	$(ENV) go build -o _site/web/app.wasm
	cp styles.css _site/web/
	cp icon.png _site/web/
	cp icon.svg _site/web/
	cp stats.png _site/web/
	go run .

vim:
	$(ENV) nvim

.PHONY: serve
serve:
	$(ENV) go build -o web/app.wasm
	cp styles.css web/
	cp icon.png web/
	cp icon.svg web/
	cp stats.png web/
	go run . --serve

test:
	go test ./...

lint:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck ./...
	go run github.com/mgechev/revive -set_exit_status ./...
