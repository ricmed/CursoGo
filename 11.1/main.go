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

func main() {
	ricardo := Cliente{
		Nome:  "Ricardo",
		Idade: 43,
		Ativo: true,
	}
	ricardo.Ativo = false
	ricardo.Endereco.Logradouro = "Rua das Flores"

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", ricardo.Nome, ricardo.Idade, ricardo.Ativo)

}
