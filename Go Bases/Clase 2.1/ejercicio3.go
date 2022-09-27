package main

import "fmt"

const (
	CategoryA = "A"
	CategoryB = "B"
	CategoryC = "C"
)

/*
3)
	Una empresa marinera necesita calcular el salario de sus empleados basándose
	en la cantidad de horas trabajadas por mes y la categoría.

	Si es categoría C, su salario es de $1.000 por hora
	Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
	Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

	Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados
	por mes y la categoría, y que devuelva su salario.
*/

func main() {
	fmt.Printf("El sueldo calculado es: %.1f\n", calculateSalary(300, CategoryC))
}

// Calcula el sueldo de una persona a partir de los minutos trabajados y su categoria
func calculateSalary(workedTieme int, category string) float64 {
	hoursWorked := float64(workedTieme) / 60
	hourValue := getHourlySalary(category)
	return hourValue(hoursWorked)
}

// Orquesta cual función que se encargara de calcular el sueldo segun su categoria
func getHourlySalary(category string) func(float64) float64 {
	switch category {
	case CategoryA:
		return categoryASalary
	case CategoryB:
		return categoryBSalary
	case CategoryC:
		return categoryCSalary
	}
	return nil

}

func categoryASalary(hoursWorked float64) (salary float64) {
	return hoursWorked * 1000.0
}
func categoryBSalary(hoursWorked float64) (salary float64) {
	return hoursWorked * 1500.0 * 1.2
}
func categoryCSalary(hoursWorked float64) (salary float64) {
	return hoursWorked * 3000.0 * 1.5
}
