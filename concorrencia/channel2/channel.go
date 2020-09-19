package main

import (
	"fmt"
	"time"
)

// Channel (canal) - é a forma de comunicação entre goroutines
// channel é um tipo

func doisTressQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 4 * base // enviando dados para o canal
}

func main() {
	c := make(chan int)
	go doisTressQuatroVezes(10, c)
	go func(base int, c chan int) {
		time.Sleep(time.Second)
		c <- 2 * base // enviando dados para o canal

		time.Sleep(time.Second)
		c <- 3 * base // enviando dados para o canal

		time.Sleep(time.Second)
		c <- 4 * base // enviando dados para o canal
	}(2, c)

	fmt.Println("A")

	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
