package main

import (
	"bytes"
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
	c := cachorro{"Nome do Cachorro", "Raça do Cachorro", 12}

	cJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal("Erro ao converter")
	}

	fmt.Println(cJSON)
	fmt.Println(bytes.NewBuffer(cJSON))

	c2 := map[string]string{
		"nome": "Teste 2",
		"raca": "Teste de Raça 2",
	}

	c2JSON, erro := json.Marshal(c2)
	if c2JSON != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(c2JSON))
}
