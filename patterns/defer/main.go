package main

// take in a string output the reverse order of strings
import (
	"fmt"
	"strings"
)

var s string = "The quick brown fox"

func reversePrint(s string) {
	ss := strings.Split(s, " ")
	for _, f := range ss {
		defer fmt.Printf("%s ", f)
	}
}

func main() {
	fmt.Println(s)
	reversePrint(s)
}
