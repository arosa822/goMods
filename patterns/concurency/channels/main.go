package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {

	fmt.Println("--simple channel---")
	s := []int{7, 2, 8, -9, 4, 0}
	fmt.Println(s)

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c

	fmt.Println(s[:len(s)/2])
	fmt.Println(x)

	fmt.Println(s[len(s)/2:])
	fmt.Println(y)
	fmt.Println(x + y)
	fmt.Println("----- end of simple channel----")

	fmt.Println("---start of buffered channel---")
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	val1, val2 := <-ch, <-ch
	fmt.Println(val1, val2)

	fmt.Println("---start of buffered channel---")

	fmt.Println("-----start range/close------")
	buf_c := make(chan int, 10)
	go fibonacci(cap(buf_c), buf_c)
	for i := range buf_c {
		fmt.Println(i)
	}

	fmt.Println("-----end range/close------")
}
