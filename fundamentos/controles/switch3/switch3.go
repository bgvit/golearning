package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println(("teste"))
	fmt.Println(tipo(6.3))
	fmt.Println(tipo(func (){}))
	fmt.Println(tipo(time.Now()))
}

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "tipo que não conheço"
	}
}