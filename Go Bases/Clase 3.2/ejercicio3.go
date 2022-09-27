package main

import "fmt"

/*
3)

	Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
	Para ello requieren realizar un programa que se encargue de calcular el precio total de
	Productos, Servicios y Mantenimientos. Debido a la fuerte demanda y para optimizar la velocidad
	requieren que el cálculo de la sumatoria se realice en paralelo mediante 3 go routines.

	Se requieren 3 estructuras:
	-	Productos: nombre, precio, cantidad.
	-	Servicios: nombre, precio, minutos trabajados.
	-	Mantenimiento: nombre, precio.

	Se requieren 3 funciones:
	-	Sumar Productos: recibe un array de producto y devuelve el precio total (precio * cantidad).
	-	Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio * media hora trabajada,
		si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
	-	Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

	Los 3 se deben ejecutar concurrentemente y al final se debe mostrar por pantalla el monto final (sumando el total de los 3).
*/
func main() {
	s1 := Service{Name: "s1", Price: 10.0, minutesWorked: 120}
	s2 := Service{Name: "s2", Price: 230.0, minutesWorked: 10}
	s3 := Service{Name: "s3", Price: 15.0, minutesWorked: 25}
	services := []Service{s1, s2, s3}

	p1 := Product{Name: "p1", Price: 300.0, Quantity: 5}
	p2 := Product{Name: "p2", Price: 450.0, Quantity: 2}
	p3 := Product{Name: "p3", Price: 8.0, Quantity: 32}
	products := []Product{p1, p2, p3}

	m1 := Maintenance{Name: "m1", Price: 245.0}
	m2 := Maintenance{Name: "m2", Price: 510.0}
	maintenances := []Maintenance{m1, m2}

	c := make(chan float64)
	var totalPrice float64
	go calculatePriceProducts(products, c)
	go calculatePriceService(services, c)
	go calculatePriceMaintenance(maintenances, c)

	for i := 0; i < 3; i++ {
		totalPrice += <-c
	}
	fmt.Printf("El precio total de productos, servicios y mantenimientos es: %.1f\n", totalPrice)
}

type Service struct {
	Name          string
	Price         float64
	minutesWorked int
}

type Product struct {
	Name     string
	Price    float64
	Quantity int
}
type Maintenance struct {
	Name  string
	Price float64
}

func calculatePriceProducts(products []Product, c chan float64) {
	var totalPrice float64
	for _, product := range products {
		totalPrice += product.Price * float64(product.Quantity)
	}
	c <- totalPrice
}
func calculatePriceService(services []Service, c chan float64) {
	var totalPrice float64
	for _, service := range services {
		totalPrice += service.Price * float64(service.minutesWorked)
	}
	c <- totalPrice
}
func calculatePriceMaintenance(maintens []Maintenance, c chan float64) {
	var totalPrice float64
	for _, mainten := range maintens {
		totalPrice += mainten.Price
	}
	c <- totalPrice
}
