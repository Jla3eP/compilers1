package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("input, output/inputFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	output, err := os.Create("input, output/outputFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	LexIt(input, output)
}
