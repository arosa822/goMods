package main

import "fmt"

var l []int
var d []int

func main() {
	l = []int{1, 2, 3, 4, 5}

	for t := 0; t < len(l); t++ {
		for _, i := range l {
			if i != l[t] {
				diff := i - l[t]
				if diff < 0 {
					diff = -diff
					d = append(d, diff)

				}
			}
		}

	}

	fmt.Println(d)
}
