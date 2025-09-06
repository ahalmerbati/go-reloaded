package functions

import (
	"fmt"
	"strings"
)

func ValidateAlphaNumerical(content string) bool {
	content = strings.ReplaceAll(content, "\n", " ")
	content = strings.ReplaceAll(content, "\t", " ")
	
	arr := []rune(content)

	for i := 0; i < len(arr); i++ {
		if (arr[i] < 32) && (arr[i] > 126) {
			fmt.Println("The content includes non-alphanumerical characters")
			return false
		}
	}
	return true
}
