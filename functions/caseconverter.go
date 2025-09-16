package functions

import (
	"fmt"
	"regexp"
	"strings"
)

func CaseConverter(content string) string {
	if !ErrorHandling(content) {
		return content
	}

	caseCmd, err := regexp.Compile(`((\b[a-zA-Z0-9']+\b(\s*['.,!?:;]\s*)){1,})\s*\((up|low|cap)\s*,?\s*(\d+)?\)`)
	if err != nil {
		fmt.Println("Error: Could not compile the regular expression.")
		return content
	}

	var result strings.Builder
	index := 0

	found := caseCmd.FindAllStringSubmatchIndex(content, -1)
	

	if len(found) == 0 {
		return content
	}

	for _, match := range found {
		
	}


}
