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
	"strconv"
)

// Sum returns the sum of all its arguments.
func Sum(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

// ConvertBase returns the value in the given base.
func ConvertBase(base int, num int) int {
	var resultInString string

	for num >= base {
		// prepend before mutating num
		resultInString = fmt.Sprint(num%base) + resultInString
		num = num / base
	}

	// insert the left over number
	resultInString = fmt.Sprint(num%base) + resultInString

	result, err := strconv.Atoi(resultInString)
	if err != nil {
		log.Fatalln(err)
	}

	return result
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
