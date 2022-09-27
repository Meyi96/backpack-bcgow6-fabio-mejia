package main

import (
	"errors"
	"fmt"
)

const (
	Minimum = "minimum"
	Average = "average"
	Maximum = "Maximum"
)

/*
4)
	Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de
	calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo, máximo
	y promedio de sus calificaciones.
	Se solicita generar una función que indique qué tipo de cálculo se quiere realizar
	(mínimo, máximo o promedio) y que devuelva otra función ( y un mensaje en caso que el cálculo
	no esté definido) que se le puede pasar una cantidad N de enteros y devuelva el cálculo que
	se indicó en la función anterior
*/

func main() {
	funcAverage, _ := operation(Average)
	funcMaximum, _ := operation(Maximum)
	funcMinimum, _ := operation(Minimum)
	_, err := operation("division")

	fmt.Printf("El promedio de calificacion fue: %.1f\n", funcAverage(3, 4, 5, 2, 3, 4, 4, 3))
	fmt.Printf("La calificacion mas alta fue: %.1f\n", funcMaximum(3, 4, 5, 2, 3, 4, 4, 3))
	fmt.Printf("La calificacion mas baja fue: %.1f\n", funcMinimum(3, 4, 5, 2, 3, 4, 4, 3))
	fmt.Println(err)

}

// Devuelve una funcion para realizar la operacion que se paso por parametro
// Si la operacion no existe lanza un error
func operation(operation string) (ope func(...int) float64, err error) {
	switch operation {
	case Maximum:
		ope = calculateMaximum
	case Minimum:
		ope = calculateMinimum
	case Average:
		ope = calculateAverage
	default:
		err = errors.New("No existe el operador.")
	}
	return

}

func calculateAverage(nums ...int) (average float64) {
	for _, num := range nums {
		average += float64(num)
	}
	return average / float64(len(nums))
}
func calculateMaximum(nums ...int) (maximum float64) {
	for i, num := range nums {
		if i == 0 {
			maximum = float64(num)
		}
		if float64(num) > maximum {
			maximum = float64(num)
		}
	}
	return
}
func calculateMinimum(nums ...int) (minimum float64) {
	for i, num := range nums {
		if i == 0 {
			minimum = float64(num)
		}
		if float64(num) < minimum {
			minimum = float64(num)
		}
	}
	return
}
