package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%f is the amount of money you have\n", 100.81999999999)

	// fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	// fmt.Printf("%12s | %s\n", "-------", "--------------")
	// fmt.Printf("%12s | %2d\n", "Laptop", 10)
	// fmt.Printf("%12s | %2d\n", "Phone", 5)
	// fmt.Printf("%12s | %2d\n", "TV", 3000)

	fmt.Printf(("%%7.3f: %7.3f\n"), 1.23456789)
	fmt.Printf(("%%7.3f: %7.3f\n"), 12222222.23456789)
	fmt.Printf(("%%7.3f: %7.3f\n"), 123.23456789)
	fmt.Printf(("%%7.3f: %7.3f\n"), 12.23456789)
	sayHi()
}
