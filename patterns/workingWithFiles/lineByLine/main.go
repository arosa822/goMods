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
	var status bool
	fmt.Println(status)

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
			status = false
		}

		// check for failure
		if state == 1 && strings.Contains(line, "bomb") {
			status = true
		}

		if state == 1 && strings.Contains(line, "down") {
			s[0]++
			state = 0
			if status != true {
				// reset status
				status = true
				s[1]++

			} else {
				s[2]++
			}
		}
		// reset states status indicates a failure, state tells us if we are in
		// the middle of a process

	}

	fmt.Printf("Number of cycles: %v\n", s[0])
	fmt.Printf("Successfull Attempts %v\n", s[1])
	fmt.Printf("Failed Attempts: %v\n", s[2])

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
