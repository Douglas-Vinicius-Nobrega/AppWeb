//* Controle.

package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Alura/AppWeb/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html")) //* variavel com todas os templates, carrega todos os templates

func Index(w http.ResponseWriter, r *http.Request) { //* func, requisita os produtos e e exibi o template

	TodosOsProdutos := models.BuscaTodosOsProdutos() //* armazena na variavel todos os produtos

	temp.ExecuteTemplate(w, "Index", TodosOsProdutos) //* exexutar o template
}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil) //* abrir o novo template, com o nosso novo formulário

}

func Insert(w http.ResponseWriter, r *http.Request) { //* inserir nosso produto no banco de dados
	if r.Method == "POST" { //* se essa requisição for igual a POST, estou querendo criar um novo produto
		nome := r.FormValue("nome") //* Forma de conseguir buscar , a informação que queremos
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64) //* essa var, retorna o valor convertido de string, para Float

		if err != nil { // se o erro não for igual a nil
			log.Println("Erro na conversão do preço: ", err)
		}

		quantidadeConvertidoParaInt, err := strconv.Atoi(quantidade) //* essa var, retorna o valor convertido de string, para int

		if err != nil { // se o erro não for igual a nil
			log.Println("Erro na conversão da quantidade: ", err)
		}

		models.CriaNovoProduto(nome, descricao, precoConvertidoParaFloat, quantidadeConvertidoParaInt)
	}

	http.Redirect(w, r, "/", 301) //* w = responsiveWrite; r = Request. Assim que preencher o formulário, retorna para a página principal.
}

func Delete(w http.ResponseWriter, r *http.Request) { //* Função que deleta um item
	idDoProduto := r.URL.Query().Get("id") //* Pegando uma informação da minha URL
	models.DeletaProduto(idDoProduto)      //* deletar o produto do banco de dados
	http.Redirect(w, r, "/", 301)          //* depois de deletar, redirecionar para a página principal
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id") //* Trazemos os parametro Id do nosso URl
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto) //* Função que ira editar os produtos
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID para int: ", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço para float64: ", err)
		}

		quantidadeConvertidoParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão do quantidade para Int: ", err)
		}

		models.AtualizarProduto(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidoParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
