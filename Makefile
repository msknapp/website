

.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -o static/wasm/app.wasm wasm/main.go

.PHONY: site
site:
	hugo

.PHONY: serve
serve:
	hugo serve -p 8080

.PHONY: cheats
cheats:
	./build-cheats

test: vet
	go test ./pkg/...

coverage: vet
	go test -coverprofile=coverage.out ./pkg/...
	go tool cover -html=coverage.out

vet:
	go vet ./pkg/...

code-to-html:
	go build -o code-to-html ./cmd/code-to-html
	[ -d ${HOME}/bin ] && mv code-to-html ${HOME}/bin