package main

import (
	"fmt"

	"github.com/odelacruz93/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvierteTexto(1993)
	fmt.Println(estado)
	fmt.Println(texto)
}
