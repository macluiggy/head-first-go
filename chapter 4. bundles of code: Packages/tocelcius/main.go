package main

import (
	"fmt"
	"log"
)

// local imports
import (
	"head/src/keyboard"
)

func main() {
	fmt.Print("Enter a temperature in Celsius: ")
	celsius, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The temperature in Fahrenheit is", celsius*1.8+32)
}
