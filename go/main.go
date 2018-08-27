package main

import "fmt"

func main() {
	i := 0
	for i < 100000000 {
		i++
	}
	fmt.Println("this is go")
}

// https://github.com/golang/go/wiki/WebAssembly
// cp $(go env GOROOT)/misc/wasm/wasm_exec.{html,js} .
// GOOS=js GOARCH=wasm go build -o serve/test.wasm main.go
