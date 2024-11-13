package main

import "fmt"

func main() {

	salarios := map[string]int{"Ricardo": 5000, "João": 4000, "Maria": 3000}
	fmt.Println(salarios["Ricardo"])
	delete(salarios, "Maria")

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é de %d\n", salario)
	}
}
