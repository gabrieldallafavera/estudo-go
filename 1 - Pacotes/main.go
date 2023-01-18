package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Teste de escrita")
	auxiliar.Escrever()

	error := checkmail.ValidateFormat("teste")
	fmt.Println(error)
}
