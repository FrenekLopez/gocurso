package ejer_interfaces

import (
	"fmt"

	"github.com/freneklopez/gocurso/interfaces"
)

func HumanoRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}
