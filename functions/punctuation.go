package functions

import (
	"fmt"
	"regexp"
	"strings"
)

func Punctuation(content string) string {
	expression, err := regexp.Compile(`\s*([.,!?:;]+)\s*`)
	if err != nil {
		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
		return content
	}

	var result strings.Builder
	index := 0

	found := expression.FindAllStringSubmatchIndex(content, -1)

	if len(found) == 0 {
		fmt.Println("No instances of puncutation marks were found in the content. The original content is returned:")
		return content
	}

	for _, match := range found {
		result.WriteString(content[index:match[0]])

		result.WriteString(content[match[2]:match[3]])

		if match[1] < len(content) {
			result.WriteString(" ")
		}

		index = match[1]
	}

	result.WriteString(content[index:])

	finalString := strings.TrimSpace(result.String())

	return finalString
}
