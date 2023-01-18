package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("string")
	generica(1)
	generica(true)

	fmt.Println(1, 2, "string", false, true, float64(12345))

	mapa := map[interface{}]interface{}{
		1:           "Teste",
		float32(34): true,
		"String":    "String",
		true:        float64(345),
	}

	fmt.Println(mapa)
}
