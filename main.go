package main

import (
	"fmt"
	"go-reloaded/functions"
	
)

func main() {
	testCases := []string{
		"There was an error   , compiling    ! the regular  , expression.   ",
	}

	for _, tc := range testCases {
		fmt.Printf("Original: \"%s\"\n", tc)
		fmt.Printf("Modified: \"%s\"\n\n", functions.Punctuation(tc))
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