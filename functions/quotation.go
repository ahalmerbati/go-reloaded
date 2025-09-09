package functions

import (
	"fmt"
	"regexp"
	"strings"
)

func Quotation(content string) string {
	expression, err := regexp.Compile(`'`)
	if err != nil {
		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
		return content
	}

	var result strings.Builder
	index := 0

	found := expression.FindAllStringSubmatchIndex(content, -1)

	if len(found) == 0 {
		fmt.Println("No instances of ' were found in the content. The original content is returned:")
		return content
	}

	if len(found)%2 != 0 {
		fmt.Println("Odd number of quotation marks found. The original content is returned:")
		return content
	}

	for i := 0; i < len(found); i = i + 2 {
		firstQuote := found[i][0]
		secondQuote := found[i+1][0]

		result.WriteString(content[index:firstQuote])
		result.WriteString(string(content[firstQuote]))

		contentBetweenQuotes := strings.TrimSpace(content[firstQuote+1 : secondQuote])

		result.WriteString(contentBetweenQuotes)

		result.WriteString(string(content[secondQuote]))

		if secondQuote+1 < len(content) && content[secondQuote+1] != ' ' {
			result.WriteString(" ")
		}

		index = secondQuote + 1
	}

	result.WriteString(content[index:])

	return result.String()
}
