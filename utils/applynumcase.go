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

	for j, word := range words {
		if j >= i {
			switch caseType {
			case "up":
				wordsAfter[j] = strings.ToUpper(word)

			case "low":
				wordsAfter[j] = strings.ToLower(word)

			case "cap":
				wordsAfter[j] = Capitalize(word)

			default:
				wordsAfter[j] = word
			}
		} else {
			wordsAfter[j] = word
		}
	}
	return wordsAfter
}
