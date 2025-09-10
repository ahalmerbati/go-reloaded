package functions

import "strings"

func ProcessText(content string) string {
	content = BinToDec(content)
	content = HexToDec(content)
	content = NumberedCaseConverter(content)
	content = UpperCaseConverter(content)
	content = LowerCaseConverter(content)
	content = CapitalizedCaseConverter(content)
	content = Quotation(content)
	content = Punctuation(content)
	content = Vowels(content)

	content = strings.TrimSpace(content)

	return content
}
