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
	stream       map[int]int
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
	r.sum = 0
	r.byteArray = nil

}

func (r *river) riverGuide(c chan []int) {

	var a []int

	a = append(a, r.number)
	for n := 0; n < 100; n++ {
		r.makeRiver()
		a = append(a, r.next)
		r.number = r.next
	}
	c <- a

}

func findCrossing(a, b []int) {
	var l []int
	for _, j := range a {
		for _, k := range b {
			if k == j {
				l = append(l, k)
			}
		}

	}
	fmt.Println(l[0])
}

func main() {
	var n, m river
	c := make(chan []int, 2)

	n.number = 1158
	m.number = 2085

	go m.riverGuide(c)
	go n.riverGuide(c)

	a, b := <-c, <-c

	fmt.Println(a)
	fmt.Println(b)
	findCrossing(a, b)

}
