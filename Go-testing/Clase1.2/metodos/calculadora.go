package metodos

import "errors"

var NotZeroDivision = errors.New("El denominador no puede ser 0")

// Función que recibe dos enteros y retorna la suma resultante
func Add(num1, num2 int) int {
	// Esta función ahora devuelve un resultado INCORRECTO
	return num1 + num2
}

// Función que recibe dos enteros y retorna la resta o diferencia resultante
func Subtract(num1, num2 int) int {
	return num1 - num2
}

func division(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, NotZeroDivision
	}
	return num1 / num2, nil
}
