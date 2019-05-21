package main

import "fmt"

// CC:
var expected string = "0 0 00 0000 0 000 00 0000 0 00"

//
type binString struct {
	asciiVal string //ascii value of strings
	binCode  []int  //resulting code
	binVal   []int
	binCount []int
}

type keys struct {
	keyMap map[int32]int
}

//convert s (string) to ascii
func (b *binString) stringToBinary(s string) {
	b.asciiVal = ""
	for _, i := range s {
		b.asciiVal = fmt.Sprintf("%s%.7b\n", b.asciiVal, i)
	}
	fmt.Println(b.asciiVal)
}

// convert to ascii to 0's and 1's based on mapping
func (b *binString) convertToBinary(k *keys) {
	for _, j := range b.asciiVal {
		val, ok := k.keyMap[j]
		if ok {
			b.binCode = append(b.binCode, val)
		}
		fmt.Println(b.binCode)
	}
}

// scan resulting binCode for unique characters
func (b *binString) scanUnique() {
	/// add stuff in here

	tempCount := 1
	for n, j := range b.binCode {
		if n == 0 {
			tempCount = 1
		} else if j != b.binCode[n-1] {
			b.binVal = append(b.binVal, b.binCode[n-1])
			b.binCount = append(b.binCount, tempCount-1)

			tempCount = 1
		}
		tempCount++
	}

	tempCount = 0
	for n := len(b.binCode) - 1; n > 0; n-- {
		if b.binCode[n] != b.binCode[n-1] {
			b.binVal = append(b.binVal, b.binCode[n])
			b.binCount = append(b.binCount, tempCount)
			break
		}
		tempCount++
	}

}

func (b binString) display() {
	for t := 0; t < len(b.binVal); t++ {
		if b.binVal[t] == 0 {
			fmt.Print(" 00 ")
		} else {
			fmt.Print(" 0 ")
		}
		for n := 0; n < b.binCount[t]; n++ {
			fmt.Print("0")
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

	b.stringToBinary("%")
	k := newKey()
	b.convertToBinary(k)

	b.scanUnique()

	fmt.Println(b.binCode)

	fmt.Printf("\nE:%s", expected)

	fmt.Printf("\nR:")
	b.display()
	fmt.Printf("\n%v\n%v\n", b.binVal, b.binCount)

}
