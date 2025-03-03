package main

import "fmt"

func main() {
	var variable_one string = "Varivél 1"
	variable_two := "Variavel 2"

	fmt.Print(variable_one)
	fmt.Print((variable_two))

	var (
		variable_three string = "lalala"
		variable_four  string = "lalala"
	)

	fmt.Print(variable_three, variable_four)

	variable_five, variable_six := "Varivél 5", "Varivél 6"
	fmt.Print(variable_five, variable_six)

	const constant_one string = "Constant 1"
	fmt.Print(constant_one)
}
