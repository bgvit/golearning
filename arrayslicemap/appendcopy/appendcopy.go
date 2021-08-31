package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int
	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1) // copia do slice1 para o slice2 (destiny, source) // copy não expande o tamanho do slice
	fmt.Println(slice2)
}
