package main

import (
	"fmt"
	"go-reloaded/functions"
)

func main() {
	testCases := []string{
		"Punctuation tests are ... kinda boring ,what do you think ?",
	}

	
	for _, tc := range testCases {
		fmt.Printf("Original: \"%s\"\n", tc)

		x := functions.Punctuation(tc)
		
		fmt.Printf("Processed: \"%s\"\n\n", x)
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

// 	sampleFile := os.Args[1]
// 	//resultFile := os.Args[2]

// 	_, err := os.ReadFile(oldFile)
// 	if err != nil {
// 		fmt.Println("Error reading file")
// 		return
// 	}
// }
