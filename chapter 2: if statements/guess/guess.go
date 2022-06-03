// guess challenges players to guess a random number between 1 and 100.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	println(seconds)
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a number between 1 and 100.\nCan you guess it?")
	// fmt.Println(target, "aqui esta el target")
	// fmt.Printf("A integer: %d\n", target)

	reader := bufio.NewReader(os.Stdin)

	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You've got", 10-guesses, "guesses left.")

		fmt.Println("Make a guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess == target {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		} else if guess < target {
			fmt.Println("Too low. Guess again.")
			// main()
		} else if guess > target {
			fmt.Println("Too high. Guess again.")
			// main()
		}
	}

	if !success {
		fmt.Println("Sorry, you failed. The number I was thinking of was", target)
	}
}
