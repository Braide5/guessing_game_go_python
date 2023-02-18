package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to the amazing game!")

	// Declaring the random_number and assigning it a random number
	source := rand.NewSource(time.Now().UnixNano())
	random_generator := rand.New(source)

	random_number := random_generator.Intn(100)
	fmt.Println("Enter a number between 0 and 100")

	for {

		// Declaring the guess variable to hold an integer
		var guess int

		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}

		if guess < random_number {
			fmt.Println("you guessed too low, please try again")
		} else if guess > random_number {
			fmt.Println("you guessed too high, please try again")
		} else {
			fmt.Println("you guessed correctly")
			break
		}
	}
}
