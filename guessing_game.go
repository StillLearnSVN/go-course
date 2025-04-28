package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number between 1 and 100
	randomNumber := random.Intn(100) + 1

	//
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have selected a random number between 1 and 100.")
	fmt.Println("Can you guess what it is?")
	fmt.Println("You have 10 attempts to guess the number.")
	fmt.Println("If you want to quit, type 'exit'.")
	fmt.Println("Good luck!")
	fmt.Println("")

	// Initialize variables
	var guess int
	var attempts int
	var userInput string
	// Loop until the user guesses the number or runs out of attempts
	for attempts < 10 {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&userInput)

		// Check if the user wants to exit
		if userInput == "exit" {
			fmt.Println("Thanks for playing! Goodbye!")
			return
		}

		// Convert user input to an integer
		fmt.Sscanf(userInput, "%d", &guess)

		attempts++

		if guess < randomNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > randomNumber {
			fmt.Println("Too high! Try again.")
		} else if guess == randomNumber {
			fmt.Printf("Congratulations! You guessed the number %d in %d attempts.\n", randomNumber, attempts)
			fmt.Println("Thanks for playing! Goodbye!")
			break
		} else {
			fmt.Println("Invalid input. Please enter a number between 1 and 100.")
		}
		
		// Check if the user has used all attempts
		// If the user has used all attempts, reveal the number and exit
		if attempts == 10 {
			fmt.Printf("Sorry, you've used all your attempts. The number was %d.\n", randomNumber)
			fmt.Println("Thanks for playing! Goodbye!")
			return
		}
	}
}
