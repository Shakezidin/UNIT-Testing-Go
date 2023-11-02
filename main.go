package main

import (
	"fmt"
	"github.com/Shakezidin/unittesting/math" 
)

func main() {
	resultAdd := math.Add(1, 5)
	resultSubtract := math.Subtract(7, 3)
	resultDivision := math.Division(10, 2)
	resultMultiplication := math.Multiplication(3, 4)


	// Print the results
	fmt.Printf("Addition Result: %d\n", resultAdd)
	fmt.Printf("Subtraction Result: %d\n", resultSubtract)
	fmt.Printf("Division Result: %f\n", resultDivision)
	fmt.Printf("Multiplication Result: %d\n", resultMultiplication)
}
