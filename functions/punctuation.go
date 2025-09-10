package functions

import (
	"fmt"
	"regexp"
	"strings"
)

func Punctuation(content string) string {
	expression1, err := regexp.Compile(`\s*([.,!?:;]+)`)
	if err != nil {
		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
		return content
	}

	content = expression1.ReplaceAllString(content, "$1")

	expression2, err := regexp.Compile(`([.,!?:;])\s*([^.,!?;:\s])`)
	if err != nil {
		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
		return content
	}
	content = expression2.ReplaceAllString(content, "$1 $2")

	expression3, err := regexp.Compile(`([.,!?:;]+)(\s+)([.,!?:;]+)`)
	if err != nil {
		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
		return content
	}
	content = expression3.ReplaceAllString(content, "$1$3")

	return strings.TrimSpace(content)
}
