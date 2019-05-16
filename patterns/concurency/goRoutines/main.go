package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("Goroutine_1")
	go f("Goroutine_2")

	go func(msg string) {
		for i := 0; i < 5; i++ {
			fmt.Println(msg, ":", i)
		}
	}("we are doing this")

	fmt.Scanln()
	fmt.Println("done")
}
