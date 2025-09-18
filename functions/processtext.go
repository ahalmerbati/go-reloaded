package functions

import (
	"strings"
)

// The function is used for processing text by applying a series of transformations to the input string
func ProcessText(content string) string {
	if !ErrorHandling(content) {
		return content
	}

	//content = CheckCaseConverterSpacing(content)

	content = Vowels(content)
	content = ToDec(content)
	content = CaseConverter(content)
	content = Punctuation(content)
	content = Quotation(content)

	content = strings.TrimSpace(content)

	return content
}
