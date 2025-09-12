package functions

import (
	"fmt"
	"go-reloaded/utils"
	"regexp"
	"strconv"
	"strings"
)

// The function finds specific commands in a string and applies case conversion to a given number of preceding words
func NumberedCaseConverter(content string) string {
	validCmd, err := regexp.Compile(`(?i)([.,!?:;]*)\s*\((up|low|cap),\s*(-?\d+)\)`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}

	invalidCmd, err := regexp.Compile(`\(\s*\(+\s*(up|low|cap)\s*,\s*-?\d+\s*\)+\s*\)|\)\s*\((up|low|cap),\s*-?\d+\)\(`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}

	if invalidCmd.MatchString(content) {
		fmt.Println("Error: The command format is invalid. Command must be in the format (up, <number>), (low, <number>), or (cap, <number>) with no extra parentheses.")
		return content
	}

	var result strings.Builder
	index := 0
	found := validCmd.FindAllStringSubmatchIndex(content, -1)

	if len(found) == 0 {
		return content
	}

	const maxWords = 1000
	flag := false

	for i, match := range found {
		precedingContent := content[index:match[0]]

		if i > 0 && strings.TrimSpace(content[found[i-1][1]:match[0]]) == "" {
			if !flag {
				fmt.Println("Error: Found multiple commands following a single word. Only the first command will be applied.")
				flag = true
			}
			result.WriteString(content[match[0]:match[1]])
			index = match[1]
			continue
		} else {
			flag = false
		}

		expression2, err := regexp.Compile(`(\b[a-zA-Z0-9]+\b)`)
		if err != nil {
			fmt.Println("Error: Could not compile the regular expression.")
			return content
		}
		words := expression2.FindAllString(precedingContent, -1)

		punctuation := content[match[2]:match[3]]
		caseType := content[match[4]:match[5]]
		numStr := content[match[6]:match[7]]

		count, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error: Could not convert the string to a number.")
			result.WriteString(precedingContent)
			result.WriteString(content[match[0]:match[1]])
			index = match[1]
			continue
		}

		if count > maxWords {
			fmt.Printf("Error: The number of words to modify (%d) exceeds the limit of %d.\n", count, maxWords)
			result.WriteString(precedingContent)
			result.WriteString(content[match[0]:match[1]])
			index = match[1]
			continue
		}

		if count <= 0 {
			fmt.Println("Error: The number of words to modify cannot be equal to or less than 0.")
			result.WriteString(precedingContent)
			result.WriteString(content[match[0]:match[1]])
			index = match[1]
			continue
		}

		if len(words) < count {
			fmt.Println("Error: Not enough words were found to apply the case conversion.")
			result.WriteString(precedingContent)
			result.WriteString(content[match[0]:match[1]])
			index = match[1]
			continue
		}

		modifiedWords := utils.ApplyNumCase(words, caseType, count)
		modifiedContent := utils.ReplaceLastWords(precedingContent, modifiedWords, count)

		result.WriteString(modifiedContent)
		result.WriteString(punctuation)
		index = match[1]
	}
	result.WriteString(content[index:])
	return result.String()
}
