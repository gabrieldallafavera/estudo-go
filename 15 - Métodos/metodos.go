package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando dados do Usuário %s\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Teste nome", 20}
	usuario1.salvar()

	usuario2 := usuario{"Teste2 nome", 30}
	usuario2.salvar()

	fmt.Println(usuario2.maiorDeIdade())

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
