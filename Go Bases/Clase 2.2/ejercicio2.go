package main

import (
	"fmt"
	"math"
)

/*
2)
	Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que
	represente una matriz de datos.
	Para ello requieren una estructura Matrix que tenga los métodos:
	-	Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
	-	Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
	La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del a
	ncho, si es cuadrática y cuál es el valor máximo.

*/

func main() {
	matrix := Matrix{Height: 3, Width: 4}
	matrix.setValues([]float64{2, 4, 2, 4, 2, 4, 6, 8, 5, 4, 65, 4})
	matrix.printMatrix()
	fmt.Printf("Matriz cuadratica: %v\n", matrix.isQuadratic())
	fmt.Printf("El valor maximo es: %.0f\n", matrix.maxValueMatrix())

}

type Matrix struct {
	Values []float64
	Height int
	Width  int
}

func (matrix *Matrix) setValues(values []float64) {
	matrix.Values = values
}
func (matrix Matrix) printMatrix() {
	for i := 0; i < matrix.Height; i++ {
		fmt.Printf("%.0f\n", matrix.Values[i*matrix.Width:(i*matrix.Width)+matrix.Width])
	}
}

func (matrix Matrix) isQuadratic() bool {
	return matrix.Height == matrix.Width && matrix.Width != 0
}

func (matrix Matrix) maxValueMatrix() float64 {
	max := -math.MaxFloat64
	for _, value := range matrix.Values {
		if value > max {
			max = value
		}
	}
	return max
}
