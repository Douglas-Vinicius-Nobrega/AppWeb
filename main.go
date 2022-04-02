package main

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) // carregar todos os templates, para o serv

func main() {
	http.HandleFunc("/", index)       // toda vez que tiver uma requisição para a barra(/), quem vai atender é a função index
	http.ListenAndServe(":8000", nil) // subindo um servido, na porta 8000 (8 mil)
}

func index(w http.ResponseWriter, r *http.Request) { // Toda vez que tiver um requisição, vamos ter esses dois parametros
	temp.ExecuteTemplate(w, "Index", nil) // executando nosso template de index dentro dessa função
}
