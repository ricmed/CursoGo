package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
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

func main() {
	ricardo := Cliente{
		Nome:  "Ricardo",
		Idade: 43,
		Ativo: true,
	}

	ricardo.Desativar()

}
