package main

import "fmt"

var soma = func(a, b int) int {
	return a+b
}


func main() {
	fmt.Println(2,3)

	teste := soma
	fmt.Println(teste(3,54))
	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(2,3))
}