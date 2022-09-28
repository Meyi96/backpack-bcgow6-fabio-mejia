package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Info struct {
	Numbers []int
	Timer   time.Duration
	Method  string
}

func main() {
	executeSortTest(100)
	executeSortTest(1000)
	executeSortTest(10000)
}

func bubbleSort(numbers []int, c chan Info) {
	tm := time.Now()
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
				temp := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = temp
			}
		}
	}
	info := &Info{Numbers: numbers, Timer: time.Since(tm), Method: "Burbuja"}
	c <- *info
}

func insercionSort(numbers []int, c chan Info) {
	tm := time.Now()
	var j, aux int
	for i := 1; i < len(numbers); i++ {
		aux = numbers[i]
		j = i - 1
		for (j >= 0) && (aux < numbers[j]) {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = aux
	}
	info := &Info{Numbers: numbers, Timer: time.Since(tm), Method: "Insercion"}
	c <- *info
}

func selectionSort(numbers []int, c chan Info) {
	tm := time.Now()
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] > numbers[j] {
				temp := numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}
	info := &Info{Numbers: numbers, Timer: time.Since(tm), Method: "Selecion"}
	c <- *info
}

func executeSortTest(size int) {
	c := make(chan Info)
	numbers := rand.Perm(size)
	go selectionSort(numbers, c)
	go bubbleSort(numbers, c)
	go insercionSort(numbers, c)
	for i := 0; i < 3; i++ {
		info := <-c
		fmt.Printf("El metodo %-9s: se demora %-8v microsegundos con un arreglo de tamaño %d\n", info.Method, info.Timer.Nanoseconds(), size)
	}
	fmt.Println()
}
