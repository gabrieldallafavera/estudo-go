package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cJSON := `{"nome":"Teste1","raca":"Teste de Raça","idade":12}`
	var c cachorro
	if erro := json.Unmarshal([]byte(cJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	c2JSON := `{"nome":"Teste2","raca":"Teste de Raça 2"}`
	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(c2JSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
}
