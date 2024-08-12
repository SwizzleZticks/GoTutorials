package main

import (
	"fmt"
	"strconv"
)

var userInput string

func firstLoop() {
	for true {
		fmt.Println("Hello, World!")
		fmt.Println("Would you like to continue? (y/n)")

		fmt.Scanln(&userInput)

		if userInput == "n" {
			break
		}
		if userInput == "y" {
			continue
		} else {
			fmt.Println("Invalid input")
		}
	}
}
func secondLoop() {
	for {
		fmt.Print("Enter a number: ")
		fmt.Scanln(&userInput)

		number, _ := strconv.Atoi(userInput)

		for i := number; i >= 0; i-- {
			fmt.Printf("%v ", i)
		}
		fmt.Println()
		for i := 0; i <= number; i++ {
			fmt.Printf("%v ", i)
		}
		fmt.Println()

		fmt.Println("Would you like to continue? (y/n)")

		fmt.Scanln(&userInput)

		if userInput == "n" {
			break
		}
		if userInput == "y" {
			continue
		} else {
			fmt.Println("Invalid input")
		}
	}
}
func doorLoop(code int, guessCount int) bool {
	var result bool

	for {
		fmt.Println("Enter door code guess: ")
		fmt.Scanln(&userInput)
		userGuess, _ := strconv.Atoi(userInput)

		if userGuess == code && guessCount > 0 {
			fmt.Println("Your guess is right! Welcome to my house.")
			result = true
			break
		}
		if userGuess != code && guessCount > 0 {
			fmt.Println("That password was incorrect please try again.")
			guessCount--
		}
		if guessCount == 0 {
			fmt.Print("Too many incorrect attempts!")
			result = false
			break
		}
	}
	return result
}
func thirdLoop() {
	doorLocked := true
	const doorCode = 13579
	guessCount := 5

	for doorLocked {
		doorLocked = doorLoop(doorCode, guessCount)
	}
}

func main() {
	firstLoop()
	secondLoop()
	thirdLoop()
}
