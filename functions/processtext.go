package functions

import "strings"

// The function is used for processing text by applying a series of transformations to the input string 
func ProcessText(content string) string {
	content = ToDec(content)
	content = NumberedCaseConverter(content)
	content = CaseConverter(content)
	content = Punctuation(content)
	content = Quotation(content)
	content = Vowels(content)

	content = strings.TrimSpace(content)

	return content
}
