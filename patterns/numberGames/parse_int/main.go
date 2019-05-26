package main

// this script takes any int and splits it up into []int that make it up
// i.e. 708 converts to [7,0,8]
// 4095 converts to [4,0,9,5]

import "fmt"

func main() {
	num := 60
	fmt.Println(num % 10)
}
