package main

import (
	"errors"
	"fmt"
)

/*
2)
	Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo
	de “Error()”,  se implemente “errors.New()”.
*/

func main() {
	var salary int = 15000

	if salary < 150000 {
		fmt.Println(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
		return
	}
	fmt.Println("Debe pagar impuesto")

}
