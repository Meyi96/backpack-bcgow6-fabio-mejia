package main

import "fmt"

// var 1nombre string				Error
// var apellido string				Bien
// var int edad						Error
// 1apellido := 6					Error
// var licencia_de_conducir = true	Error
// var estatura de la persona int	Error
// cantidadDeHijos := 2				Bien

func main() {
	fmt.Println("Hay 5 con errores sintaticos")
	var nombre string
	var apellido string
	var edad int
	apellido1 := "Mejia"
	var licenciaConducion = true
	var estaturaPersona int
	cantidadDeHijos := 2
	fmt.Println(nombre, apellido, edad, apellido1, licenciaConducion, estaturaPersona, cantidadDeHijos)
}
