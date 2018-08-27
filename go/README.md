# Compare pure JS and go wasm in browser

## Prerequisites

* Go 1.11

## Building

```bash
GOOS=js GOARCH=wasm go build -o serve/test.wasm main.go
```

## Testing by yourself

```bash
go run go/serve/main.go
```

Results in firefox 62.0b20
--------------------------

JS: 7922ms
Go:  156ms

Useless test, but impressive results.