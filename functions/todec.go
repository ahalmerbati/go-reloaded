package functions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// The function converts hexadecimal and binary numbers within a string to their decimal equivalents
func ToDec(content string) string {
	validCmd, err := regexp.Compile(`(\b(?:[0-9a-zA-Z]+|[01]+)\b)([\s.,!?:;]*)\((hex|bin)\)`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}

	multiCommandExpression, err := regexp.Compile(`\s*\((hex|bin)\)(\s*\((hex|bin)\)){1,}`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}
	if multiCommandExpression.MatchString(content) {
		fmt.Println("Found multiple commands following a single number. Only the first command will be applied.")
	}

	invalidCmd, err := regexp.Compile(`\(\s*\(+\s*(bin|hex)\s*\)+\s*\)|\)+\s*\((bin|hex)\)\)+|\(\s*(bin|hex)\s*\)+`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}

	if invalidCmd.MatchString(content) {
		fmt.Println("Error: The command format is invalid. Command must be in the format (hex) or (bin) with no extra parentheses.")
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

		oldNum := content[match[2]:match[3]]
		punctuation := content[match[4]:match[5]]
		command := content[match[6]:match[7]]

		var base int
		switch command {
		case "hex":
			base = 16
		case "bin":
			base = 2
		}

		decNum, err := strconv.ParseInt(oldNum, base, 64)
		if err != nil {
			fmt.Printf("Error: Could not convert %s to a decimal number.\n", oldNum)
			result.WriteString(oldNum)
			result.WriteString(punctuation)
		} else {
			result.WriteString(fmt.Sprintf("%d", decNum))
			result.WriteString(punctuation)
		}
		index = match[1]
	}
	result.WriteString(content[index:])

	return result.String()
}
