
all: build wasm web

run:
	go run ./cmd/main

run-debug:
	go run -tags="example,ebitendebug" .\cmd\main\

build:
	go build -o .dist/spritely ./cmd/main

test:
	go test ./...

wasm:
	GOOS=js GOARCH=wasm go build -o .dist/spritely.wasm ./cmd/main

web:
	mkdir -p .dist \
		&& cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js .dist/ \
		&& cp -R assets .dist/ \
		&& cp dev/index.html .dist/ \
		&& cp dev/main.html .dist/

clean:
	rm -rf .dist/
 