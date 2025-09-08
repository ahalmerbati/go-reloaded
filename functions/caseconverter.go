package functions

import (
	"fmt"
	"go-reloaded/utils"
	"regexp"
	"strings"
)

func CaseConverter(content, pattern string, converter func(string) string) string {
	expression, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("the regex pattern is invalid. the original content is returned.")
		return content
	}

	var result strings.Builder
	index := 0
	found := expression.FindAllStringSubmatchIndex(content, -1)

	if len(found) == 0 {
		fmt.Println("No instances of '(up)', '(low)', or '(cap)' found in the content. The original content is returned:")
		return content
	}

	for _, match := range found {
		result.WriteString(content[index:match[0]])

		word := content[match[2]:match[3]]

		convertedWord := converter(word)

		result.WriteString(convertedWord)

		index = match[1]
	}
	result.WriteString(content[index:])

	if strings.Contains(result.String(), strings.Split(pattern, `\s*`)[1]) {
		fmt.Println("Some words were not fully converted due to invalid input or other ASCII characters.")
	}

	return result.String()
}

func UpperCaseConverter(content string) string {
	return CaseConverter(content, `(\b[a-zA-Z]+\b)\s*\(up\)`, strings.ToUpper)
}

func LowerCaseConverter(content string) string {
	return CaseConverter(content, `(\b[a-zA-Z]+\b)\s*\(low\)`, strings.ToLower)
}

func CapitalizedCaseConverter(content string) string {
	return CaseConverter(content, `(\b[a-zA-Z]+\b)\s*\(cap\)`, utils.Capitalize)
}
