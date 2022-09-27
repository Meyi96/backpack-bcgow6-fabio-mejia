package main

import "fmt"

type User struct {
	name     string
	lastName string
	age      int
	mail     string
	password string
}

/*
1)
	Una empresa de redes sociales requiere implementar una estructura usuario con
	funciones que vayan agregando información a la estructura. Para optimizar y ahorrar
	memoria requieren que la estructura de usuarios ocupe el mismo lugar en memoria para
	el main del programa y para las funciones.
	La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
	Y deben implementarse las funciones:
	-	Cambiar nombre: me permite cambiar el nombre y apellido.
	-	Cambiar edad: me permite cambiar la edad.
	-	Cambiar correo: me permite cambiar el correo.
	-	Cambiar contraseña: me permite cambiar la contraseña.

*/

func main() {
	user := User{name: "Andres", lastName: "Mejia", age: 23, mail: "exam@exam.com", password: "12345"}
	fmt.Printf("Usuario: %v\t Direccion de memoria:%p\n", user, &user)
	user.setFullName("Carlos", "Lopez")
	user.setAge(34)
	user.setMail("test@test.com")
	user.setPassword("54321")
	fmt.Printf("Usuario: %v\t Direccion de memoria:%p\n", user, &user)

}

func (user *User) setFullName(name string, lastName string) {
	user.name = name
	user.lastName = lastName
}
func (user *User) setAge(age int) {
	user.age = age
}
func (user *User) setMail(mail string) {
	user.mail = mail
}
func (user *User) setPassword(password string) {
	user.password = password
}
