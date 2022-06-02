package main

import (
	"fmt"
	"log"
)

// local imports
import (
	// "head/src/github.com/macluiggy/keyboard"
	// "keyboard" // si tienes go.mod no va a funcionar de esta forma
	"github.com/headfirstgo/keyboard"
)

func main() {
	fmt.Print("Enter a temperature in Celsius: ")
	celsius, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The temperature in Fahrenheit is", celsius*1.8+32)
}
