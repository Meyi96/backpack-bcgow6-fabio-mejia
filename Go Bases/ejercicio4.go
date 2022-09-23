package main

import "fmt"

// var apellido string = "Gomez"	Bien
// var edad int = "35"				Error
// boolean := "false";				Error
// var sueldo string = 45857.90		Error
// var nombre string = "Julián"		Bien

func main() {
	fmt.Println("Hay 3 con errores sintaticos")
	var apellido string = "Gomez"
	var edad int = 35
	boolean := false
	var sueldo float64 = 45857.90
	var nombre string = "Julián"
	fmt.Println(nombre, apellido, edad, boolean, sueldo)
}
