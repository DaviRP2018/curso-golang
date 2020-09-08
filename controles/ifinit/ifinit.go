package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	fmt.Println("s", s)
	r := rand.New(s)
	fmt.Println("r", r)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { // tb suportado no switch
		fmt.Println("Ganhou!!!")
	} else {
		fmt.Println("Perdeu!")
	}

	// i é indisponível fora do escopo do if else
	// fmt.Println(i)
}
