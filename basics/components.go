/*
Package basics implements some of the things taught in the "Basics" section of "A Tour of Go".

  Declaring variables, calling functions, and all that jazz.

	commands:
		add  - arithmetic addition
		b2   - convert to binary
		b8   - convert to octal
		b16  - convert to hexadecimal
		div  - arithmetic division
		exp  - raise to the power of
		mod  - arithmetic modulus
		mul  - arithmetic multiplication
		sub  - arithmetic subtraction
		sqrt - square root of

	example:
		$ tog add 30 80 45
		$ tog sqrt 36 81 48
		$ tog b2 30
		$ tog b16 30*14
*/
package basics

import (
	"fmt"
	"log"
)

// Sum returns the sum of all its arguments.
func Sum(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

// ConvertBase returns a value in the given base.
func ConvertBase(base int, num int) (result string) {
	if base < 2 || base > 16 {
		log.Fatalln("Omo! base '" + fmt.Sprint(base) + "' is not supported. \n Only bases 2 to 16 is supported.")
	}

	for num >= base {
		// prepend before mutating num
		result = fmt.Sprint(num%base) + result
		num = num / base
	}

	switch num % base {
	case 10:
		result = "A" + result
	case 11:
		result = "B" + result
	case 12:
		result = "C" + result
	case 13:
		result = "D" + result
	case 14:
		result = "E" + result
	case 15:
		result = "F" + result
	default:
		result = fmt.Sprint(num%base) + result
	}

	return
}

// Divide returns the division of all its arguments from left to right.
func Divide(nums ...int) (total float64) {
	total = float64(nums[0])
	for i := 1; i < len(nums); i++ {
		total /= float64(nums[i])
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
