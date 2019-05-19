package main

import "fmt"

func stringToBinary(s string) (b string) {
	binaryString := ""
	for _, i := range s {
		binaryString = fmt.Sprintf("%s%b\n", binaryString, i)
	}
	return binaryString
}

func main() {
	fmt.Println(stringToBinary("C"))
}
