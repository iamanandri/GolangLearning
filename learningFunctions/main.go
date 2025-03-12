package main

import "fmt"

func main() {
	//Series of numbers
	numbers := []float64{98, 93, 77, 82, 83}

	//average := average(numbers)

	fmt.Printf("The average is: %f", average(numbers))
}

func average(x []float64) float64 {
	var total float64

	for _, v := range x {
		total += v
	}

	return (total / float64(len(x)))
}
