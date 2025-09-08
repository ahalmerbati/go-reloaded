package functions

import (
	"fmt"
	"go-reloaded/utils"
	"regexp"
	"strconv"
	"strings"
)

func NumberedCaseConverter(content string) string {
	expression, err := regexp.Compile(`[\s.,!?:;]*\((up|low|cap),\s*(\d+)\)`)
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
		wordsBefore := expressionWordsBefore.FindAllString(content[:match[0]], -1)

		wordsAfter := utils.ApplyNumCase(wordsBefore, caseType, count)

		result.WriteString(strings.Join(wordsAfter, " "))

		index = match[1]
	}
	result.WriteString(content[index:])

	return result.String()
}
