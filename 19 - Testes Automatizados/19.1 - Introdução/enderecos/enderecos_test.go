package enderecos_test

import (
	. "introducao-testes/enderecos" // Utilizar o . entende que esse pacote pode ser utilizado por padrão
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// go test --cover
// Mostra a porcentagem de cobertura dos testes

// go test --coverprofile <nomeDoArquivo.txt>
// Gera um arquivo para ser analisado relacionado os testes

// go tool cover --func=<nomeDoArquivo.txt>
// Mostra no terminal o que cobriu

// go tool cover --html=<nomeDoArquivo.txt>
// Mostra em uma pagina html o que faltou cobrir no teste

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel() // Permite rodar os testes em paralelo

	cenarioDeTeste := []cenarioDeTeste{
		{"Rua Teste", "Rua"},
		{"Avenida Teste", "Avenida"},
		{"Rodovia Teste", "Rodovia"},
		{"Estrada Teste", "Estrada"},
		{"Praça Teste", "Tipo Inválido"},
		{"RUA Teste", "Rua"},
		{"ESTRADA Teste", "Estrada"},
	}

	for _, cenario := range cenarioDeTeste {
		if retorno := TipoDeEndereco(cenario.enderecoInserido); retorno != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", retorno, cenario.retornoEsperado)
		}
	}
}
