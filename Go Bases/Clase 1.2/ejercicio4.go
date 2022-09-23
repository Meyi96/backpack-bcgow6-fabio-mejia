package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	employeeName := "Benjamin"
	//info employee
	employeeInfo(employeeName)
	//Number of employees over 21
	countEmployeesCondition(21)
	//add federico as an employee
	addUser("Federico", 25)
	//Delete Pedro
	deleteUserByName("Pedro")
}

func employeeInfo(name string) {
	fmt.Printf("1.\nNombre empleado: %s \nEdad: %d \n", name, employees[name])
}

func countEmployeesCondition(umbral int) {
	var count int
	for _, element := range employees {
		if element > umbral {
			count++
		}
	}
	fmt.Printf("2. \nEmpleados mayores de %d años: %d \n", umbral, count)
}

func addUser(name string, age int) {
	fmt.Printf("3. \nCantidad de empleados antes de Agregar usuario %d \n", len(employees))
	if _, ok := employees[name]; !ok {
		employees[name] = age
		fmt.Printf("Cantidad de empleados despues de Agregar usuario %d \n", len(employees))
	} else {
		fmt.Println("La llave ya existe")
	}
}

func deleteUserByName(name string) {
	fmt.Printf("4. \nCantidad de empleados antes de borrar usuario %d \n", len(employees))
	if _, ok := employees[name]; ok {
		delete(employees, name)
		fmt.Printf("Cantidad de empleados despues de borrar usuario %d \n", len(employees))
	} else {
		fmt.Println("Usuario no encontrado")
	}
}
