start:
	npm run start

wasm:
	GOARCH=wasm GOOS=js go build -o golang/wasm/main.wasm golang/wasm/main.go

server: golang/server.go
	go run $<