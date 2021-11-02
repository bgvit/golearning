package main

import (
	"fmt"

	"github.com/cod3rcursos/html"
)

func main() {
	c := juntar(
		html.Titulo("https://www.google.com.br/", "https://tecnoblog.net/meiobit/"),
		html.Titulo("https://www.youtube.com.br/", "https://www.amazon.com"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// multiplexar - misturar (mensagens) num canal
func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}
