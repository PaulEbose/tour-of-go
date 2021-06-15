package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/paulebose/tog/basics"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("	Expected a command!")
		fmt.Println(usageMsg)
		os.Exit(1)
	}

	var nums []float64

	for _, v := range os.Args[2:] {
		n, err := strconv.ParseFloat(v, 64)
		if err == nil {
			nums = append(nums, n)
		}
	}

	switch os.Args[1] {
	case "add":
		fmt.Println(basics.Sum(nums...))
	case "b2":
		// could easily have done this
		// n, err := strconv.ParseInt(v, 0, 0)
		fmt.Println(basics.ConvertBase(2, floatToInt(nums)...))
	case "b8":
		fmt.Println(basics.ConvertBase(8, floatToInt(nums)...))
	case "b16":
		fmt.Println(basics.ConvertBase(16, floatToInt(nums)...))
	case "div":
		fmt.Println(basics.Divide(nums...))
	case "mul":
		fmt.Println(basics.Multiply(nums...))
	case "sub":
		fmt.Println(basics.Subtract(nums...))
	case "sqrt":
		fmt.Println(basics.Sqrt(nums...))
	default:
		fmt.Printf("	Invalid command: \"%s\"\n", os.Args[1])
		fmt.Println(usageMsg)
	}
}

var usageMsg = `
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
`

func floatToInt(nums []float64) []int {
	var numInts []int
	for _, v := range nums {
		numInts = append(numInts, int(v))
	}
	return numInts
}
