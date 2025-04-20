package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays Internos")

	fmt.Println("------------------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}
