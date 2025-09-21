package functions

import (
	"go-reloaded/utils"
	"log"
	"strconv"
	"strings"
)

func CaseConverter(content string) string {
	words := strings.Fields(content)
	result := make([]string, 0)
	index := 0

	for index < len(words) {
		word := words[index]

		if strings.HasPrefix(word, "(") && strings.HasSuffix(word, ")") && utils.CheckCommand(word) {
			command := strings.Trim(word, "()")
			if strings.Contains(command, ",") {
				seg := strings.SplitN(command, ",", 2)
				if len(seg) == 2 {
					commandToApply := strings.TrimSpace(seg[0])
					numStr := strings.TrimSpace(seg[1])
					if num64, err := strconv.ParseInt(numStr, 10, 32); err == nil && num64 > 0 {
						num := int(num64)
						utils.ApplyNumCase(result, commandToApply, num)
					} else if num64 <= 0 {
						log.Fatal("Error: For the command to apply the number has to be a positive integer.")
					} else if num64 >= 1000 {
						log.Fatal("Error: The number is too large. Pick any number greater than 0 and less than 1000")
					}
				}
			} else {
				utils.ApplySingleCase(result, command)
			}
			index++
		} else {
			result = append(result, word)
			index++
		}
	}
	return strings.Join(result, " ")
}
