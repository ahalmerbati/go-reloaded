package functions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func BinToDec(content string) string {
	expression, err := regexp.Compile(`(\b[01]+\b)\s*\(bin\)`)
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

		binNum := content[match[2]:match[3]]

		decNum, err := strconv.ParseInt(binNum, 2, 64)
		if err != nil {
			result.WriteString(content[match[0]:match[1]])
		} else {
			result.WriteString(fmt.Sprintf("%d", decNum))
		}

		index = match[1]
	}
	result.WriteString(content[index:])

	if strings.Contains(result.String(), "(bin)") {
		fmt.Println("Some binary numbers were not fully converted due to invalid input or other ASCII characters.")
	}

	return result.String() 
}
