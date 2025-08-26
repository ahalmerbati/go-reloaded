package functions

import "strings"

func HexToDec(content string) string {
	contentCheck := strings.Contains(content, "(hex)")

	if !(contentCheck) {
		return content
	}
	
	arr := []rune(content)

	if contentCheck {
		for i := 0; i < len(arr); i++ {
			if string(arr[i:i+5]) == "(hex)" {
				// for j:= 0; j < i; j++ {

				// }
			}
		}
	}
	return string(arr)
}
	

