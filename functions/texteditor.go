package functions

import (
	"fmt"
)

func TextEditor(content string) {
	arr := []rune(content)

	for i:= 0; i < len(arr); i++ {
		if !(arr[i] >= 32 && arr[i] <= 126) {
			fmt.Println("The text editor only accepts printable ASCII characters")
			return
		}
	}
}