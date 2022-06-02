package main

import (
	"fmt"
	"log"
)

var total float64

func main() {
	// var a error = nil
	// println(a)
	// amount, err := paintCalculator(10, 20)
	// println(err)
	// fmt.Printf("%.2f liters of paint required\n", amount)
	amount, err2 := paintCalculator(4.5, -3.5)
	fmt.Println(err2) // con formato
	// print(err2, "err2")   // sin formato
	// println(err2, "err2") // sin formato
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("%.2f liters of paint required\n", amount)

	amount, e := paintCalculator(4.5, 3.5)
	fmt.Println(e)
	fmt.Printf("%.2f liters of paint required\n", amount)
	// functionD("Hello", 5)
	fmt.Printf("The total is %.2f\n", total)
}

func paintCalculator(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("width of %.2f is less than zero", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("height of %.2f is less than zero", height)
	}
	area := width * height
	amount := area / 10.0
	total += amount
	// fmt.Printf("%.2f liters of paint required\n", amount)
	fmt.Printf(("The current total is %.2f\n"), total)
	return amount, nil
	// println("no voy a imprimir nada")
}

func functionD(a string, b int) /*string*/ {
	// if false {
	// 	return "Hello"
	// }

	for i := 0; i < b; i++ {
		print(a, "\n")
	}
	// return "fhfh"
}
