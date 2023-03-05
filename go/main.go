package main

import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "welcome!")
}

func main() {
		http.HandleFunc("/", handler)
		http.ListenAndServe(":8080", nil)
	}

func main() {
	fmt.Println("welcome to the amazing game!!")

	// Creating a random number source using the current time
	source := rand.NewSource(time.Now().UnixNano())

	// Creating a random number generator using the random number source
	random_generator := rand.New(source)

	// Getting a random number between 0 and 100
	random_number := random_generator.Intn(100)
	fmt.Println("Enter a number between 0 and 100!")

	for {

		// Declaring the guess variable to hold an integer
		var guess int

		// Reading from the console and storing the value in the guess variable
		_, err := fmt.Scanln(&guess)

		// If the user enters a non-integer value, we print an error message
		// and continue
		if err != nil {
			fmt.Println("Please enter a valid number")
			continue
		}

		// If the user enters a number less than 0 or greater than 100,
		// we print an error message and continue
		if guess < random_number {
			fmt.Println("you guessed too low, please try again!")
		} else if guess > random_number {S
			fmt.Println("you guessed too high, please try again!")
		} else {
			fmt.Println("you guessed correctly")
			break
		}
	}
}
