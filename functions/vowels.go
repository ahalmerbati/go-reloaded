package functions

import (
	"strings"
)

func Vowels(content string) string {
	vowels := "aeiouhAEIOUH"
	words := strings.Fields(content)
	resultWords := make([]string, 0)

	for i := 0; i < len(words); i++ {
		if words[i] == "a" || words[i] == "A" || words[i] == "an" || words[i] == "An" {
			if i+1 < len(words) && strings.ContainsAny(string(words[i+1][0]), vowels) {
				if words[i] == "a" {
					resultWords = append(resultWords, "an")
				} else if words[i] == "A" {
					resultWords = append(resultWords, "An")
				} else if words[i] == "An" {
					resultWords = append(resultWords, "A")
				} else if words[i] == "an" {
					resultWords = append(resultWords, "a")
				}
				continue
			}
		}
		resultWords = append(resultWords, words[i])

	}
	return strings.Join(resultWords, " ")
}
