package functions

import (
	"fmt"
	"go-reloaded/utils"
	"regexp"
	"strings"
)

// The function processes a string to modify text enclosed within single quotation marks
func Quotation(content string) string {
	expression, err := regexp.Compile(`'`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}

	var quotationPosition []int
	found := expression.FindAllStringIndex(content, -1)
	for _, match := range found {
		i := match[0]
		if !utils.IsApostrophe(content, i) {
			quotationPosition = append(quotationPosition, i)
		}
	}

	if len(quotationPosition) == 0 {
		return content
	}

	if len(quotationPosition)%2 != 0 {
		fmt.Println("Error: Odd number of quotation marks found.")
		return content
	}

	var result strings.Builder
	index := 0

	for i := 0; i < len(quotationPosition); i = i + 2 {
		firstQuote := quotationPosition[i]
		secondQuote := quotationPosition[i+1]

		result.WriteString(content[index:firstQuote])
		result.WriteString("'")

		contentBetweenQuotes := strings.TrimSpace(content[firstQuote+1 : secondQuote])

		result.WriteString(contentBetweenQuotes)
		result.WriteString("'")

		if secondQuote+1 < len(content) && content[secondQuote+1] != ' ' {
			result.WriteString(" ")
		}

		index = secondQuote + 1
	}

	result.WriteString(content[index:])

	return result.String()
}
