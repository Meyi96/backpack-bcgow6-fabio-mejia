package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./customerss.txt"
	defer fmt.Println("ejecución finalizada")
	data := loadCustomers(path)
	fmt.Println(string(data))

}

func loadCustomers(path string) (data []byte) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	return
}
