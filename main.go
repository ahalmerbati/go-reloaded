package main

import (
	"fmt"
	"go-reloaded/functions"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: Invalid length of arguments")
		return
	}

	sampleFile := os.Args[1]
	resultFile := os.Args[2]

	if !strings.HasSuffix(sampleFile, ".txt") || !strings.HasSuffix(resultFile, ".txt") {
		fmt.Println("Error: The file to be read from and the file to be written to have to both be a txt file")
		return
	}

	if sampleFile == resultFile {
		fmt.Println("The input file cannot be the same as the output file")
		return
	}

	if !functions.CheckFileSize(sampleFile) {
		fmt.Println("Error: The file size exceeds the maximum limit of 100KB")
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

	finalContent := functions.ProcessText(string(content))

	err = os.WriteFile(resultFile, []byte(finalContent), 0644)
	if err != nil {
		fmt.Println("Error: Could not write into file")
		return
	}

	fmt.Println("Successfully proccessed sample.txt and wrote the output to result.txt")
}
