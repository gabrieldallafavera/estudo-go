package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	stringConexao := "usuario:senha@/banco?charset=utf8&parseTime=True&loc=local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}
	// defer db.Close()

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
