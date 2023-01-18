package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	fmt.Println(texto, numeros)
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6))

	escrever("Teste de escrita", 1, 2, 3, 4, 5)
}
