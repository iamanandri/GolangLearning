package main

import (
	"fmt"
	"strings"
)

func main() {

	var running = true
	for running == true {
		var answer string
		fmt.Print("Should I keep it going!? (Type \"yes\" or \"no\"): ")
		fmt.Scan(&answer)

		if strings.ToLower(answer) != "no" && strings.ToLower(answer) != "yes" {
			fmt.Println("Please type in either yes or no!")
		} else if strings.ToLower(answer) == "no" {
			fmt.Println("Oh aight then")
			running = false
			break
		}

		running = true
	}
}
