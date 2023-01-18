package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto) // Mantem a referencia inicial
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
