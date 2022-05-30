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
	fmt.Println(target, "aqui esta el target")

	reader := bufio.NewReader(os.Stdin)

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
			fmt.Println("You guessed it!")
			break
		} else if guess < target {
			fmt.Println("Too low. Guess again.")
			// main()
		} else if guess > target {
			fmt.Println("Too high. Guess again.")
			// main()
		}
	}
}
