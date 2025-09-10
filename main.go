package main

import (
	"fmt"
	"go-reloaded/functions"
)

func main() {
	testCases := []string{
		"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.",
	}

	for _, tc := range testCases {
		fmt.Printf("Original: \"%s\"\n", tc)

		y := functions.UpperCaseConverter(tc)
		z := functions.LowerCaseConverter(y)
		r := functions.CapitalizedCaseConverter(z)
		x := functions.NumberedCaseConverter(r)

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
