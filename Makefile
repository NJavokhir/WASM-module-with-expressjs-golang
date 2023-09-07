start:
	npm run start

update_wasm:
	GOARCH=wasm GOOS=js go build --o public/wasm/main.wasm public/wasm/main.go