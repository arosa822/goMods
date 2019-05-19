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

	binaryRep := stringToBinary("C")

	decoder := make(map[int32]int)
	decoder[48] = 0
	decoder[49] = 1
	decoder[10] = 9
	decoded := []int{}

	// take in int32's and place in list of 0's and 1's
	for n, j := range binaryRep {
		fmt.Printf("%d: %v %T\n", n, j, j)
		b, ok := decoder[j]
		if ok {
			//store b in string[]
			decoded = append(decoded, b)
		}

	}

	fmt.Println(decoded)
	//var state int = 0
	//var count int = 0
	count := []int{}
	val := []int{}
	var tempCount int

	for n, j := range decoded {

		if n == 0 {
			fmt.Println("first iteration")
			tempCount = 1
		} else if j != decoded[n-1] {
			val = append(val, decoded[n-1])
			count = append(count, tempCount-1)

			tempCount = 1

		}

		tempCount++
	}
	fmt.Println(val)
	fmt.Println(count)
}
