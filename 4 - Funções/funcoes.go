package main

import (
	"errors"
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) (string, error) {

		if txt == "" {
			return txt, errors.New("String não pode estar vazia.")
		}

		return txt, nil
	}

	retorno, erro := f("Texto de função f")

	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println(retorno)
	}
}
