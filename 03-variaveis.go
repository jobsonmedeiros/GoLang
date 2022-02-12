package main

import "fmt"

// Não é possível declarar uma variável ou pacote e não usar.
// Go não deixará compilar

func main() {

	var a = "inicial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple" // forma rápida de declarar e iniciar uma variável
	fmt.Println(f)
}
