//* Funções de configurações, de acesso ao banco de dados.

package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB { //* Conectando com o banco de dados
	conexao := "user=postgres dbname=alura_loja password=dNvidia11 host=localhost sslmode=disable" //* passando dados do usuario
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
