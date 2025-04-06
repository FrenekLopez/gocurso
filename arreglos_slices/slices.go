package arreglos_slices

import "fmt"

var tablaS []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 98, 21, 674, 18, 36, 78, 9}

func MuestroSlices() {
	fmt.Println(tablaS)
	porcion := arreglo[3:]   //Slice cread desde un vector, desde la posicion 3
	porcion2 := arreglo[:5]  // Slice cread desde un vector, desde la posicion 0 asta la 5
	porcion3 := arreglo[6:7] //Slice cread desde un vector, desde la posicion 6 asta la 7
	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("largo %d capacidad %d", len(elementos), cap(elementos))
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nlargo %d capacidad %d ", len(nums), cap(nums))

}
