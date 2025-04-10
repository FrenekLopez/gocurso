package main

import (
	"fmt"

	"github.com/freneklopez/gocurso/goroutines"
)

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
	/*numero, mensaje := ejercicios.ConvirtiendoaNumero("rfrf")
	  fmt.Println(numero)
	  fmt.Println(mensaje)*/
	/*teclado.IngresoNumeros()*/
	//fmt.Println(ejercicios.Multiplicacion())

	// files.GrabaTabla()
	// files.SumaTabla()
	// files.LeoArchivo()
	// funciones.Calculo()
	// funciones.LlamarClosure()
	//funciones.Exponencia(2)
	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.MuestroSlices()
	//arreglos_slices.Capacidad()
	//mapas.MuestrarMapas()
	//users.AltaUsusario()

	//Pedro := new(modelos.Hombre)
	//e.HumanoRespirando(Pedro)

	//Maria := new(modelos.Mujer)
	//e.HumanoRespirando(Maria)

	//defer_panic.VemosDefer()
	//defer_panic.EjemploPanic()

	go goroutines.MiNombreLento("Eric Frenek")

	fmt.Println("Este estoy aqui")
	var x string
	fmt.Scanln(&x)
}
