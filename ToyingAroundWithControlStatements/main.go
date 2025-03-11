package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	var answer int
	fmt.Println("***** USING CONTROL STATEMENTS IN GO *****")
	fmt.Print("Technical note: There are no differences in these statements output-wise, just the handling")
	running := true

	for running == true{
		fmt.Println("We have: ")
		fmt.Println("1. If-Else Statement")
		fmt.Println("2. Switch Statement")
		fmt.Println("0. Exit")
		fmt.Print("Which one do you want? (eg: Type \"1\" for If-Else): ")
		fmt.Scan(&answer)

		switch answer {
		case 1: usingIfStatements()
		case 2: usingSwitchStatements()
		case 0: 
			fmt.Println("Thanks for using the program!")
			running = false
			os.Exit(0)
		default: 
			fmt.Println("That's not part of the options")
			fmt.Println("Once again. ")
		}
	}

	os.Exit(0)
}

func usingIfStatements() {
	var running = true
	for running == true {
		var answer string
		fmt.Print("Should I keep it going!? (Type \"yes\" or \"no\"): ")
		fmt.Scan(&answer)

		if strings.ToLower(answer) == "yes" {
			fmt.Println("Then let's keep it going!")
		} else if strings.ToLower(answer) == "no" {
			fmt.Println("Oh aight then")
			running = false
			break
		} else {
			fmt.Println("Please type in either yes or no!")
		}

		running = true
	}
}

func usingSwitchStatements() {
	var running bool = true
	for running == true {
		var answer string
		fmt.Print("Should I keep it going? (Type \"yes\" or \"no\")")
		fmt.Scan(&answer)

		switch answer {
		case "no":
			fmt.Println("Oh aight then")
			running = false
		case "yes":
			fmt.Println("Then let's keep it going!")
		default:
			fmt.Println("Please type \"yes\" or \"no\"")
		}
	}
}
