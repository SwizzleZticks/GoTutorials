package main

import (
	"fmt"
	"strconv"
	"strings"
)

var userInput string

func firstLoop() {
loopyLoop:
	for {
		fmt.Println("Hello, World!")
		fmt.Println("Would you like to continue? (y/n)")

		fmt.Scan(&userInput)

		switch strings.ToLower(userInput) {
		case "y":
			continue
		case "n":
			break loopyLoop
		default:
			fmt.Println("That was a invalid, please type y or n")
		}
	}
}
func secondLoop() {
loopyLoop:
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
		fmt.Scan(&userInput)

		switch strings.ToLower(userInput) {
		case "y":
			continue
		case "n":
			break loopyLoop
		default:
			fmt.Println("That was a invalid, please type y or n")
		}
	}
}
func doorLoop(code int, guessCount int) bool {
	var result bool

guessLoop:
	for guessCount > 0 {
		fmt.Println("Enter door code guess: ")
		fmt.Scanln(&userInput)
		userGuess, _ := strconv.Atoi(userInput)

		if userGuess == code {
			result = false
			fmt.Println("Your guess is right! Welcome to my house.")
			break guessLoop
		} else if userGuess != code {
			fmt.Println("That password was incorrect please try again.")
			guessCount--
		}
		if guessCount == 0 {
			result = true
			fmt.Print("Too many incorrect attempts!")
			break guessLoop
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
