package main

import "fmt"

// CC:
var expected string = "0 0 00 0000 0 000 00 0000 0 00"

//
type binString struct {
	asciiVal string //ascii value of strings
	binCode  []int  //resulting code
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
	for n, j := range b.asciiVal {
		val, ok := k.keyMap[j]
		if ok {
			b.binCode = append(b.binCode, val)
		}
	}
}

// constructor function for keymap
func newKey() *keys {
	var k keys
	k.keyMap = make(map[int32]int)
	k.keyMap[48] = 0
	k.keyMap[49] = 1
	return &k
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
		} else if j != b.binCode[n-1] && b.binCode[n-1] != 9 {
			val = append(val, b.binCode[n-1])
			count = append(count, tempCount-1)

			tempCount = 1

		}

		tempCount++
	}
	fmt.Println(val)
	fmt.Println(count)

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
	fmt.Printf("\n Expected: %s", expected)
}
