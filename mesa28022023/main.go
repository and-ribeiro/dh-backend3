package main

import "fmt"

type IProduto interface {
	setNome(nome string)
	setPreco(preco float64)
	getNome() string
	getPreco() float64
}

type Produto struct {
	nome  string
	preco float64
}

func (p *Produto) setNome(nome string) {
	p.nome = nome
}

func (p *Produto) getNome() string {
	return p.nome
}

func (p *Produto) setPreco(preco float64) {
	p.preco = preco
}

func (p *Produto) getPreco() float64 {
	return p.preco
}

type Pequeno struct {
	Produto
}

type Medio struct {
	Produto
}

type Grande struct {
	Produto
}

func novoPequeno(nome string, custo float64) IProduto {
	return &Pequeno{
		Produto{
			nome:  nome,
			preco: custo,
		},
	}
}

func novoMedio(nome string, custo float64) IProduto {
	return &Medio{Produto{
		nome:  nome,
		preco: custo + (custo * 0.03),
	},
	}
}

func novoGrande(nome string, custo float64) IProduto {
	return &Grande{Produto{
		nome:  nome,
		preco: custo + (custo * 0.06) + 50,
	},
	}
}

func factory(tipoProduto string, nome string, custo float64) (IProduto, error) {
	if tipoProduto == "Pequeno" {
		return novoPequeno(nome, custo), nil
	}
	if tipoProduto == "Medio" {
		return novoMedio(nome, custo), nil
	}
	if tipoProduto == "Grande" {
		return novoGrande(nome, custo), nil
	}
	return nil, fmt.Errorf("Tipo errado de produto")
}

func main() {
	produtinho, _ := factory("Pequeno", "Arroz", 3)
	produto, _ := factory("Medio", "Garrafa 2L", 5)
	produtao, _ := factory("Grande", "Geladeira", 1500)

	fmt.Println(produtinho.getPreco())
	fmt.Println(produto.getPreco())
	fmt.Println(produtao.getPreco())

}
