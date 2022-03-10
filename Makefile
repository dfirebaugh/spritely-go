
all: build build-windows wasm web

run:
	go run ./cmd/main

run-debug:
	go run -tags="example,ebitendebug" .\cmd\main\

build:
	go build -o .dist/spritely ./cmd/main

build-windows: # cross-compile to windows exe
	GOOS=windows go build -o .dist/spritely.exe ./cmd/main

test:
	go test ./...

wasm:
	GOOS=js GOARCH=wasm go build -o .dist/spritely.wasm ./cmd/main

.PHONY: web
web:
	mkdir -p .dist \
		&& cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js .dist/ \
		&& cp -R assets .dist/ \
		&& cp web/index.html .dist/ \
		&& cp web/main.html .dist/

clean:
	rm -rf .dist/
 