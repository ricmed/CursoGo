package main

import "fmt"

type ID int

var (
	b bool = true
	c int
	d string = "Ricardo"
	e float64
	g ID = 1
)

func main() {
	var a string
	f := "X"
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
	println(g)

	fmt.Printf("O tipo de E é %T", e)
}
