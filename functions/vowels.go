package functions

import (
	"fmt"
	"regexp"
	"strings"
)

// The function replaces any instance of "a" with "an" when it is followed by a word that starts with a vowel or 'h'
func Vowels(content string) string {
	expression, err := regexp.Compile(`\b([aA])\b((?:\s*\([a-z]+\))*)\s+([aAeEiIoOuUhH][a-zA-Z]*)`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}

	var result strings.Builder
	index := 0
	found := expression.FindAllStringSubmatchIndex(content, -1)

	if len(found) == 0 {
		return content
	}

	for _, match := range found {
		result.WriteString(content[index:match[0]])

		letterA := content[match[2]:match[3]]
		replacement := "an"

		if letterA == "A" {
			replacement = "An"
		}

		result.WriteString(replacement)
		result.WriteString(content[match[3]:match[1]])

		index = match[1]
	}

	result.WriteString(content[index:])

	return result.String()
}
