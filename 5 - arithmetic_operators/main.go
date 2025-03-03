package main

import "fmt"

func main() {
	var number1 int16 = 10
	var number2 int16 = 25

	soma := number1 + number2
	fmt.Println(soma)

	//ATRIBUIÇÃO
	variable2 := "String"
	fmt.Println(variable2)

	// RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//OPERADORES LOGICOS
	// fmt.Println(true && true)
	fmt.Print(true || false)
	fmt.Print(!true)

	//OPERADORES UNARIOS
	number := 10
	number++ // aumenta 1
	number += 15
	number--
	fmt.Println(number)

	//OPERADORES TERNARIOS
	var text string
	if number > 5 {
		text = "Maior que 5"
	} else {
		text = "Menor que 5"
	}
	fmt.Println(text)
}
