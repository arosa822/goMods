package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var age int
	fmt.Println("please enter your age:")
	_, err := fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Println("value entered was not an integer..")
		os.Exit(1)

	}
	// reading a string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ := reader.ReadString('\n')

	fmt.Println("Your name is ", name, " and your age is ", age)

	fmt.Printf("name type: %T\nage type: %T\n", name, age)
}
