package main

import "fmt"

var total float64

func main() {
	paintCalculator(10, 20)
	paintCalculator(4.5, 3.5)
	// functionD("Hello", 5)
	fmt.Printf("The total is %.2f\n", total)
}

func paintCalculator(width float64, height float64) float64 {
	area := width * height
	amount := area / 10.0
	total += amount
	fmt.Printf("%.2f liters of paint required\n", amount)
	fmt.Printf(("The current total is %.2f\n"), total)
	return amount
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
