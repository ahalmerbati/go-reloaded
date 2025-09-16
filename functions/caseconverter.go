package functions

import (
	"fmt"
	"go-reloaded/utils"
	"regexp"
	"strings"
)

func CaseConverter(content string) string {
	if !ErrorHandling(content) {
		return content
	}

	cmdExpression, err := regexp.Compile(`\(\s*(up|low|cap)\s*(,\s*(\d+)\s*)?\)`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}
	wordExpression, err := regexp.Compile(`(\b[A-Za-z0-9]+\b)`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}

	matches := cmdExpression.FindAllStringSubmatchIndex(content, -1)
	wordsToApply := wordExpression.FindAllString(content, -1)

	var result strings.Builder
	index := 0

	for _, match := range matches {
		command := content[match[2]:match[3]]

		if match[4] == -1 {
			words := content[index:match[2]]
			wordsToApply = append(wordsToApply, words)
			lastWord := wordsToApply[len(wordsToApply)-1]
			modifiedContent := utils.ApplySingleCase(lastWord, command)
			result.WriteString(content[index:])
			result.WriteString(modifiedContent)
			index = match[1]
		}
	}
	result.WriteString(content[index:])
	return result.String()
}

// } else {
// 	numStr := content[match[4]:match[5]]
// 	count, err := strconv.Atoi(numStr)
// 	if err != nil {
// 		fmt.Println("Error: Could not convert the string to a number.")
// 		// rewrite the preceding content and continue
// 	}
// }
