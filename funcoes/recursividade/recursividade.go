package main

import "fmt"

func main() {

	resultado, _ := fatorial(5)
	fmt.Println(resultado)

	_, err := fatorial(-4)

	if err != nil {
		fmt.Println(err)
	}

}

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}

}
