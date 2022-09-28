package main

import (
	"errors"
	"fmt"
	"math"
)

/*
4)
Vamos a hacer que nuestro programa sea un poco más complejo.
A.	Desarrolla las funciones necesarias para permitir a la empresa calcular:
	a)	Salario mensual de un trabajador según la cantidad de horas trabajadas.
		-	La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
		-	Dicha función deberá retornar más de un valor (salario calculado y error).
		-	En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar
			el 10% en concepto de impuesto.
		-	En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo,
			la función debe devolver un error. El mismo deberá indicar “error: el trabajador no puede haber
			trabajado menos de 80 hs mensuales”.
	b)	Calcular el medio aguinaldo correspondiente al trabajador
		-	Fórmula de cálculo de aguinaldo: [mejor salario del semestre] / 12 * [meses trabajados en el semestre].
		-	La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que
			se ingrese un número negativo.
B.	Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”,
	“fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu función “main()”.
*/

func main() {
	valuePerHour := 1000.0
	setHoursWorked := []int{80, 100, 440, 120, 180}
	//Calculo de salarios
	salaries, err := calculateSalaries(valuePerHour, setHoursWorked)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Los salarios por mes: %v  cantidad de meses: %d\n", salaries, len(salaries))
	//Calculo de bonus
	bouns, err := calculateBonus(salaries)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("El aguinaldo es de: %.0f $\n", bouns)

}

func calculateSalary(valuePerHour float64, hoursWorked int) (salary float64, err error) {
	if hoursWorked < 80 {
		err = errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
		return
	}
	salary = float64(hoursWorked) * valuePerHour
	if salary > 150000.0 {
		salary *= .9
	}
	return
}

func calculateBonus(salaries []float64) (bonus float64, err error) {
	betterSalary := -math.MaxFloat64
	for _, salary := range salaries {
		if salary < 0 {
			err = errors.New("error: salario negativo")
			return
		}
		if betterSalary < salary {
			betterSalary = salary
		}

	}
	bonus = betterSalary / 12.0 * float64(len(salaries))
	return
}

func calculateSalaries(valuePerHour float64, setHoursWorked []int) (salaries []float64, err error) {
	if valuePerHour < 1 {
		err = fmt.Errorf("error: el valor de hora trabajada debe ser mayor a cero, valor: %.0f", valuePerHour)
		return
	}
	for _, hourWorked := range setHoursWorked {
		var temp float64
		temp, err = calculateSalary(valuePerHour, hourWorked)
		salaries = append(salaries, temp)
	}
	return
}
