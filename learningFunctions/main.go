package main

import "fmt"

func main() {
	//Series of numbers
	numbers := []float64{98, 93, 77, 82, 83}

	average := average(numbers)

	fmt.Printf("The average is: %f", average)
}

func average(x []float64) float64 {
	var total float64

	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	return (total / float64(len(x)))
}
