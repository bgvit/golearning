package main

import "fmt"

func main() {
	x, y := troca(2, 3)
	fmt.Println(x, y)
	
	segundo, primeiro := troca(7,1)
	fmt.Println(segundo, primeiro)
}

func troca(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
}
