package main

import "fmt"

func main(){
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(2.1))
}

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
		case 10:
			fallthrough//se o switch cair nesse caso, ele irá executar o bloco de código do imediatamente próximo
		case 9:
			return "A"
		case 8,7:
			return "C"
		case 6,5:
			return "C"
		case 4,3:
			return "C"
		case 2,1,0:
			return "E"
		default:
			return "Nota inválida"
	}
}


// package main

// func main(){
	
// }