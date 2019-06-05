package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	arg := os.Args[1]

	file, err := os.Open(arg)
	s := make([]int, 3)
	//fmt.Println(s)

	var state byte
	state = 0
	var status bool

	var errorLog []string
	//fmt.Println(status)

	//fmt.Println(state)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		if state == 0 && strings.Contains(line, "initializing") {
			state = 1
			status = false
		}

		// check for failure
		if state == 1 && strings.Contains(line, "bomb") {
			status = true
			errorLog = append(errorLog, line)

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

	fmt.Println("Errors Found:")

	for _, i := range errorLog {
		fmt.Println(i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
