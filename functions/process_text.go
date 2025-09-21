package functions

import (
	"go-reloaded/utils"
	"strings"
)

// The function is used for processing text by applying a series of transformations to the input string
func ProcessText(content string) string {
	content = Punctuation(content)

	if !ErrorHandling(content) {
		return content
	}

	content = utils.CheckCaseConverterSpacing(content)
	content = utils.CheckASCIIBrackets(content)

	content = Vowels(content)
	content = ToDec(content)
	content = CaseConverter(content)
	content = Punctuation(content)
	content = Quotation(content)
	content = Vowels(content)

	content = strings.TrimSpace(content)

	return content
}
