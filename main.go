package main

import (
	"fmt"
	"go-reloaded/functions"
	"os"
)
func main() {
	if len(os.Args) != 3 {
		fmt.Print("Invalid length of arguments")
		return
	}

	oldFile := os.Args[1]

	_, err := os.ReadFile(oldFile)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	content, _ := os.ReadFile(oldFile)
	oldContent := string(content)

	functions.TextEditor(oldContent)
	// functions.HexToDec(oldContent)
	fmt.Print(content) 
}