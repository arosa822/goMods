package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("file.txt")
	s := make([]int, 3)
	fmt.Println(s)

	var state byte
	state = 0
	fmt.Println(state)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if state == 0 && strings.Contains(line, "initializing") {
			state = 1
		}

		if state == 1 && strings.Contains(line, "down") {
			s[0]++
		}

	}

	fmt.Printf("Number of cycles: %v\n", s[0])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
