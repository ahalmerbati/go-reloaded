package functions

import "strings"

// The function is used for processing text by applying a series of transformations to the input string
func ProcessText(content string) string {
	content = ToDec(content)
	content = Vowels(content)
	content = Punctuation(content)
	content = NumberedCaseConverter(content)
	content = CaseConverter(content)
	content = Quotation(content)

	content = strings.TrimSpace(content)

	return content
}
