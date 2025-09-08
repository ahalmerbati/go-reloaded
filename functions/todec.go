package functions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ToDec(content, pattern string, base int) string {
	expression, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("The regex pattern is invalid. The original content is returned.")
		return content
	}

	var result strings.Builder
	index := 0
	found := expression.FindAllStringSubmatchIndex(content, -1)

	if len(found) == 0 {
		fmt.Println("No instances of '(hex)' or '(bin)' found in the content. The original content is returned:")
		return content
	}

	for _, match := range found {
		result.WriteString(content[index:match[0]])

		oldNum := content[match[2]:match[3]]

		decNum, err := strconv.ParseInt(oldNum, base, 64)
		if err != nil {
			result.WriteString(content[match[0]:match[1]])
		} else {
			result.WriteString(fmt.Sprintf("%d", decNum))
		}
		index = match[1]
	}
	result.WriteString(content[index:])

	if strings.Contains(result.String(), strings.Split(pattern, `\s*`)[1]) {
		fmt.Println("Some words were not fully converted due to invalid input or other ASCII characters.")
	}

	return result.String()
}

func HexToDec(content string) string {
	return ToDec(content, `(\b[0-9a-fA-F]+\b)\s*\(hex\)`, 16)
}

func BinToDec(content string) string {
	return ToDec(content, `(\b[01]+\b)\s*\(bin\)`, 2)
}
