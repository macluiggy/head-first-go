// reporta si un estudiante ha aprobado o no dependiendo de su nota
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// "fmt"


func main()  {
	fmt.Print("Enter your grade: ")
	reader := bufio.NewReader(os.Stdin) // newreader es para que espere por input, y stdin es de donde se lee en este caso el teclado
	input, err := reader.ReadString('\n') // lee hasta que encuentra un salto de linea
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64) // convierte el string a float64
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passed"
	} else {
		status = "failed"
	}
	fmt.Println("A grado of", grade, "is", status)
	// println(input)
	// fmt.Println("!false")

}