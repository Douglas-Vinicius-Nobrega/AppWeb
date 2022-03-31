package main

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*html")) // carregar todos os templates, para o serv

func main() {
	http.HandleFunc("/", index)       // toda vez que tiver uma requisição para a barra, quem vai atender é o index
	http.ListenAndServe(":8000", nil) // subindo um servido, na porta 8000 (8 mil)
}

func index(w http.ResponseWriter, r *http.Request) { //
	temp.ExecuteTemplate(w, "Index", nil)
}
