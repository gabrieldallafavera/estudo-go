package main

import "fmt"

type endereco struct {
	rua    string
	numero int16
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	fmt.Println("Campo struct")

	var u usuario
	u.nome = "Teste"
	u.idade = 21
	fmt.Println(u)

	u2 := usuario{
		nome:  "teste",
		idade: 32,
	}
	fmt.Println(u2)

	u3 := usuario{nome: "teste"}
	fmt.Println(u3)

	u4 := usuario{
		"Teste 2",
		32,
		endereco{
			"rua 1",
			12,
		},
	}
	fmt.Println(u4)
}
