package main

import (
	"errors"
	"fmt"
)

const (
	Dog       = "dog"
	Cat       = "cat"
	Hamster   = "hamster"
	Tarantula = "tarantula"
)

/*
4)
	Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
	Por el momento solo tienen tarántulas, hamsters, perros, y gatos, pero se espera que puedan
	haber muchos más animales que refugiar.

	perro necesitan 10 kg de alimento
	gato 5 kg
	Hamster 250 gramos.
	Tarántula 150 gramos.

	Se solicita:
	Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal
	especificado y que retorne una función y un mensaje (en caso que no exista el animal)
	Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del
	tipo de animal especificado.
*/

func main() {
	dogAnimal, _ := animal(Dog)
	catAnimal, _ := animal(Cat)
	hamsterAnimal, _ := animal(Hamster)
	tarantulaAnimal, _ := animal(Tarantula)
	_, err := animal("pez")

	var amount float64
	amount += dogAnimal(5)
	amount += catAnimal(4)
	amount += hamsterAnimal(8)
	amount += tarantulaAnimal(3)
	fmt.Printf("Catindad de comida por todos los animales es:  %.2f kg\n", amount)
	fmt.Println(err)

}

// Devuelve una funcion para realizar el calculo de la comida una especie
// segun el parametro de animal
func animal(animal string) (ope func(int) float64, err error) {
	switch animal {
	case Dog:
		ope = quantityDogFood
	case Cat:
		ope = quantityCatFood
	case Hamster:
		ope = quantityHamsterFood
	case Tarantula:
		ope = quantitytarantulaFood
	default:
		err = errors.New("No existe el animal.")
	}
	return

}

func quantityDogFood(number int) float64 {
	return float64(number) * 10.0
}
func quantityCatFood(number int) float64 {
	return float64(number) * 5.0
}
func quantityHamsterFood(number int) float64 {
	return float64(number) * .25
}
func quantitytarantulaFood(number int) float64 {
	return float64(number) * .15
}
