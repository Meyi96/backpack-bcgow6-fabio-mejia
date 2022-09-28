package main

import (
	"fmt"
)

/*
1)
	En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
	Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el
	salario ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor
	a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/

func main() {
	var salary int = 15000
	notValidMinimumSalary := NotValidMinimumSalary{message: "error: el salario ingresado no alcanza el mínimo imponible"}

	if salary < 150000 {
		fmt.Println(&notValidMinimumSalary)
		return
	}
	fmt.Println("Debe pagar impuesto")

}

type NotValidMinimumSalary struct {
	message string
}

func (e *NotValidMinimumSalary) Error() string {
	return fmt.Sprint(e.message)
}
