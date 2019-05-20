package main

import "fmt"

// CC:
var expected string = "0 0 00 0000 0 000 00 0000 0 00"

//
type binString struct {
	asciiVal string //ascii value of strings
	binCode  []int  //resulting code
	binVal   []int
	bincount []int
}

type keys struct {
	keyMap map[int32]int
}

//convert s (string) to ascii
func (b *binString) stringToBinary(s string) {
	b.asciiVal = ""
	for _, i := range s {
		b.asciiVal = fmt.Sprintf("%s%b\n", b.asciiVal, i)
	}
}

// convert to ascii to 0's and 1's based on mapping
func (b *binString) convertToBinary(k *keys) {
	for _, j := range b.asciiVal {
		val, ok := k.keyMap[j]
		if ok {
			b.binCode = append(b.binCode, val)
		}
	}
}

// scan resulting binCode for unique characters
func (b *binString) scanUnique() {
	/// add stuff in here

}

// constructor function for keymap
func newKey() *keys {
	var k keys
	k.keyMap = make(map[int32]int)
	k.keyMap[48] = 0
	k.keyMap[49] = 1
	return &k
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

var b binString

func main() {

	b.stringToBinary("CC")
	k := newKey()
	b.convertToBinary(k)

	fmt.Println(b.binCode)

	count := []int{}
	val := []int{}
	var tempCount int

	for n, j := range b.binCode {

		if n == 0 {
			fmt.Println("first iteration")
			tempCount = 1
		} else if j != b.binCode[n-1] && n != 9 {
			val = append(val, b.binCode[n-1])
			count = append(count, tempCount-1)

			tempCount = 1

		}

		tempCount++
	}

	// reverse search for the last element
	tempCount = 0
	for n := len(b.binCode) - 1; n > 0; n-- {
		if n == len(b.binCode)-1 {
			tempCount = 1
		} else if n < len(b.binCode)-2 && b.binCode[n] != b.binCode[n+1] {
			val = append(val, b.binCode[n+1])
			count = append(count, tempCount-1)
			break
		}
		tempCount++

	}

	fmt.Println(val)
	fmt.Println(count)
	fmt.Printf("S:")

	for t := 0; t < len(val); t++ {
		if val[t] == 0 {
			fmt.Print("00 ")
		} else {
			fmt.Print("0 ")
		}

		for n := 0; n < count[t]; n++ {
			fmt.Print("0")
		}
		fmt.Print(" ")

	}
	fmt.Printf("\nE:%s", expected)
}
