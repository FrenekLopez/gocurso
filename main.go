package main

import (
	"fmt"

	"github.com/freneklopez/gocurso/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1577)
	fmt.Println(estado)
	fmt.Println(texto)
}
