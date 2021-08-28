package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiros. Não se pode fazer qualquer operação de aritmética em cima dos ponteiros
	// p vira a referência
	var p *int = nil

	// o ponteiro i agora aponta para o local na memória em que é guardado o valor da variável i
	p = &i

	// processo de desreferenciar o ponteiro para acessar o valor daquele ponteiro e alterar diretamente na memória(o incremento tá alterando diretamente)
	*p++
	i++
	fmt.Println(p, *p, i, &i)
}

/*
package main

func main(){

}
*/
