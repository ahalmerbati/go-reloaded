package functions

import (
	"fmt"
	"go-reloaded/utils"
	"regexp"
	"strings"
)

// The function finds specific commands in a string and applies case conversion to the word immediately preceding the command
func SingleCaseConverter(content string) string {
	validCmd, err := regexp.Compile(`(?i)(\b[a-zA-Z0-9]+\b)([.,!?:;]*)\s*\((up|low|cap)\)`)

	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}

	var result strings.Builder
	index := 0
	found := validCmd.FindAllStringSubmatchIndex(content, -1)

	if len(found) == 0 {
		return content
	}

	if !ErrorHandling(content) {
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
