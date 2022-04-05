// reporta si un estudiante ha aprobado o no dependiendo de su nota
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// "fmt"


func main()  {
	fmt.Print("Enter your grade: ")
	reader := bufio.NewReader(os.Stdin) // newreader es para que espere por input, y stdin es de donde se lee en este caso el teclado
	input, err := reader.ReadString('\n') // lee hasta que encuentra un salto de linea
	if err != nil {
		log.Fatal(err)
	}
	println(input)
	// fmt.Println("!false")

}