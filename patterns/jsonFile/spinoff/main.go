package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// flags struct which contains an array of flags
type Flags struct {
	Flags []Flag `json:"flags"`
}

// Flag struct which contains all the flags we are looking for
type Flag struct {
	Start    string `json:"start"`
	End      string `json:"end"`
	FailFlag string `json:"failFlag"`
}

func main() {
	// open our jsonFile
	jsonFile, err := os.Open("flags.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened .json file")

	defer jsonFile.Close()

	//read our opened jsonFile as a bytearray
	byteValue, _ := ioutil.ReadAll(jsonFile)

	//Initialize our Flags array
	var flags Flags

	// unmarshal our byteArray which contains our json's file content into
	// 'flags which we defined above
	json.Unmarshal(byteValue, &flags)

	fmt.Println(flags.Flags[0].Start)
}
