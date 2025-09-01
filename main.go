package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Print("Invalid length of arguments")
		return
	}

	oldFile := os.Args[1]
	//newFile := os.Args[2]

	_, err := os.ReadFile(oldFile)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	
}
