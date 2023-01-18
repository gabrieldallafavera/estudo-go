package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Teste1", "Teste2", "Teste3"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "Teste" {
		fmt.Println(indice, string(letra))
	}
}
