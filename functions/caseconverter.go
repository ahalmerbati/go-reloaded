package functions

import (
	"fmt"
	"go-reloaded/utils"
	"regexp"
	"strings"
)

func CaseConverter(content string) string {
	multiCommandExpression, err := regexp.Compile(`\s*\((up|low|cap)\)(\s*\((up|low|cap)\)){1,}`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}
	if multiCommandExpression.MatchString(content) {
		fmt.Println("Found multiple commands following a single word. Only the first command will be applied.")
	}

	expression, err := regexp.Compile(`(\b[a-zA-Z]+\b)[\s.,!?:;]*\((up|low|cap)\)`)
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

		word := content[match[2]:match[3]]

		caseType := content[match[4]:match[5]]

		convertedWord := utils.ApplySingleCase(word, caseType)

		result.WriteString(convertedWord)

		index = match[1]
	}
	result.WriteString(content[index:])

	return result.String()
}
