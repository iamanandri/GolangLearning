package main

import (
	"fmt"
	//"strings"
)

func main(){
	fmt.Println("***** OBLIGATORY FIZZBUZZ PROBLEM *****")
	fmt.Println()

	//Print 1-100
	//Multiples of 3 will print Fizz
	//Multiples of 5 will print Buzz
	//Multiple of both 3 and 5 will print FizzBuzz

	for i := 1; i <= 100; i++ {
		if i % 5 == 0 && i % 3 == 0 {
			fmt.Print("FizzBuzz ")
		}  else if i % 5 == 0 {
			fmt.Print("Buzz ")
		} else if i % 3 == 0 {
			fmt.Print("Fizz ")
		} else {
			fmt.Print(i, " ")
		}
	}
}