# Compare pure JS and go wasm in browser

## Prerequisites

* Go 1.11

## Just look at it for yourself

[Here on surge!](http://playwithwasm-golang.surge.sh/)

## Building

```bash
GOOS=js GOARCH=wasm go build -o serve/test.wasm main.go
```

## Testing by yourself

```bash
go run go/serve/main.go
```

## Results

go to http://127.0.0.1:8080/go/serve/

Results in firefox 62.0b20:
* JS: 7922ms
* Go:  156ms

Useless test, but impressive results.