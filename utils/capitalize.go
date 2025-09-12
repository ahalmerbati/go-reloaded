package utils

import (
	"fmt"
	"strings"
)

// The function is called to capitalize the first letter of a string and turn the rest of the word to lowercase.
func Capitalize(str string) string {
	if len(str) == 0 {
		fmt.Println("Error: Cannot capitalize an empty string")
		return str
	}
	return strings.ToUpper(str[:1]) + strings.ToLower(str[1:])
}
