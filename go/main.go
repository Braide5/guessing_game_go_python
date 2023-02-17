package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("welcome to the amazing game!")
	random_number := rand.Intn(100)

	for {
		fmt.Println("Enter a number between 0 and 100")
		var guess int
		if _, err := fmt.Scanln("%d", &guess); err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}

		if guess < random_number {
			fmt.Println("you guessed too low")
		} else if guess > random_number {
			fmt.Println("you guessed too high")
		} else {
			fmt.Println("you guessed correctly")
		}
		break
	}
}
