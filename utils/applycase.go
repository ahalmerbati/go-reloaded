package utils

import (
	"strings"
)

func ApplyNumCase(words []string, caseType string, count int) []string {
	wordsAfter := make([]string, len(words))

	i := len(words) - count
	if i < 0 {
		i = 0
	}

	for j := i; j < len(wordsAfter); j++ {
		switch caseType {
		case "up":
			wordsAfter[j] = strings.ToUpper(wordsAfter[j])
		case "low":
			wordsAfter[j] = strings.ToLower(wordsAfter[j])
		case "cap":
			wordsAfter[j] = Capitalize(wordsAfter[j])
		}
	}
	return wordsAfter
}

func ApplySingleCase(word string, caseType string) string {
	var convertedWord string

	switch caseType {
	case "up":
		convertedWord = strings.ToUpper(word)
	case "low":
		convertedWord = strings.ToLower(word)
	case "cap":
		convertedWord = Capitalize(word)
	}

	return convertedWord
}
