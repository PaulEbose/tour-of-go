/*
Package basics implements some of the things taught in the "Basics" section of "A Tour of Go".

  Declaring variables, calling functions, and all that jazz.

	commands:
		add  - arithmetic addition
		b2   - convert to binary
		b8   - convert to octal
		b16  - convert to hexadecimal
		div  - arithmetic division
		mul  - arithmetic multiplication
		sub  - arithmetic subtraction
		sqrt - square root of

	example:
		$ tog add 30 80 45
		$ tog sqrt 36 81 48
		$ tog b2 30
*/
package basics

import (
	"fmt"
	"log"
)

// Divide returns the division of all its arguments from left to right.
func Divide(nums ...float64) (total float64) {
	total = nums[0]
	for i := 1; i < len(nums); i++ {
		total /= nums[i]
	}
	return
}

// Multiply returns the multiplication of all its arguments.
func Multiply(nums ...float64) (total float64) {
	total = nums[0]
	for i := 1; i < len(nums); i++ {
		total *= nums[i]
	}
	return
}

// Subtract returns the subtraction of all its arguments from left to right.
func Subtract(nums ...float64) (total float64) {
	total = nums[0]
	for i := 1; i < len(nums); i++ {
		total -= nums[i]
	}
	return
}

// Sum returns the sum of all its arguments.
func Sum(nums ...float64) (sum float64) {
	for _, num := range nums {
		sum += num
	}
	return
}

// Sqrt returns the Sqrt of all its arguments.
func Sqrt(nums ...float64) (total []float64) {
	for _, num := range nums {
		sqrt := 1.0

		for range [12]int{} {
			sqrt -= (sqrt*sqrt - num) / (2 * sqrt)
		}

		total = append(total, sqrt)
	}
	return
}

// ConvertBase returns a value in the given base.
func ConvertBase(base int, nums ...int) (result []string) {
	if base < 2 || base > 16 {
		log.Fatalln("Omo! base '" + fmt.Sprint(base) + "' is not supported. \n Only bases 2 to 16 is supported.")
	}

	for _, num := range nums {
		var currentValue string

		for num >= base {
			// prepend before mutating num
			currentValue = fmt.Sprint(num%base) + currentValue
			num = num / base
		}

		switch num % base {
		case 10:
			currentValue = "A" + currentValue
		case 11:
			currentValue = "B" + currentValue
		case 12:
			currentValue = "C" + currentValue
		case 13:
			currentValue = "D" + currentValue
		case 14:
			currentValue = "E" + currentValue
		case 15:
			currentValue = "F" + currentValue
		default:
			currentValue = fmt.Sprint(num%base) + currentValue
		}

		result = append(result, currentValue)
	}

	return
}

/**
OBJECTIVES :
	- use multiple results (return values)
	- use byte data type
	- use named return values
	- use "naked" return for short functions
	- use short declarations with multiple initializers
	- use basic complex math functions
	- use variadic functions
*/
