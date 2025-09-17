package functions

import (
	"fmt"
	"go-reloaded/utils"
	"strings"
)

func CaseConverter(content string) string {
	if !ErrorHandling(content) {
		return content
	}

	words := strings.Fields(content)
	result := make([]string, 0)
	index := 0

	for index < len(words) {
		char := words[index]

		if strings.HasPrefix(char, "(") && strings.HasSuffix(char, ")") && CheckCommand(char) {
			command := strings.Trim(char, "()")
			if strings.Contains(command, ",") {
				seg := strings.SplitN(command, ",", 2)
				if len(seg) == 2 {
					commandToApply := strings.TrimSpace(seg[0])
					numStr := strings.TrimSpace(seg[1])

				}
			} else {
				utils.ApplySingleCase()
			}
		}

	}
}

// } else {
	// 	numStr := content[match[4]:match[5]]
	// 	count, err := strconv.Atoi(numStr)
	// 	if err != nil {
		// 		fmt.Println("Error: Could not convert the string to a number.")
		// 		// rewrite the preceding content and continue
		// 	}
		// }
		
		// cmdExpression, err := regexp.Compile(`\(\s*(up|low|cap)\s*(,\s*(\d+)\s*)?\)`)
		// if err != nil {
		// 	fmt.Println("Error: Could not compile the regular expression.")
		// 	return content
		// }
		// wordExpression, err := regexp.Compile(`(\b[A-Za-z0-9]+\b)`)
		// if err != nil {
		// 	fmt.Println("Error: Could not compile the regular expression.")
		// 	return content
		// }
		
		// matches := cmdExpression.FindAllStringSubmatchIndex(content, -1)
		// if len(matches) == 0 {
		// 	fmt.Println("Error: No valid commands found in the content.")
		// 	return content
		// }
		
		// var result strings.Builder
		// index := 0
		
		// for _, match := range matches {
		// 	wordsToApply := wordExpression.FindAllString(content[:match[0]], -1)
		
		// 	command := content[match[2]:match[3]]
		
		// 	if match[4] == -1 {
		// 		words := content[index:match[2]]
		
		// 		wordsToApply = append(wordsToApply, words)
		
		// 		lastWord := wordsToApply[len(wordsToApply)-1]
		
		// 		var modifiedWords []string
		
		// 		modifiedWord := utils.ApplySingleCase(lastWord, command)
		
		// 		modifiedWords = append(modifiedWords, modifiedWord)
		
		// 		modifiedContent := utils.ReplaceLastWords(content[:match[0]], modifiedWords, 1)
		
		// 		result.WriteString(modifiedContent)
		
		// 		index = match[1]
		// 	}
		// }
		// result.WriteString(content[index:])
		// return result.String()