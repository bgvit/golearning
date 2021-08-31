package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// mapas devem ser inicializados. A forma abaixo não serve. É preciso usar make
	//var aprovados map[int]string

	aprovados := make(map[int]string)
	aprovados[12345678978] = "Maria"
	aprovados[12325678978] = "Pedro"
	aprovados[112325678978] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d) \n", nome, cpf)
	}

	fmt.Println(aprovados[12325678978])
	delete(aprovados, 12325678978)
	//vazio
	fmt.Println(aprovados[12325678978])
}
