package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	isLoopActive := true

	studentNames := [4]string{"Josh", "Eric", "Anthony", "Felipe"}
	studentTowns := [4]string{"Ligonier", "Sussex", "Sussex", "Santiago"}
	studentFoods := [4]string{"Barbeque", "Pizza", "Chicken", "Salad"}

	for isLoopActive {
		index := GetUserInput("Which student are you inquiring about? Enter a number (1-4):")

		prompt := fmt.Sprintf("What would you like to know about %s? Enter \"favorite food\" or \"hometown\": ", studentNames[index-1])
		studentCategory := GetStudentCategory(prompt)
		PrintUserInfo(studentCategory, studentNames[index-1], studentFoods[index-1], studentTowns[index-1])

		isLoopActive = GetUserChoice()
	}
}

func GetUserChoice() bool {
	for {
		var userInput string
		fmt.Print("Do you want to inquire about another student? (y/n): ")
		fmt.Scan(&userInput)

		switch strings.ToLower(userInput) {
		case "y":
			return true
		case "n":
			return false
		default:
			fmt.Println("Invalid choice. Please enter 'y' or 'n'.")
		}
	}
}

func PrintUserInfo(category string, name string, food string, town string) {
	switch category {
	case "favorite food":

		fmt.Printf("%s's favorite food is: %s\n", name, food)
	case "hometown":
		fmt.Printf("%s's hometown is: %s\n", name, town)
	default:
		fmt.Println("Invalid category.")
	}
}

func GetStudentCategory(message string) string {
	for {
		fmt.Println(message)
		reader := bufio.NewReader(os.Stdin)

		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]

		if input == "hometown" || input == "favorite food" {
			return input
		} else {
			fmt.Println("That category does not exist.")
		}
	}
}

func GetUserInput(message string) int {
	for {
		var userInput string

		fmt.Println(message)
		fmt.Scan(&userInput)

		number, err := strconv.Atoi(userInput)
		if err != nil || number < 1 || number > 4 {
			fmt.Println("Invalid input, please enter a number from 1 to 4.")
		} else {
			return number
		}
	}
}
