package functions

import (
	"fmt"
	"regexp"
)

func ErrorHandling(content string) bool {
	multiCommandSingle, err := regexp.Compile(`\s*\((up|low|cap)\)(\s*\((up|low|cap)\)){1,}`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return false
	}
	if multiCommandSingle.MatchString(content) {
		fmt.Println("Error: Found multiple commands following a single word. Only the first command will be applied.")
		return false
	}

	invalidCmdSingle, err := regexp.Compile(`\(\s*\(+\s*(up|low|cap)\s*\)+\s*\)|\)\s*\((up|low|cap)\)\(`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return false
	}

	if invalidCmdSingle.MatchString(content) {
		fmt.Println("Error: The command format is invalid. Command must be in the format (up), (low), or (cap) with no extra parentheses.")
		return false
	}

	invalidCmdNumbered, err := regexp.Compile(`\(\s*\(+\s*(up|low|cap)\s*,\s*-?\d+\s*\)+\s*\)|\)\s*\((up|low|cap),\s*-?\d+\)\(`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return false
	}

	if invalidCmdNumbered.MatchString(content) {
		fmt.Println("Error: The command format is invalid. Command must be in the format (up, <number>), (low, <number>), or (cap, <number>) with no extra parentheses.")
		return false
	}

	return true
}
