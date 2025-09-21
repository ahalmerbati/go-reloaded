package main

import (
	"fmt"
	"go-reloaded/functions"
	"go-reloaded/utils"
	"log"
	"os"
)

func main() {
	utils.ValidateInputFiles()

	sampleFile := os.Args[1]
	resultFile := os.Args[2]

	content, err := os.ReadFile(sampleFile)
	if err != nil {
		log.Fatal("Error: Could not read the file")
	}

	finalContent := functions.ProcessText(string(content))

	err = os.WriteFile(resultFile, []byte(finalContent), 0644)
	if err != nil {
		log.Fatal("Error: Could not write into file")
	}

	fmt.Println("Successfully processed sample.txt and wrote the output to result.txt")
}
