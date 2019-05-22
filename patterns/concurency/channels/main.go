package main

import (
	"fmt"
	//"time"
)

func multiplyByTwo(num int, out chan<- int) {
	result := num * 2
	out <- result
}
func main() {
	n := 3

	out := make(chan int)

	go multiplyByTwo(n, out)

	fmt.Println(<-out)

}
