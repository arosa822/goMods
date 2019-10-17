package main

import (
	"fmt"
)

// Sum calculates the sum of x and y
func Sum(x, y int) int {
	return x + y
}

func main() {
	fmt.Printf("Sum: %d\n", Sum(5, 5))

}
