package main

import (
	"net/http"

	"github.com/Alura/AppWeb/routes"
)

func main() {

	routes.CarregaRotas()             //* buscar na pasta routes, a func que carrega as rotas
	http.ListenAndServe(":8000", nil) //* subir nosso servidos

}
