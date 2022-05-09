//* Faz a conexão com o postgre e busca todos os produtos.

package models

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := ConectaComBancoDeDados() //* abrindo conexão com banco de dados

	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) { //* Os parametros para minha função funcionar
	db := ConectaComBancoDeDados() //* abrindo conexão com o banco de dados

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)") //* inserindo conteudo no banco de dados e o nosso script verifica se esta dando certo ou não

	if err != nil { //* se o erro for defirente de nulo
		panic(err.Error()) //* exibindo o erro
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade) //* Executar com o valores que estão no produtos, em controller
	defer db.Close()
}

func DeletaProduto(id string) {
	db := ConectaComBancoDeDados() //* conecta com o banco de dados

	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil { //* se o erro for defirente de nulo
		panic(err.Error()) //* exibindo o erro
	}

	deletarOProduto.Exec(id) //* usando esse script, pegando o valor do id

	defer db.Close()

}

func EditaProduto(id string) Produto {
	db := ConectaComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade

	}
	defer db.Close()
	return produtoParaAtualizar

}

func AtualizarProduto(id int, nome, decricao string, preco float64, quantidade int) {
	db := ConectaComBancoDeDados()

	AtualizarProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5") //* Preparar o banco de dados
	if err != nil {
		panic(err.Error())
	}
	AtualizarProduto.Exec(nome, decricao, preco, quantidade, id)
	defer db.Close()
}
