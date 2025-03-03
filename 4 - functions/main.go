package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculateMatematic(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

func main() {
	sum_total := sum(10, 20)
	fmt.Println((sum_total))

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Teste func")

	result_sum, _ := calculateMatematic(10, 15)
	fmt.Println(result_sum)
}
