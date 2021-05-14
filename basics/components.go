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

// Sum returns the sum of all its arguments.
func Sum(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
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
