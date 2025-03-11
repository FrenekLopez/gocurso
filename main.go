package main

import (
	"fmt"

	"github.com/freneklopez/gocurso/ejercicios"
	/*"runtime"*/)

func main() {
	/*estado, texto := variables.ConviertoaTexto(1577)
	fmt.Println(estado)
	fmt.Println(texto)*/
	/*if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es Windows", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}*/
	numero, mensaje := ejercicios.ConvirtiendoaNumero("rfrf")
	fmt.Println(numero)
	fmt.Println(mensaje)
}
