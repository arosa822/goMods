package main

// this script takes any int and splits it up into []int that make it up
// i.e. 708 converts to [7,0,8]
// 4095 converts to [4,0,9,5]

import (
	"fmt"
)

type river struct {
	number, next int
	byteArray    []byte
	sum          byte
}

func (r *river) makeRiver() {
	var rightMost, tempIntVar int
	tempIntVar = r.number

	// get the list of digits that make a number
	for {
		rightMost = tempIntVar % 10
		// append right most int to begining of list
		r.byteArray = append(r.byteArray, byte(rightMost))
		// take off last processed number
		tempIntVar /= 10
		if tempIntVar == 0 {
			break
		}
	}

	// add up the numbers that make up intvar
	for _, j := range r.byteArray {
		r.sum += j
	}
	r.next = r.number + int(r.sum)

}

func main() {
	var n river
	n.number = 123
	n.makeRiver()
	fmt.Println(n.byteArray)
	fmt.Println(n.sum)
	fmt.Println(n.next)
	/*
		var intVar = 123
		var rightMost, tempIntVar int
		var byteArray []byte
		tempIntVar = intVar

		// get the list of digits that make a number
		for {
			rightMost = tempIntVar % 10
			// append right most int to begining of list
			byteArray = append(byteArray, byte(rightMost))
			// take off last processed number
			tempIntVar /= 10
			if tempIntVar == 0 {
				break
			}
		}

		// add up the numbers that make up intvar
		fmt.Println(byteArray)
		var k byte = 0
		for _, j := range byteArray {
			k += j
		}

	*/

}
