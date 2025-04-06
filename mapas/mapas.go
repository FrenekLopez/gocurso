package mapas

import (
	"fmt"
)

func MuestrarMapas() {
	paises := make(map[string]string)
	//fmt.Println(paises)

	paises["Mexico"] = "D.F"
	paises["Argentina"] = "Buenos Aires"
	//fmt.Println(paises)
	//fmt.Println(paises["Argentina"])

	campionato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"Chivas":      37,
		"Boca Junior": 19,
	}
	fmt.Println(campionato)

	//for equipo, puntaje := range campionato {
	//	fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	//}
	delete(campionato, "Real Madrid")
	fmt.Println(campionato)

	puntaje, existe := campionato["Chivas"]
	fmt.Printf("El porcentaje capturado es %d, y el equipo existe = %t, \n", puntaje, existe)

}
