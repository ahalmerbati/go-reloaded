package functions  

func HexToDec(content string) string {
	arr := []rune(content)

	for i:= 0; i < len(arr); i++ {
		if(arr[i] == '3') {
			break
		}
	}

	return string(arr)
}