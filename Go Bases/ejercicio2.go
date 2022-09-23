package main

import "fmt"

var temperature int64 = 22
var atmoPressure float64 = 1015
var humidity float64 = 63.2

func main() {
	fmt.Printf("Temperatura: %[1]d C° Tipo: %[1]T\nPresión: %.1[2]f hPa Tipo: %[2]T\nHumedad: %.1[3]f %% Tipo: %[3]T\n",
		temperature, atmoPressure, humidity)
}
