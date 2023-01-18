package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1, v2)

	v1++
	fmt.Println(v1, v2)

	// PONTEIRO É UMA REFERENCIA DE MEMÓRIA
	var v3 int
	var pont *int

	v3 = 100
	pont = &v3

	fmt.Println(v3, pont)
	fmt.Println(v3, *pont)
}
