package main

import "fmt"

func main() {
	var variavel1 string = "teste de variável"
	variavel2 := "teste de variável"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Teste1"
		variavel4 string = "Teste2"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "teste3", 12

	fmt.Println(variavel5, variavel6)

	const constante1 string = "teste"
	fmt.Println(constante1)

	variavel1, variavel2 = variavel2, variavel1
	fmt.Println(variavel1, variavel2)
}
