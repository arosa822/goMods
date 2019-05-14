package main

import "fmt"

type square struct {
	l, w float64
}

type circle struct {
	r float64
}

//type shape struct {
//	square
//	circle
//}

type area interface {
	area()
	update()
}

func (c *circle) area() float64 {
	return 3.14 * c.r
}

func (s square) area() float64 {
	return s.l * s.w
}

// pass by reference -- use this to modify type values
func (c *circle) update(r float64) {
	c.r = r
}

func (s *square) update(l, w float64) {
	s.l = l
	s.w = w
}

// pass by reference
func main() {
	var box = square{l: 3, w: 10}
	box.update(30, 10)
	fmt.Printf("Shape: %T\n", box)
	fmt.Printf("l:%.2f w:%.2f\n", box.l, box.w)
	fmt.Printf("area: %.2f\n", box.area())

	var hole = circle{r: 10}
	fmt.Printf("Shape: %T\n", hole)
	fmt.Printf("r: %.2f\n", hole.r)
	fmt.Printf("area: %.2f\n", hole.area())

}
