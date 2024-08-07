package main

import (
	"fmt"
	"strconv"
)

func main() {
	getText("Enter some text: ")
	fmt.Println()
	getAddNumber("Enter a number: ")
	fmt.Println()
	getAddDecimalNumber("Enter a decimal number: ")
	fmt.Println()
	addTwoNumbers("Enter a number to add: ")
	fmt.Println()
	addTwoDecimalNumbers("Enter a decimal number to add: ")
	fmt.Println()
	multiplyTwoNumbers("Enter a number to multiply: ")
	fmt.Println()
	multiplyTwoDecimalNumbers("Enter a decimal number to multiply: ")
	fmt.Println()
}

func getText(text string) {
	var userInput string
	fmt.Println(text)
	fmt.Scanln(&userInput)

	fmt.Println(userInput)
}

func getAddNumber(text string) {
	var userInput string
	fmt.Println(text)
	fmt.Scanln(&userInput)

	num, err := strconv.Atoi(userInput)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num + 1)
}

func getAddDecimalNumber(text string) {
	var userInput string
	fmt.Println(text)
	fmt.Scanln(&userInput)

	num, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num + .5)
}

func addTwoNumbers(text string) {
	var firstInput string
	var secondInput string

	fmt.Println(text)
	fmt.Scanln(&firstInput)

	fmt.Println("Enter a another number: ")
	fmt.Scanln(&secondInput)

	number1, err := strconv.Atoi(firstInput)
	if err != nil {
		fmt.Println(err)
	}
	number2, err := strconv.Atoi(secondInput)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The sum of the numbers is: ", number1+number2)
}

func addTwoDecimalNumbers(text string) {
	var firstInput string
	var secondInput string

	fmt.Println(text)
	fmt.Scanln(&firstInput)

	fmt.Println("Enter a another number: ")
	fmt.Scanln(&secondInput)

	number1, err := strconv.ParseFloat(firstInput, 64)
	if err != nil {
		fmt.Println(err)
	}
	number2, err := strconv.ParseFloat(secondInput, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The sum of the numbers is: ", number1+number2)
}

func multiplyTwoNumbers(text string) {
	var firstInput string
	var secondInput string

	fmt.Println(text)
	fmt.Scanln(&firstInput)

	fmt.Println("Enter a another number to multiply: ")
	fmt.Scanln(&secondInput)

	number1, err := strconv.Atoi(firstInput)
	if err != nil {
		fmt.Println(err)
	}
	number2, err := strconv.Atoi(secondInput)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The sum of the numbers multiplied is: ", number1*number2)
}
func multiplyTwoDecimalNumbers(text string) {
	var firstInput string
	var secondInput string

	fmt.Println(text)
	fmt.Scanln(&firstInput)

	fmt.Println("Enter a another decimal number to multiply: ")
	fmt.Scanln(&secondInput)

	number1, err := strconv.ParseFloat(firstInput, 64)
	if err != nil {
		fmt.Println(err)
	}
	number2, err := strconv.ParseFloat(secondInput, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("The sum of the decimal numbers multiplied is: ", number1*number2)
}
