package functions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// The function converts hexadecimal and binary numbers within a string to their decimal equivalents
func ToDec(content string) string {
	multiCommandExpression, err := regexp.Compile(`\s*\((hex|bin)\)(\s*\((hex|bin)\)){1,}`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}
	if multiCommandExpression.MatchString(content) {
		fmt.Println("Found multiple commands following a single number. Only the first command will be applied.")
	}

	expression, err := regexp.Compile(`(\b(?:[0-9a-fA-F]+|[01]+)\b)[\s.,!?:;]*\((hex|bin)\)`)
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

		oldNum := content[match[2]:match[3]]
		command := content[match[4]:match[5]]

		var base int
		switch command {
		case "hex":
			base = 16
		case "bin":
			base = 2
		}

		decNum, err := strconv.ParseInt(oldNum, base, 64)
		if err != nil {
			result.WriteString(content[match[0]:match[1]])
		} else {
			result.WriteString(fmt.Sprintf("%d", decNum))
		}
		index = match[1]
	}
	result.WriteString(content[index:])

	return result.String()
}
