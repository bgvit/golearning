package main

import (
	"fmt"
	"time"
)

func main() {

	// Channel (canal) -> é a forma de comunicação entre goroutines
	// É um ponto de parada (obriga sincronismo)
	// channel é um tipo
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b := <-c, <-c // recebendo os dados do canal
	fmt.Println(a, b)
	fmt.Println(<-c)
}

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base
	
	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

