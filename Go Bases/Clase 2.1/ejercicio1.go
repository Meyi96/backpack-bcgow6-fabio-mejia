package main

import (
	"errors"
	"fmt"
)

/*
1)
	Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento
	de depositar el sueldo, para cumplir el objetivo es necesario crear una función que
	devuelva el impuesto de un salario.
	Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del
	sueldo y si gana más de $150.000 se le descontará además un 10%.
*/

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// input := scanner.Text()
	// salary, _ := strconv.ParseFloat(input, 8)
	salary := 160000.0
	// salary := 16000.0
	// salary := 1600.0
	// salary := -1600.0
	tax, err := calculatePayrollTax(salary)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\nEl impuesto para un salario de %.1f es de : %.1f\n", salary, tax)
	}
}

func calculatePayrollTax(salary float64) (tax float64, err error) {
	if salary < 1 {
		err = errors.New("El salario debe ser mayor a cero.")
	}
	switch {
	case 50000 < salary && salary <= 150000:
		tax = salary * .17
	case salary > 150000:
		tax = salary * .27
	default:
		tax = 0
	}
	return
}
