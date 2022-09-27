package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
2)
	La misma empresa necesita leer el archivo almacenado, para ello requiere que:
	se imprima por pantalla mostrando los valores tabulados, con un título
	(tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad),
	el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando
	precio por cantidad)
*/

func main() {
	data, err := os.ReadFile("./products.txt")

	if err != nil {
		panic("Ocurrio un problema con la lectura del archivo.")
	}

	registers := strings.Split(string(data), "\n")
	totalPrice := 0.0
	//Imprimir encabezado
	printRegister(registers[0], true)
	for _, register := range registers[1:] {
		if register != "" {
			//Imprimir registros y guardar el precio para sumarlo
			currentPrice := printRegister(register, false)
			totalPrice += currentPrice
		}
	}
	//Imprimir total
	fmt.Printf(" \t\t%10.2f\t\t \n", totalPrice)
}

func printRegister(register string, header bool) (price float64) {
	v := strings.Split(register, ",")
	fmt.Printf("%s\t\t%10s\t\t%10s\n", v[0], v[1], v[2])
	if !header {
		price, err := strconv.ParseFloat(v[1], 64)
		if err != nil {
			panic("No se puede convertir string como flotante.")
		}
		return price
	}
	return
}
