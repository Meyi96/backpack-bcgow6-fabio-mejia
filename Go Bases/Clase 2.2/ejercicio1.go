package main

import "fmt"

/*
1)
	Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el
	detalle de los datos de cada uno de ellos/as, de la siguiente manera:
		Nombre: [Nombre del alumno]
		Apellido: [Apellido del alumno]
		DNI: [DNI del alumno]
		Fecha: [Fecha ingreso alumno]
	Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
	Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI,
	Fecha y que tenga un método detalle.
*/

func main() {
	student := Student{Name: "Fabio", LastName: "Mejia", Dni: "1113696432", Date: "19/09/22"}
	student.studentDetail()
}

type Student struct {
	Name     string
	LastName string
	Dni      string
	Date     string
}

func (s Student) studentDetail() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %s\nFechaa: %s\n", s.Name, s.LastName, s.Dni, s.Date)
}
