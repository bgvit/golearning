package main

import (
	"fmt"
	"reflect"
)

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	// poderia ser carro1 := ferraru{"f40", false, 0}, mas deixei assim para visualização da linha 27
	var carro1 ferrari = ferrari{"f40", false, 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2, reflect.TypeOf(carro1), reflect.TypeOf(carro2))
}
