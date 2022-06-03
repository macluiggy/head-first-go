package main

import "fmt"

func main() {
	amount := 6
	double(&amount)
	fmt.Println(amount)
}

func double(x *int) {
	*x *= 2
	// return x
}
