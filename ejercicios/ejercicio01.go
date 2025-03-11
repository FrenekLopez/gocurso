package ejercicios

import (
	"strconv"
)

func ConvirtiendoaNumero(texto string) (int, string) {
	numero, error := strconv.Atoi(texto)
	if error != nil {
		return 0, "HUBO UN ERROR"
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "No es mayor a 100"
	}
}
