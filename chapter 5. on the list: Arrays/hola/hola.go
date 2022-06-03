package main

import "fmt"

func main() {
	var primes [3]int
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5

	var a string
	fmt.Println("fgg", a)

	primes = [3]int{2, 3, 5}
	fmt.Println(primes[2])
	// println(primes)
	fmt.Printf("%T\n", primes)
	fmt.Printf("%#v\n", primes)

	var notes = [3]string{"do", "re", "mi"}
	// for i := 0; i < len(notes); i++ {
	// fmt.Println(i, notes[i])
	// }

	for _, note := range notes {
		// fmt.Println(note, i, notes[i])
		fmt.Println(note)

	}
	// for i, _ := range notes {
	// 	// fmt.Println(note, i, notes[i])
	// 	fmt.Println(i)

	// }
}
