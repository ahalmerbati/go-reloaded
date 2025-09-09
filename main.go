package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("invalid length of arguments")
		return
	}

	sampleFile := os.Args[1]
	//resultFile := os.Args[2]

	_, err := os.ReadFile(sampleFile)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
}
