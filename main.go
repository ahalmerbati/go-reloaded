package main

import (
	"fmt"
	"go-reloaded/functions"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("invalid length of arguments")
		return
	}

	sampleFile := os.Args[1]
	resultFile := os.Args[2]

	content, err := os.ReadFile(sampleFile)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	finalContent := functions.ProcessText(string(content))

	err = os.WriteFile(resultFile, []byte(finalContent), 0644)
	if err != nil {
		fmt.Println("Error writing to file")
		return
	}

	fmt.Println("Successfully proccessed sample.txt and wrote the output to result.txt")
}
// Ask Ridha where to do the if statment for if its non-ascii characters / non-printables