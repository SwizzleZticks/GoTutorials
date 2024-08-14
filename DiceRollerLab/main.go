package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
)

var userInput string

func main() {
	activeLoop := true
	for activeLoop {
		diceSides := GetDiceSidesNumber()

		PromptUser()
		firstDice := randRange(1, diceSides)
		secondDice := randRange(1, diceSides)

		fmt.Printf("You rolled %v and %v", firstDice, secondDice)
		fmt.Println()

		if diceSides == 6 {
			fmt.Println(CheckWinningCombos(firstDice, secondDice))
			fmt.Println(PrintDiCeCombos(firstDice, secondDice))
		}

		activeLoop = GetUserChoice()
		userInput = ""
	}
}
func PrintDiCeCombos(firstDice, secondDice int) string {
	var result string

	if firstDice == 1 && secondDice == 1 {
		result = "Snake Eyes: Two 1's"
	}
	if firstDice == 1 && secondDice == 2 {
		result = "Ace Deuce: A 1 and 2"
	}
	if firstDice == 6 && secondDice == 6 {
		result = "Box Cars: Two 6's"
	} else {
		result = ""
	}
	return result
}

func CheckWinningCombos(firstDice, secondDice int) string {
	diceTotal := firstDice + secondDice
	var result string

	if diceTotal == 11 || diceTotal == 7 {
		result = fmt.Sprintf("Win: Total: %v", diceTotal)
	}
	if diceTotal == 2 || diceTotal == 3 || diceTotal == 11 {
		result = fmt.Sprintf("Craps: Total: %v", diceTotal)
	} else {
		result = ""
	}

	return result
}

func GetUserChoice() bool {
	for {
		fmt.Print("Do you want to roll the dice again? (y/n): ")
		fmt.Scan(&userInput)

		switch strings.ToLower(userInput) {
		case "y":
			return true
		case "n":
			return false
		}
	}
}

func randRange(min, max int) int {
	return rand.IntN(max+1-min) + min
}

func PromptUser() {
	fmt.Println("Press enter to roll the dice...")
	fmt.Scanln()
}

func GetDiceSidesNumber() int {
	for {
		print("Please enter a number for the sides of the 2 dice: ")
		fmt.Scan(&userInput)

		number, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Invalid input, please enter a number for dice sides: ")
			fmt.Scan(&userInput)
		} else {
			return number
		}
	}
}
