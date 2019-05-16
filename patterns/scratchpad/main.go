package main

import "fmt"

func main() {
	b1 := []byte{2, 2}
	s := []string{"el0", "el1", "el3"}
	fmt.Println(b1)
	fmt.Println(b1[0])
	fmt.Printf("%t\n", b1[0] > 1)
	fmt.Println(s)

	fmt.Println("-------------")

	fmt.Printf("string: %T\n", s[1])
	ascii := []byte(s[1])
	fmt.Printf("converted string: %T\n", ascii)
	//var number int = 3
	//	b := []byte{number}
	//	fmt.Printf("string: %s, ascii: %v\n int: %d byte to compare: %v\n", s[1], ascii[0], number, b)
	//	fmt.Printf("ascii > %d: %t\n", number, ascii[0] > b[0])

}
