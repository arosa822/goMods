package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(values chan int) {
	value := rand.Intn(100)
	fmt.Println("Calculated random value: {},", value)
	values <- value
}

func tick(duration int, clock chan int) {
	n := 0
	for {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("tick")
		n++
		if n > duration {
			break
		}
	}
	clock <- n

}

func main() {
	fmt.Println("Go Channel Tutorial")

	values := make(chan int)
	defer close(values)

	go CalculateValue(values)
	value := <-values
	fmt.Println(value)
}
