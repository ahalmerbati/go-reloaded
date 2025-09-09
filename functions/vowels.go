package functions

import (
	"fmt"
	"regexp"
	"strings"
)

func Vowels(content string) string {
	expression, err := regexp.Compile(`\b([aA])\b\s+([aAeEiIoOuUhH])`)
	if err != nil {
		fmt.Println("the regex pattern is invalid. the original content is returned.")
		return content
	}

	var result strings.Builder
	index := 0
	found := expression.FindAllStringSubmatchIndex(content, -1)

	if len(found) == 0 {
		fmt.Println("No instances of the letter a followed by a word starting with a vowel or h found in the content. The original content is returned:")
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
