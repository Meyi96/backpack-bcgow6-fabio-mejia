package main

import "fmt"

/*
2)

	Una importante empresa de ventas web necesita agregar una funcionalidad para agregar
	productos a los usuarios. Para ello requieren que tanto los usuarios como los
	productos tengan la misma dirección de memoria en el main del programa como en las funciones.
	Se necesitan las estructuras:
	-	Usuario: Nombre, Apellido, Correo, Productos (array de productos).
	-	Producto: Nombre, precio, cantidad.
	Se requieren las funciones:
	-	Nuevo producto: recibe nombre y precio, y retorna un producto.
	-	Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el
		producto al usuario.
	-	Borrar productos: recibe un usuario, borra los productos del usuario.
*/
func main() {
	p1 := newProduct("Gafas", 2000.0)
	p2 := newProduct("Gorra", 35000.0)
	u1 := User{Name: "Fabio", LastName: "Mejia", Mail: "algo@algo.com"}

	fmt.Printf("Producto: %v\t Direccion de memoria:%p\n", p1, &p1)
	fmt.Printf("Producto: %v\t Direccion de memoria:%p\n", p2, &p2)

	addProductToUser(&u1, &p1, 3)
	addProductToUser(&u1, &p2, 5)

	fmt.Printf("Productos del usuario: %v\n", u1.Products)
	deleteProductsToUser(&u1)
	fmt.Printf("Productos del usuario: %v\n", u1.Products)
}

type User struct {
	Name     string
	LastName string
	Mail     string
	Products []Product
}

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func newProduct(name string, price float64) Product {
	return Product{Name: name, Price: price}
}
func deleteProductsToUser(user *User) {
	user.Products = nil
}

func addProductToUser(u *User, product *Product, cantidad int) {
	product.Quantity = cantidad
	u.Products = append(u.Products, *product)
}
