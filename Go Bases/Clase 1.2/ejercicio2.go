package main

import "fmt"

func main() {
	age := 32
	lengthService := 2
	salary := 500000.0
	calculateLoan(age, lengthService, salary)

}

func calculateLoan(age int, lengthService int, salary float64) {
	if (age <= 22) || (lengthService <= 1) {
		fmt.Println("Lo sentimos usted no ha sido aprobado para el presamo.")
	} else if salary > 100000 {
		fmt.Println("Felicitaciones usted ha sido aprobado para el presamo sin cobro de intereses.")
	} else {
		fmt.Println("Felicitaciones usted ha sido aprobado para el presamo con cobro de intereses.")
	}
}
