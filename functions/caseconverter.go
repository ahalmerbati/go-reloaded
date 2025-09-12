package functions

import (
	"fmt"
	"go-reloaded/utils"
	"regexp"
	"strings"
)

// The function finds specific commands in a string and applies case conversion to the word immediately preceding the command
func CaseConverter(content string) string {
	validCmd, err := regexp.Compile(`(?i)(\b[a-zA-Z]+\b)([.,!?:;]*)\s*\((up|low|cap)\)`)

	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}

	multiCommandExpression, err := regexp.Compile(`\s*\((up|low|cap)\)(\s*\((up|low|cap)\)){1,}`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}
	if multiCommandExpression.MatchString(content) {
		fmt.Println("Error: Found multiple commands following a single word. Only the first command will be applied.")
	}

	invalidCmd, err := regexp.Compile(`\(\s*\(+\s*(up|low|cap)\s*\)+\s*\)|\)\s*\((up|low|cap)\)\(`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}

	if invalidCmd.MatchString(content) {
		fmt.Println("Error: The command format is invalid. Command must be in the format (up), (low), or (cap) with no extra parentheses.")
		return content
	}

	var result strings.Builder
	index := 0
	found := validCmd.FindAllStringSubmatchIndex(content, -1)

	if len(found) == 0 {
		return content
	}

	for _, match := range found {
		result.WriteString(content[index:match[0]])

		word := content[match[2]:match[3]]
		punctuation := content[match[4]:match[5]]
		caseType := content[match[6]:match[7]]

		if caseType != "up" && caseType != "low" && caseType != "cap" {
			fmt.Println("Error: Command must be lowercase and one of (up, low, cap).")
			result.WriteString(word + punctuation)
			index = match[1]
			continue
		}

		convertedWord := utils.ApplySingleCase(word, caseType)

		result.WriteString(convertedWord + punctuation)

		index = match[1]
	}
	result.WriteString(content[index:])

	return result.String()
}
