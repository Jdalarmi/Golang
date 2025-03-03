package main

import (
	"fmt"
	"module/assistent"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	assistent.Write()

	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Print(erro)
}
