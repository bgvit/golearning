package main

import "fmt"

func main() {

	slice_original := []int{1, 2, 3, 4, 5}
	slice1 := slice_original
	//fmt.Println(slice_original, slice1)
	slice_original[2] = 1
	//fmt.Println(slice_original,slice1)

	slice_original[3] = 1000
	fmt.Println(slice1)

}
