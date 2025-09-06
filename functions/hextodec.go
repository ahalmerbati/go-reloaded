package functions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func HexToDec(content string) string {
	expression, err := regexp.Compile(`(\b[0-9a-fA-F]+\b)\s*\(hex\)`)
	if err != nil {
		return "invalid regex pattern"
	}

	var result strings.Builder
	index := 0

	found := expression.FindAllStringSubmatchIndex(content, -1)

	if found == nil {
		return content
	}

	for _, match := range found {
		result.WriteString(content[index:match[0]])

		hexNum := content[match[2]:match[3]]

		decNum, err := strconv.ParseInt(hexNum, 16, 64)
		if err != nil {
			result.WriteString(content[match[0]:match[1]])
		} else {
			result.WriteString(fmt.Sprintf("%d", decNum))
		}
		index = match[1]
	}
	result.WriteString(content[index:])

	if strings.Contains(result.String(), "(hex)") {
		fmt.Println("Some hexadecimal numbers were not fully converted due to invalid input or other ASCII characters.") 
	}

	return result.String()
}
