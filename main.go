package main

import (
	"fmt"
	"go-reloaded/functions"
	
)

func main() {
	testCases := []string{
		"There it was. A amazing rock!",
		"Look its a amazing rock.",
		"A intelligent warrior has arrived.",
		"He saw a eagle and a horse.",
		"An apple a day keeps the doctor away.", // Should not change "An"
		"A unique situation.",
		"a", // Edge case: "a" not followed by word
		"A", // Edge case: "A" not followed by word
		"about", // Edge case: "a" within a word
		"a a a", // Multiple occurrences
		"A A A", // Multiple occurrences with caps
		"a amazing A hour",
	}

	for _, tc := range testCases {
		fmt.Printf("Original: \"%s\"\n", tc)
		fmt.Printf("Modified: \"%s\"\n\n", functions.Vowels(tc))
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