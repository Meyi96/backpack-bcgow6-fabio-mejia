package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

type IdGenerator struct {
	current int
	limit   int
}
type Customer struct {
	Id      int    `csv:"Legajo"`
	Name    string `csv:"Nombre"`
	Dni     int    `csv:"Dni"`
	Phone   string `csv:"Telefono"`
	Address string `csv:"Domicilio"`
}

type FileNotRead struct {
	msg string
}

func (e FileNotRead) Error() string {
	return fmt.Sprint(e.msg)
}

func (idGe IdGenerator) nexId() (id int, err error) {
	if idGe.current+1 > idGe.limit {
		err = errors.New("error: se excede el limite de ids a generar")
	}
	id = idGe.current + 1
	return
}

func isCustomer(dni int, path string) (err error) {
	customerFile, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		err = errors.New("el archivo indicado no fue encontrado o está dañado")
	}
	defer customerFile.Close()
	customers := []*Customer{}
	if err := gocsv.UnmarshalFile(customerFile, &customers); err != nil {
		err = errors.New("hubo un problema con el mapeo de los registros a la estructura")
	}
	for _, customer := range customers {
		if dni == customer.Dni {
			return fmt.Errorf("error: el dni existe y es del cliente: %s", customer.Name)
		}
	}
	return
}
func validateFields(id int, name string, dni int, phone string, adddress string) (customer Customer, err error) {
	if (name == "") || (dni == 0) || (phone == "") || (adddress == "") {
		err = errors.New("Los campos no puede ser nulos")
		return
	}
	customer = Customer{Id: id, Name: name, Dni: dni, Phone: phone, Address: adddress}
	return
}
func printRecover() {
	err := recover()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Termino el programa")
}
func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	path := "./customers.csv"
	idGenerator := IdGenerator{current: 3, limit: 10}

	dni := 342
	name := "Carla"
	phone := "3283479403"
	address := "calle 12 # 04"

	id, err := idGenerator.nexId()
	errPanic(err)

	defer printRecover()
	err = isCustomer(dni, path)
	errPanic(err)
	customer, err := validateFields(id, name, dni, phone, address)
	errPanic(err)
	fmt.Printf("El usuario %s con dni %d se puede agregar con id %d\n", customer.Name, customer.Dni, customer.Id)
}
