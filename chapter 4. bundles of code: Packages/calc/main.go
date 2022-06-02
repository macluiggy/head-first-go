package main

import (
	"fmt"
	"head/src/calc"
	"log"
)

func main() {
	add, err := calc.Add(1, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(add)

	sub, err2 := calc.Substract(1, 2)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(sub)
}
