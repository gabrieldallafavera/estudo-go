package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{ // map[tipo das chaves]tipo dos valores
		"nome":      "Teste",
		"sobrenome": "Aqui",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Teste",
			"segundo":  "Aqui",
		},
	}
	fmt.Println(usuario2["nome"]["primeiro"])

	delete(usuario2, "nome") // Remove valores de uma chave
}
