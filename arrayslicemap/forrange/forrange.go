package main

import "fmt"

func main() {
	numeros := [...]int{2, 4, 8, 16, 32} // compilador que conta a quantidade de elementos

	for i, numero := range numeros {
		fmt.Printf("[%d] %d\n", i, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
