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
		fmt.Print("Invalid regex pattern")
		return ""
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
			fmt.Print("The hexadecimal number could not be converted to decimal")
		} else {
			result.WriteString(fmt.Sprintf("%d", decNum))
		}
		index = match[1]
	}
	result.WriteString(content[index:])

	return result.String()
}

// contentCheck := strings.Contains(content, "(hex)")

// if !(contentCheck) {
// 	return content
// }

// arr := []rune(content)

// if contentCheck {
// 	for i := 0; i < len(arr); i++ {
// 		if string(arr[i:i+5]) == "(hex)" {

// 		}
// 	}
// }
// return string(arr)
