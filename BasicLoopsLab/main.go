package main

import "fmt"

func firstLoop() {
	var userInput string

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

}
func main() {
	firstLoop()
}
