

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