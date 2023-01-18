package main

import (
	"errors"
	"fmt"
)

func main() {
	// pode ser também int8, int16, int32, int64
	var numero int = -10000
	fmt.Println(numero)

	// unasigned int, não pode usar sinal
	var numero2 uint = 1000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// UINT8 = BYTE
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroF1 float32 = 123.134
	fmt.Println(numeroF1)

	var numeroF2 float64 = 122332233.134
	fmt.Println(numeroF2)

	numeroF3 := 3245.33
	fmt.Println(numeroF3)

	var str string = "Teste"
	fmt.Println(str)

	str2 := "teste2"
	fmt.Println(str2)

	var str3 string
	fmt.Println(str3)

	var numero5 int16
	fmt.Println(numero5)

	var boolean bool = true
	fmt.Println(boolean)

	var erro error = errors.New("Inserindo novo erro")
	fmt.Println(erro)
}
