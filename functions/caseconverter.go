package functions

import (
	"fmt"
	"go-reloaded/utils"
	"regexp"
	"strconv"
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

func NumberedCaseConverter(content string) string {
	expression, err := regexp.Compile(`\s*\((up|low|cap),\s*(\d+)\)`)
	if err != nil {
		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
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
		caseType := content[match[2]:match[3]]
		numStr := content[match[4]:match[5]]

		count, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error converting string to a number. The original content is returned:")
			return content
		}

		expressionWordsBefore, err := regexp.Compile(`(\b[a-zA-Z]+\b)`)
		if err != nil {
			fmt.Println("There was an error compiling the regex expression. The original content is returned.")
			return content
		}
		wordsBefore := expressionWordsBefore.FindAllStringSubmatchIndex(content[:match[0]], -1)

		wordsAfter := utils.ApplyNumCase(wordsBefore, caseType, count)

		result.WriteString(strings.Join(wordsAfter, " "))

		index = match[1]
	}
	result.WriteString(content[index:])

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
