package main

import "fmt"

func main() {

	i := 1
	for i <= 10 { //não tem while em Go. Esse é o mais próximo.
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 { //o nosso for tradicional
		fmt.Printf("%d", i)
	}

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	// for {
	// 	laço infinito
	// }

	// No capitulo de Array, vamos ver o que é um foreach no Go.

}
