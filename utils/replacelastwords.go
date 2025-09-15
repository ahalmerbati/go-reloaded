package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// The function is called to replace the last (count) words in a string with the words from the newWords slice
func ReplaceLastWords(content string, newWords []string, count int) string {
	expressionWord, err := regexp.Compile(`(\b[a-zA-Z0-9]+\b)`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}

	wordsFound := expressionWord.FindAllStringIndex(content, -1)

	if len(wordsFound) == 0 || len(wordsFound) < count {
		fmt.Println("Error: Not enough words were found to apply the case conversion.")
		return content
	}

	startIndex := wordsFound[len(wordsFound)-count][0]
	endIndex := wordsFound[len(wordsFound)-1][1]

	var result strings.Builder
	result.WriteString(content[:startIndex])

	result.WriteString(strings.Join(newWords, " "))

	result.WriteString(content[endIndex:])

	return result.String()
}
