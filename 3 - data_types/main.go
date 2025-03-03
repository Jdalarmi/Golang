package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int64 = -100
	fmt.Println(number)

	var number2 uint32 = 1000
	fmt.Println(number2)

	//alias
	//INT32 = RUNE
	var number3 rune = 123456
	fmt.Println((number3))

	var number4 byte = 123
	fmt.Println((number4))

	var numberReal float64 = 123000.45
	println(numberReal)

	numberReal2 := 123456.76
	fmt.Println((numberReal2))

	var str string = "texto"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	var text string
	fmt.Println(text)

	var booleano bool = false
	fmt.Println(booleano)

	var erro error = errors.New("Internal error")
	fmt.Println(erro)

}
