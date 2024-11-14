package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	ricardo := Cliente{
		Nome:  "Ricardo",
		Idade: 43,
		Ativo: true,
	}
	ricardo.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\n", ricardo.Nome, ricardo.Idade, ricardo.Ativo)

}
