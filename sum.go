package main

import (
	"fmt"
	"math"
)

// sumOfDigits calculates the sum of the digits of a given integer.
func sumOfDigits(n int) int {
	n = int(math.Abs(float64(n))) // Handle negative numbers
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main() {
	var number int

	fmt.Println("Sum of Digits Calculator")
	fmt.Println("========================")
	fmt.Print("Enter a number: ")

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error: Please enter a valid integer.")
		return
	}

	result := sumOfDigits(number)
	fmt.Printf("The sum of the digits of %d is %d\n", number, result)
}
