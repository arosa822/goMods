package main

import (
	"fmt"
	"strings"
)

// the object of this program is to sort through a list of files and determine
// the MIME type with the use of a hash table

// starting values - given string
var s = []string{"apple.html", "fdaf", "apple.png", "banana.kdj", "apple.git@github.com:arosa822/goMods.gitf"}
var mimeList = []string{"text/html", "image/png", "image/gif"}
var exList = []string{"html", "png", "gif"}

func makeMap(key, val []string) map[string]string {
	mapOut := map[string]string{}
	for n, i := range key {
		mapOut[i] = val[n]
	}
	return mapOut
}

// sample values:

//var m map[string]int

//mMap := map[string]string{
//	"html": "text/html",
//	"png":  "image/png",
//	"gif":  "image/gif",
//}

func main() {

	mMap := makeMap(exList, mimeList)

	fmt.Println("----------")

	for n, i := range s {
		i := strings.Split(i, ".")
		//fmt.Printf("%d:%s\n", n, i[len(i)-1])
		var lookup string = i[len(i)-1]
		//fmt.Printf("%T\n", lookup)
		result, ok := mMap[lookup]
		//fmt.Printf("%d:%s:%t\n", n, result, ok)
		if ok {
			fmt.Printf("%d: %s :%s\n", n, i, result)
		}

	}
}
