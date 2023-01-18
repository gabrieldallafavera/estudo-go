package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // utilização de pacote implicito, quem utiliza é o database/sql
)

func main() {
	stringConexao := "usuario:senha@/banco?charset=utf8&parseTime=True&loc=local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Conexão aberta!")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
