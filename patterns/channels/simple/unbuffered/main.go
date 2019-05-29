package main

// introduction to unbuffered channels
import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(c chan int) {
	value := rand.Intn(100)
	fmt.Println("Calculated Random Value: {}:", value)
	time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Println("only executes after another goroutine performs a receive on the channel")
}
func main() {
	fmt.Println("Go -unbuffered channel- tutorial")
	valueChannel := make(chan int)
	defer close(valueChannel)

	go CalculateValue(valueChannel)
	go CalculateValue(valueChannel)

	values := <-valueChannel
	fmt.Println(values)
	fmt.Println(valueChannel)
}
