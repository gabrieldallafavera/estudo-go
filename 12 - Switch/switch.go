package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Não encontrado"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch numero {
	case 1:
		diaDaSemana = "Domingo"
		fallthrough
	case 2:
		diaDaSemana = "Segunda"
	case 3:
		diaDaSemana = "Terça"
	case 4:
		diaDaSemana = "Quarta"
	case 5:
		diaDaSemana = "Quinta"
	case 6:
		diaDaSemana = "Sexta"
	case 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Não encontrado"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	fmt.Println(diaDaSemana(2))
}
