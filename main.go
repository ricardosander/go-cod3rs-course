package main

import (
	"fmt"
	"os"
)

const PI float64 = 3.1415

func main() {

	fmt.Println("Hello Go")
	var input string

	for {

		showMenu()
		input = readCommand(input)

		switch input {
		case "0":
			justExit()
		case "1":
			explainVariables()
		case "2":
			explainTypes()
		case "3":
			explainPrintAndFormating()
		case "4":
			explainPointers()
		default:
			exitWithError()
		}
	}
}

func showMenu() {

	printLine()

	fmt.Println("Choose what you want to learn:")
	fmt.Println("	1 - How to create variables and constants")
	fmt.Println("	2 - What types we have on Go")
	fmt.Println("	3 - Print and formating")
	fmt.Println("	4 - Tell me about pointers")
	fmt.Println("	0 - Exit gracefully")
	fmt.Println("	Anything else - Just get me out of here")

	printLine()
}

func readCommand(input string) string {
	fmt.Print("Choose your Destiny: ")
	fmt.Scan(&input)
	fmt.Println("Your destiny is", input)
	return input
}

func giveMeSomeSpace() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func justExit() {
	giveMeSomeSpace()
	fmt.Println("Ok, see you later alligator")
	os.Exit(0)
}

func exitWithError() {
	giveMeSomeSpace()
	fmt.Println("Ouch, how rude")
	os.Exit(255)
}

func printLine() {
	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println()
}
