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

func main() {
	var b biteString
	b.convert("%")
	fmt.Println(b.val)
}
