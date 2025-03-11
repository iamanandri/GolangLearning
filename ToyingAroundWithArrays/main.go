package main

import (
	"fmt"
)

func main() {
	fmt.Printf("***** Toying Around with Arrays *****")
	fmt.Println("Which one do you wanna do?")
	fmt.Println("1. Calculating Average with Input")
	fmt.Println("2. Slicing and Dicing Arrays")
	fmt.Println("3. Working with Maps")
	fmt.Println("4. Maps in Maps")
	fmt.Println("5. Smallest Number in a List")

	fmt.Print("Choose with a number: ")

	var answer int
	fmt.Scan(&answer)
	switch answer {
	case 1:
		calculatingAverageWithInput()
	case 2:
		slicingAndDicingArrays()
	case 3:
		mapsHandling()
	case 4:
		mapInMaps()
	case 5:
		smallestNumberInList()
	default:
		fmt.Println("That wasn't a number >:(")
	}

}

func calculatingAverageWithInput() {
	fmt.Println("***** Calculating Average with Input *****")
	fmt.Println("Enter 5 numbers and the program will give you the TOTAL and AVERAGE")

	var numbers [5]float64
	var total float64
	for i, value := range numbers {
		fmt.Printf("Number %d: ", i+1)
		fmt.Scan(&value)
		total += value
	}

	fmt.Printf("The total is: %f\n", total)
	fmt.Printf("The average is: %f\n", total/float64(len(numbers)))
}

func slicingAndDicingArrays() {
	fmt.Println("***** Slicing and Dicing Arrays *****")

	//Commented includes different ways of slicing arrays

	//var x []float64
	//x := make([]float64, 5)
	//x := make([]float64, 5, 10)

	//[low:high]
	//arr:= [5]float64{1,2,3,4,5}
	//x := arr[1:4]

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func mapsHandling() {
	//var x map[string]int
	//elements := make(map[string]string)
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	//name, ok := elements["Ne"]

	if name, ok := elements["Ne"]; ok {
		fmt.Println(name, ok)
	}

	//fmt.Println(name, ok)
}

func mapInMaps() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Undisputed"]; ok {
		fmt.Println(el["name"], el["state"])
	} else {
		fmt.Println("That's not one of the element")
	}
}

func smallestNumberInList() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	var smallest int = x[0]

	for _, value := range x {
		if value < smallest {
			smallest = value
			fmt.Printf("Current number: %d\n", value)
		}
	}

	//slices.Sort(x)
	//fmt.Println("Smallest number is:", x[0])

	fmt.Println("Smallest number is:", smallest)

}
