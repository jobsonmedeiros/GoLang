package main

import (
	"fmt"
	"math"
)

// Go suporta constantes de caracteres, strings,
// valores booleanos e numéricos.

const s string = "constante" // const declara um valor constante

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d)) // Uma constante numérica não tem
	// tipo até receber um, como por uma conversão explícita.

	fmt.Println(math.Sin(n)) // Aqui math.sin espera um valor
	// float 64
}
