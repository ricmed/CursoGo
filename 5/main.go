package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(meuArray)
	fmt.Println(len(meuArray))
	fmt.Println(meuArray[len(meuArray)-1])

	for i, valor := range meuArray {
		fmt.Printf("O valor do índice %d é %d\n", i, valor)
	}
}
