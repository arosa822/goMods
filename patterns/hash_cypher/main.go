package main

import "fmt"

// input comes in the formm:

//{ AAAAA BBBBB CCCCC DDDDD .....}
//{ AAAAA BBBBB CCCCC DDDDD .....}
//{ AAAAA BBBBB CCCCC DDDDD .....}

// var len int = 5     // length of a character
// var ht int = 3      // height of a character
// var char string = A // desired output

// Goal: want to output 2 dimentional string
// for testing this can be of the form:
// { 11111 22222 33333 .....
// { 11111 22222 33333 .....
// { 11111 22222 33333 .....
// where 1 equates to A

// simulated input stream - pre defined dimentions 5L x 3H
type wordList struct {
	row1, row2, row3 map[int][]string
}

func cypher(r string) {
	//ascii := []byte(r)

	//asciiShift := ascii - shift

	//fmt.Printf("%s:%d\n", r, ascii)
	//fmt.Printf("byteShift: %s", asciiShift)
}

func main() {
	cypher("E")
	fmt.Println("vim-go")
}
