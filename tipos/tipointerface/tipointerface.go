package main

import "fmt"

type curso struct {
	nome string
}


func main() {
	var coisa interface{}
	fmt.Println(coisa)
	
	coisa = 3
	fmt.Println(coisa)

	type dinamico interface{}
	// eu posso utilizar a interface como um tipo dinâmico, permitindo, por exemplo, atribuição de vários valores.
	var coisa2 dinamico = "Opa"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang"}
	fmt.Println(coisa2)
}

