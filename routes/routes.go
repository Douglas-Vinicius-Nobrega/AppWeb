//* Rotas.

package routes

import (
	"net/http"

	"github.com/Alura/AppWeb/controllers"
)

func CarregaRotas() {

	http.HandleFunc("/", controllers.Index) //* carregamento das rotas
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert) //* Quando tiver um requisição, para /insert, vou chamar o insert
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit) //*
	http.HandleFunc("/update", controllers.Update)
}
