package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro": 20000.0,
		},
	}

	delete(funcsPorLetra, "P")

	for _, funcionario := range funcsPorLetra {
		for nome, salario := range funcionario {
			fmt.Println(nome, salario)
		}
	}
}
