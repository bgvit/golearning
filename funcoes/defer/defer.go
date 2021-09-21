package main

import "fmt"

func main() {
	obterValorAprovado(5001)

}

func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	} 
	fmt.Println("Valor baixo...")
	return numero
}
