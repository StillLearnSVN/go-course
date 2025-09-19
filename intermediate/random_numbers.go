package main

import (
	"math/rand"
	"fmt"
)

func main() {
	// fmt.Println(rand.Intn(101))

	for {
		// Show the menu
		fmt.Println("Welcome to the Dice Game!")
		fmt.Println("1. Roll the dice!")
		fmt.Println("2. Exit")
		fmt.Print("Choose an option(1/2): ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		if choice == 2 {
			fmt.Println("Exiting the game. Goodbye!")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1
		
		fmt.Printf("You rolled a %d and a %d.\n", die1, die2)
		fmt.Printf("Total: %d\n", die1 + die2)

		// Ask if the user wants to play again
		fmt.Print("Do you want to roll again? (y/n): ")
		var again string
		_, err = fmt.Scan(&again)
		if err != nil || (again != "y" && again != "n") {
			fmt.Println("Invalid input. assumming no.")
			break
		}
		if again == "n" {
			fmt.Println("Thanks for playing! Goodbye!")
			break
		}

	}
}
