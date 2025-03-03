package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	number uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u user
	u.name = "jeferson"
	u.age = 21
	fmt.Println(u)

	address_example := address{109}

	user2 := user{"Davi", 21, address_example}
	fmt.Println(user2)

	user3 := user{name: "teste"}
	fmt.Println((user3))
}
