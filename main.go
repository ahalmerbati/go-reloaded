package main

import (
	"fmt"
	"go-reloaded/functions"
	"os"
	"strings"
)

func main() {
	sampleFile := os.Args[1]
	resultFile := os.Args[2]

	if len(os.Args) != 3 {
		fmt.Println("Error: Invalid length of arguments")
		return
	}

	if !strings.HasSuffix(sampleFile, ".txt") || !strings.HasSuffix(resultFile, ".txt") {
		fmt.Println("Error: The file to be read from and the file to be written to have to both be a txt file")
		return
	}

	content, err := os.ReadFile(sampleFile)
	if err != nil {
		fmt.Println("Error: Could not read file")
		return
	}

	if !functions.CheckAsciiAndPrintables(string(content)) {
		fmt.Println("Error: Cannot process content as it contains non-ASCII or non-printable characters.")
		return
	}

	if !functions.CheckFileSize(sampleFile) {
		return
	}

	finalContent := functions.ProcessText(string(content))

	err = os.WriteFile(resultFile, []byte(finalContent), 0644)
	if err != nil {
		fmt.Println("Error: Could not write into file")
		return
	}

	fmt.Println("Successfully proccessed sample.txt and wrote the output to result.txt")
}

// if theres a (  so like hello )(low)( it should remove the bracket