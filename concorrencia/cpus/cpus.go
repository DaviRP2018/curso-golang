package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Quantidade de processadores: %d", runtime.NumCPU())
}
