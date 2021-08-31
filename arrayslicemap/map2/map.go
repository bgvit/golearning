package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"José":         11325.45,
		"Gab":          15456.78,
		"Pedro Junior": 1200.0}
	funcESalarios["Rafael Filho"] = 1350.0
	delete(funcESalarios, "inexistente")
	//No for, o valor sempre será a chave, valor. Se usar _, salario -> valor, se usar nome -> chave
	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}
}
