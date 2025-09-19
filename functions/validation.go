package functions

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

// The function verifies that a given string contains only ASCII and printable characters
func CheckAsciiAndPrintables(content string) bool {
	content = strings.ReplaceAll(content, "\n", " ")
	content = strings.ReplaceAll(content, "\t", " ")

	expression, err := regexp.Compile(`\s+`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return false
	}
	content = expression.ReplaceAllString(content, " ")

	for _, r := range content {
		if r > unicode.MaxASCII || !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

// The function verifies if the file is within the maximum size limit of 100KB
func CheckFileSize(filePath string) bool {
	const maxSize int64 = 102400 // 100KB

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Error: Could not get file information")
		return false
	}

	if fileInfo.Size() > maxSize {
		fmt.Println("Error: The file size exceeds the maximum limit of 100KB")
		return false
	}

	return true
}

func CheckCommand(command string) bool {
	cmd := strings.Trim(command, "()")

	for _, v := range cmd {
		if !unicode.IsLetter(v) && !unicode.IsNumber(v) && v != ',' && v != ' ' && v != 10 {
			return false
		}
	}
	return true
}

func CheckCaseConverterSpacing(content string) string {
	arr := []rune(content)
	var result strings.Builder

	for i := 0; i < len(arr); i++ {
		if arr[i] != '(' {
			result.WriteRune(arr[i])
			continue
		}

		for _, r := range arr[i:] {
			if r == ')' {
				result.WriteRune(')')
				break
			}
			if r != ' ' {
				result.WriteRune(r)
			}
		}
	}

	return result.String()
}
