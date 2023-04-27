// Created by: Mariam Kasim
// Created on: Apr 2023
//
// This program generates a random number between 1 and 6 and asks the user to guess the number.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// This function is the main function
	// variables
	var guess int
	var answer int

	// input
	fmt.Println("Enter a number between 1 and 6: ")
	fmt.Scanln(&guess)

	// process
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	answer = rand.Intn(max-min) + min

	// output
	if guess == answer {
		fmt.Println("You guessed correctly!")
	}
	if guess != answer {
		fmt.Println("You guessed wrong! The answer was", answer)
	}
	fmt.Println("\nDone.")
}
