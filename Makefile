ENV := GOARCH=wasm GOOS=js

clean:
	rm -rf _site/ web/

build:
	$(ENV) go build -o _site/web/app.wasm
	cp styles.css _site/web/
	cp icon.png _site/web/
	cp icon.svg _site/web/
	go run .

vim:
	$(ENV) nvim

.PHONY: serve
serve:
	# install wasm toolchain: `rustup target install wasm32-unknown-unknown`
	# install wasm-server-runner: `cargo install wasm-server-runner`
	CARGO_TARGET_WASM32_UNKNOWN_UNKNOWN_RUNNER=wasm-server-runner cargo run --target wasm32-unknown-unknown

test:
	go test ./...

lint:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck ./...
	go run github.com/mgechev/revive -set_exit_status ./...
