package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Teste de canal"
	canal <- "Teste 2"

	fmt.Println(<-canal)
	fmt.Println(<-canal)
}
