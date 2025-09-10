package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func ReplaceLastWords(content string, newWords []string, count int) string {
	expressionWord, err := regexp.Compile(`(\b[a-zA-Z0-9]+\b)`)
	if err != nil {
		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
		return content
	}

	wordsFound := expressionWord.FindAllStringIndex(content, -1)

	if len(wordsFound) == 0 || len(wordsFound) < count {
		fmt.Print("Not enough words were found to apply the case conversion. The original content is returned:")
		return content
	}

	startIndex := wordsFound[len(wordsFound)-count][0]
	endIndex := wordsFound[len(wordsFound)-1][1]

	var result strings.Builder
	result.WriteString(content[:startIndex])

	result.WriteString(strings.Join(newWords[len(newWords)-count:], " "))

	result.WriteString(content[endIndex:])

	return result.String()
}
