package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("this is a test")
	go compute(10)
	go compute(10)

	// insert these so go routines do not terminate before main does
	//var input string
	//fmt.Scanln(&input)

}
