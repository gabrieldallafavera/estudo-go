package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 12
	fmt.Println(array1)

	array2 := [5]string{"teste1", "teste2", "teste3", "teste4", "teste5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // Fixa o tamnho do array na quantidade elementos inseridos
	fmt.Println(array3)

	slice := []int{1, 2, 3} // Valor pode ser dinamico
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 4)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "posição alterada"
	fmt.Println(slice2)

	// Arrays Internos
	slice3 := make([]float32, 10, 15) // tipo, pula tantos, capacidade maxima

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho
	fmt.Println(cap(slice3)) // capacidade

	// Quando o slice atinge a capacidade maxima e insere mais ele dobra a capacidade
}
