package main

import (
	"fmt"
)

func main() {
	var printValue string = "Hello, World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	result, remaineder := intDivision(numerator, denominator)
	var err error
	if err != nil {
		fmt.Print(err.Error())
	} else if remaineder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with a remainder of %v", result, remaineder)
	}
	fmt.Printf("The result of %d divided by %d is %d with a remainder of %d\n", numerator, denominator, result, remaineder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {
	if denominator == 0 {
		return 0, 0
	}
	var result int = numerator / denominator
	var remaineder int = numerator % denominator
	return result, remaineder
}
