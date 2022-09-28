package main

import "fmt"

/*
3)
Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y
retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
-	Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
-	Pequeño: El costo del producto (sin costo adicional)
-	Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
-	Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.
Requerimientos:
-	Crear una estructura “tienda” que guarde una lista de productos.
-	Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
-	Crear una interface “Producto” que tenga el método “CalcularCosto”
-	Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
-	Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
-	Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
-	Interface Producto:
	-	El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
	Interface Ecommerce:
	-	El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
	-	El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda

*/

const (
	smallProduct  = "small"
	mediumProduct = "medium"
	largeProduct  = "large"
)

type product interface {
	calculateCost() float64
}
type ecommerce interface {
	total() float64
	addProduct(p product)
}
type Product struct {
	name        string
	price       float64
	typeProduct string
}
type SmallProduct struct {
	product Product
}
type MediumProduct struct {
	product Product
}
type LargeProduct struct {
	product Product
}
type Store struct {
	products []product
}

func (s SmallProduct) calculateCost() float64 {
	return s.product.price
}
func (m MediumProduct) calculateCost() float64 {
	return m.product.price * 1.03
}
func (l LargeProduct) calculateCost() float64 {
	return (l.product.price * 1.06) + 2500
}
func (s *Store) addProduct(p product) {
	s.products = append(s.products, p)
}
func (s Store) total() (total float64) {
	for _, product := range s.products {
		total += product.calculateCost()
	}
	return
}
func newProduct(name string, typeProduct string, price float64) product {
	switch typeProduct {
	case smallProduct:
		p := Product{name: name, price: price, typeProduct: typeProduct}
		return SmallProduct{product: p}
	case mediumProduct:
		p := Product{name: name, price: price, typeProduct: typeProduct}
		return MediumProduct{product: p}
	case largeProduct:
		p := Product{name: name, price: price, typeProduct: typeProduct}
		return LargeProduct{product: p}
	}
	return nil
}
func newStore() ecommerce {
	return &Store{}
}
func main() {
	meli := newStore()
	meli.addProduct(newProduct("gafas", smallProduct, 15000.0))
	meli.addProduct(newProduct("Ventilador", mediumProduct, 135000.0))
	meli.addProduct(newProduct("Nevera", largeProduct, 1250000.0))
	meli.addProduct(newProduct("Libreta", smallProduct, 23000.0))
	fmt.Printf("El costo total es de: %.0f$\n", meli.total())
}
