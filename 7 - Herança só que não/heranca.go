package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa    // recebe os mesmos campos do strct pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("\"Herança\"")

	pessoa := pessoa{"Teste", "teste", 12, 178}

	estudante := estudante{pessoa, "engenharia", "USP"}

	fmt.Println(estudante)
	fmt.Println(estudante.nome)
}
