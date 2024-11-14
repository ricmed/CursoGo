package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum_with_error(20, 35)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)

}

func sum(a int, b int) int {
	return a + b
}

func sum_with_error(a, b int) (int, error) {
	if a+b > 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
