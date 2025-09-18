package utils

import (
	"fmt"
	"strings"
)

// The function applies a specified case conversion to the last (count) words in a slice of strings
func ApplyNumCase(words []string, caseType string, count int) {

	startIndex := len(words) - count
	if startIndex < 0 {
		startIndex = 0
	}

	for i := startIndex; i < len(words); i++ {
		switch caseType {
		case "up":
			words[i] = strings.ToUpper(words[i])
		case "low":
			words[i] = strings.ToLower(words[i])
		case "cap":
			words[i] = Capitalize(words[i])
		}
	}
}

// The function applies a specificed case converstion to a single word
func ApplySingleCase(words []string, caseType string) {
	if len(words) == 0 {
		fmt.Println("Error: No word was found preceding the command. Cannot apply single case command.")
		return
	}

	index := len(words) - 1

	switch caseType {
	case "up":
		words[index] = strings.ToUpper(words[index])
	case "low":
		words[index] = strings.ToLower(words[index])
	case "cap":
		words[index] = Capitalize(words[index])
	}
}
