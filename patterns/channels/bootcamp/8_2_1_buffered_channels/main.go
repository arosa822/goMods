package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2

	c3 := func() {
		fmt.Println("in goRoutin 1")
		c <- 3
	}
	c4 := func() {
		fmt.Println("in goRoutine 2")
		c <- 4
	}

	go c3()
	go c4()
	fmt.Println(<-c)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
