package main

import (
	"fmt"
	"go-reloaded/functions"
	"os"
)
func main() {
	oldFile := os.Args[1]

	if len(os.Args) != 3 {
		fmt.Print("Invalid length of arguments")
		return
	}

	_, err := os.ReadFile(oldFile)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	content, _ := os.ReadFile(oldFile)
	oldContent := string(content)

	functions.HexToDec(oldContent)
}