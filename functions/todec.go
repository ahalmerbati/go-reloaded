package functions

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// The function converts hexadecimal and binary numbers within a string to their decimal equivalents
func ToDec(content string) string {
	validCmd := regexp.MustCompile(`(\b(?:[0-9a-fA-F]+|[01]+)\b)\s*\((hex|bin)\)`)

	var result strings.Builder
	lastIndex := 0

	for _, match := range validCmd.FindAllStringSubmatch(content, -1) {
		matchStart := strings.Index(content, match[0])
		matchEnd := matchStart + len(match[0])

		result.WriteString(content[lastIndex:matchStart])

		oldNum := strings.TrimSpace(match[1])
		command := strings.TrimSpace(match[2])

		var base int
		switch command {
		case "hex":
			base = 16
		case "bin":
			base = 2
		}

		decNum, err := strconv.ParseInt(oldNum, base, 64)
		if err != nil {
			fmt.Printf("Error: Could not convert %s to a decimal number.\n", oldNum)
			result.WriteString(oldNum + " (" + command + ")")
		} else {
			result.WriteString(fmt.Sprintf("%d", decNum))
		}

		lastIndex = matchEnd
	}

	result.WriteString(content[lastIndex:])

	checkCmd := regexp.MustCompile(`\s*\((hex|bin)\)`)
	if checkCmd.MatchString(result.String()) {
		log.Fatal("Error: The command format is invalid. Command must be in the format <number> (hex) or <number> (bin).")
	}

	return result.String()
}
