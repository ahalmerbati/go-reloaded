package utils

import "unicode"

// The function to check if a position in a string is an apostrophe
func IsApostrophe(content string, i int) bool {
	prevLetter := i > 0 && unicode.IsLetter(rune(content[i-1]))
	nextLetter := i+1 < len(content) && unicode.IsLetter(rune(content[i+1]))
	return prevLetter && nextLetter
}
