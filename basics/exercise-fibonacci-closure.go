package basics

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	nextNumber := 0
	n := 0
	tmp := 0

	return func() int {
		if nextNumber <= 0 {
			nextNumber++
			return n
		}

		tmp = n + nextNumber
		n = nextNumber
		nextNumber = tmp

		return n
	}
}

func Fibonacci() {
	f := fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Println(f()) // 0,1,1,2,3,5,8,13,21,34,55,89,144...
	}
}
