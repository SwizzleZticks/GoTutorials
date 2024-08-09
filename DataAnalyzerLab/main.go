package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var userInput int
	isWithinRange := false

	fmt.Print("Please enter your name: ")
	var userName string
	fmt.Scanln(&userName)

	for true {
		for !isWithinRange {
			fmt.Printf("%s, Please enter a number between  1 and 100: ", userName)
			var input string
			fmt.Scan(&input)

			var err error
			userInput, err = strconv.Atoi(input)
			if err != nil {
				continue
			}

			if userInput >= 0 && userInput <= 100 {
				isWithinRange = true
			}
		}
		if userInput%2 != 0 && userInput < 60 {
			fmt.Printf("%s, that number is odd and less than 60.", userName)
		}
		if userInput%2 == 0 && userInput >= 2 && userInput <= 24 {
			fmt.Printf("%s, that number is even and less than 25.", userName)
		}
		if userInput%2 == 0 && userInput >= 26 && userInput <= 60 {
			fmt.Printf("%s, that number is even and between 26 than 60 inclusive.", userName)
		}
		if userInput%2 == 0 && userInput > 60 {
			fmt.Printf("%s, that number is even and greater than 60.", userName)
		}
		if userInput%2 != 0 && userInput > 60 {
			fmt.Printf("%s, that number is odd and greater than 60.", userName)
		}

		fmt.Println()
		fmt.Print("Do you want to stop the program (y/n): ")
		var userChoice string
		fmt.Scanln(&userChoice)

		if strings.ToLower(userChoice) == "y" {
			break
		}
		if strings.ToLower(userChoice) == "n" {
			userInput = 0
			isWithinRange = false
		}
	}

}
