package main

import (
	"fmt"
	"os"
)

/*
1)
	Una empresa que se encarga de vender productos de limpieza necesita:
	- Implementar una funcionalidad para guardar un archivo de texto, con la
	información de productos comprados, separados por punto y coma (csv).
	- Debe tener el id del producto, precio y la cantidad.
	- Estos valores pueden ser hardcodeados o escritos en duro en una variable.

*/

func main() {
	products := productsInitializer()

	data := "Id,Precio,Cantidad\n"
	for _, product := range products {
		data += fmt.Sprintf("%s,%.2f,%d\n", product.Id, product.Price, product.Quantity)
	}
	err := os.WriteFile("./products.txt", []byte(data), 0644)

	if err != nil {
		panic("Ocurrio un problema con la escritura del archivo.")
	}
}

type Product struct {
	Id       string
	Price    float64
	Quantity int
}

func productsInitializer() []Product {
	p1 := Product{Id: "111223", Price: 30012.00, Quantity: 1}
	p2 := Product{Id: "444321", Price: 1000000.00, Quantity: 4}
	p3 := Product{Id: "434321", Price: 50.50, Quantity: 1}
	return []Product{p1, p2, p3}
}
