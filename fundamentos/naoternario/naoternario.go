package main

import "fmt"

//Não tem operador ternario
//A única maneira de trabalhar com if é da forma abaixo:
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main(){
	fmt.Println(obterResultado(6.2))
}
