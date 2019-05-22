package main

import "fmt"

type biteString struct {
	val string
}

func (b *biteString) convert(s string) {
	for _, i := range s {
		b.val = fmt.Sprintf("%s%.8b\n", b.val, i)
	}
}

var s1 string = "a"
var s2 string = "b"

func main() {
	var b biteString
	b.convert("%")
	fmt.Println(b.val)
	fmt.Println(s1 + s2)
}
