package main

import (
	"fmt"
	"go-reloaded/functions"
	"strings"
)

func main() {
	// A list of test cases to run
	testCases := []string{
		"this is so exciting (up, 2)",
		"hello world, this is a test (low, 3)",
		"a new sentence is capitalized (cap, 4)",
		"multiple modifiers in one sentence (up, 1) and (low, 2) here",
		"what if there are no modifiers here",
		"This is a test case with a large number (up, 100)",
	}

	fmt.Println("--- Running Numbered Case Converter Tests ---")
	for _, tc := range testCases {
		result := functions.NumberedCaseConverter(tc)
		fmt.Printf("Input:  %s\n", tc)
		fmt.Printf("Output: %s\n", result)
		fmt.Println("-------------------------------------------")
	}

	fmt.Println("\n--- Running Single Word Converter Tests ---")
	testCases2 := []string{
		"this is so exciting (up)",
		"hello world, this is a test (low)",
		"a new sentence is capitalized (cap)",
	}

	for _, tc := range testCases2 {
		var result string
		if strings.Contains(tc, "(up)") {
			result = functions.UpperCaseConverter(tc)
		} else if strings.Contains(tc, "(low)") {
			result = functions.LowerCaseConverter(tc)
		} else if strings.Contains(tc, "(cap)") {
			result = functions.CapitalizedCaseConverter(tc)
		}
		fmt.Printf("Input:  %s\n", tc)
		fmt.Printf("Output: %s\n", result)
		fmt.Println("-------------------------------------------")
	}
}

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	if len(os.Args) != 3 {
// 		fmt.Println("invalid length of arguments")
// 		return
// 	}

// 	oldFile := os.Args[1]
// 	//newFile := os.Args[2]

// 	_, err := os.ReadFile(oldFile)
// 	if err != nil {
// 		fmt.Println("Error reading file")
// 		return
// 	}
// }