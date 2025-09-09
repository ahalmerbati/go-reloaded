package functions

import (
	"fmt"
	"regexp"
	"strings"
)

func Punctuation(content string) string {
	expression, err := regexp.Compile(`\s*([.,!?;:])\s*`)
	if err != nil {
		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
		return content
	}

	var result strings.Builder
	index := 0

	found := expression.FindAllStringSubmatchIndex(content, -1)

	if len(found) == 0 {
		fmt.Println("No instances of punctuation marks were found in the content. The original content is returned:")
		return content
	}

	for _, match := range found {
		result.WriteString(content[index:match[0]])
		result.WriteString(content[match[2]:match[3]])
		result.WriteString(" ")
		index = match[1]
	}

	result.WriteString(content[index:])

	expression2, err := regexp.Compile(`([.,!?;:])\s+`)
	if err != nil {
		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
		return content
	}

	finalString := expression2.ReplaceAllString(result.String(), "$1 ")

	return strings.TrimSpace(finalString)
}

func GroupPunctuation(content string) string {

}
