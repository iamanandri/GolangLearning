package main

import "fmt"

func main() {
	fmt.Println("=== Average Function ===")
	//Series of numbers
	numbers := []float64{98, 93, 77, 82, 83}	
	//average := average(numbers)

	fmt.Printf("The average is: %f\n", average(numbers))

	//Returning Multiple Types
	fmt.Println("=== Returning Multiple Types ===")
	fmt.Println(returningMultipleTypes())

	//Variadic Function
	fmt.Println("=== Add Three Numbers ===")
	vfInt := []int{5,5,5,5,5}
	fmt.Println(addThreeNumbers(vfInt...))

	//Closures
	fmt.Println("=== Closures with Increment Values ===")
	incrementValue := 0
	incrementFunc := func() int {
		incrementValue++
		return incrementValue
	}

	for i := 0; i < 5; i++{
		println(incrementFunc())
	}
}

func average(x []float64) float64 {
	var total float64

	for _, v := range x {
		total += v
	}

	return (total / float64(len(x)))
}

func returningMultipleTypes() (int, string, float64){
	return 10, "Lionel Messi", 98.98
}

func addThreeNumbers(numbers ...int) int {
	var total int
	for _, value := range numbers {
		total += value
	}

	return total
}

