package main

import (
	"errors"
	"fmt"
)

/*
2)
	Un colegio necesita calcular el promedio (por alumno) de sus calificaciones.
	Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros
	y devuelva el promedio y un error en caso que uno de los números ingresados sea negativo
*/

func main() {

	average, err := calculateGradeAverage(4.0, 5.0, 3.0) //Promedio: 4.0
	// average, err := calculateGradeAverage(4.0, 5.0, 3.0, 5.0, 4.5) //Promedio: 4.3
	// average, err := calculateGradeAverage(4.0, 5.0, -3.0, 5.0, 4.5) //Error por calificación negativa
	if err == nil {
		fmt.Printf("El promedio de las notas es: %.1f \n", average)
	} else {
		fmt.Println(err)
	}

}

func calculateGradeAverage(grades ...float64) (average float64, err error) {
	for _, grade := range grades {
		if grade < 0 {
			average, err = 0, errors.New("No pueden existir notas negativas.")
			break
		}
		average += grade
	}
	return average / float64(len(grades)), nil
}
