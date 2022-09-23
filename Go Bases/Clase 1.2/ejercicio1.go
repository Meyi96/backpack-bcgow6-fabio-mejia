package main

import "fmt"

func main() {
	name := "Fabio"
	fmt.Printf("Longitud del mensaje: %d \n", len(name))
	for _, char := range name {
		fmt.Println(string(char))
	}
}
