package functions

import "strings"

func ProcessText(content string) string {
	content = BinToDec(content)
	content = HexToDec(content)
	content = NumberedCaseConverter(content)
	content = CaseConverter(content)
	content = Punctuation(content)
	content = Quotation(content)
	content = Vowels(content)

	content = strings.TrimSpace(content)

	return content
}
