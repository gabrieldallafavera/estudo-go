package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Teste")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Aqui")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // espera chegar em 0
}
