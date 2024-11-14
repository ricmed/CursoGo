package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("Cliente %s desativado\n", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	ricardo := Cliente{
		Nome:  "Ricardo",
		Idade: 43,
		Ativo: true,
	}

	minhaEmpresa := Empresa{}

	Desativacao(ricardo)
	Desativacao(minhaEmpresa)

}
