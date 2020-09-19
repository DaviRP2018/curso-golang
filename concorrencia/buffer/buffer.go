package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 2
	ch <- 4
	ch <- 8
	ch <- 16
	ch <- 32
	fmt.Println("Executou!")
	ch <- 64
}

func main() {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
