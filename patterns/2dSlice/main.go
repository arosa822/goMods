/*
parseNum package
This is a package which contains a function for converting characters
(charToInt), in the form of runes, into numerals. It also contains a
function for parsing numerals from strings (parseNum).
*/
package main

import (
	"errors"
	"fmt"
)

/*
CharToNum function
This is a function which recieves a rune and converts it to a numeral
based on its ascii code by matching it against the numeral ascii codes
in a for loop. It subtracts 48 from the code to get the actual numeral,
as the ascii codes for numerals are 48-57 and simple subtraction gives
you the actual numeral.
*/
func CharToNum(r rune) (int, error) {
	for i := 48; i <= 57; i++ {
		if int(r) == i {
			return (int(r) - 48), nil
		}
	}
	return -1, errors.New("type: rune was not int")
}

/*
ParseNum function
This is a function which serves the purpose of identifying numerals inside
strings and returning them in a slice. It loops over the string, passing each
character to the charToNum function and identifying whether it should append the
output to the array by testing whether or not the error evaluates to nil.
*/
func ParseNum(str string) []int {
	var nums []int
	for _, val := range str {
		num, err := CharToNum(val)
		if err != nil {
			continue
		}
		nums = append(nums, num)
	}
	return nums
}

func main() {
	fmt.Println(ParseNum("000123456789000"))
}
