package main

import "fmt"

func main() {
	func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 10)
	}("teste de texto")
}
