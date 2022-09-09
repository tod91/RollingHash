package main

import (
	"fileio"
	"fmt"
)

func main() {
	file, err := fileio.Open("test.txt")
	if err != nil {
		fmt.Errorf("ERROR: ", err)
	}
	//fmt.Printf()
}
