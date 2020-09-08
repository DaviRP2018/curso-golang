package main

import "fmt"

func fatorial(n uint) (result uint) {
	result = 1
	for n > 0 {
		n--
		result += result * n
	}
	return
}

func fatorialDoLeo(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}
