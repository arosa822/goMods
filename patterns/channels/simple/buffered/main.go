package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(c chan int) {
	value := rand.Intn(100)
	fmt.Printf("Calculated Random Value: %d\n", value)
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Println("this executes regardless as the send is now non-blocking")
}

func main() {
	fmt.Println("buffered channel Tutorial")
	valueChannel := make(chan int, 2)
	defer close(valueChannel)

	go CalculateValue(valueChannel)
	go CalculateValue(valueChannel)

	values := <-valueChannel
	fmt.Println(values)

	time.Sleep(1000 * time.Millisecond)
}
