package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// revisão: nesse caso, * é usado para desreferenciar,
	// ou seja, nos dar acesso ao valor do qual o ponteiro aponta
	*n++
}

func main() {
	n := 1
	inc1(n) // passagem de parâmetro por valor
	fmt.Println(n)

	// revisão: & usado para obter o endereço da variável
	inc2(&n) // passagem de parâmetro por referência
	fmt.Println(n)
}
