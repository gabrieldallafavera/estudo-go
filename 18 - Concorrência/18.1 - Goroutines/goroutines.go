package main

import (
	"fmt"
	"time"
)

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	// CONCORRÊNCIA != PARALELISMO

	go escrever("Teste") // gouroutine
	escrever("Aqui")
}
