package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess the number!")
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100) + 1

	for {
		fmt.Println("Please input your guess.")

		var guessedNumber int
		_, err := fmt.Scanf("%d", &guessedNumber)
		if err != nil {
			continue
		}
		fmt.Println("You guessed:", guessedNumber)

		if guessedNumber == secretNumber {
			fmt.Println("You win!")
			return
		} else if guessedNumber < secretNumber {
			fmt.Println("To small!")
		} else if guessedNumber > secretNumber {
			fmt.Println("To big!")
		}
	}
}
