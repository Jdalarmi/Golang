package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays and Slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [2]string{"Posição 1", "Posição 2"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} //Fixa o tamanho de acordo com que for passado
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15} //tamanho muda de acordo com necessidade
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)
}
