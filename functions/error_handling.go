package functions

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ErrorHandling(content string) bool {
	multiCommandSingle, err := regexp.Compile(`\s*\((up|low|cap)\)(\s*\((up|low|cap)\)){1,}`)
	if err != nil {
		log.Fatal("Error: Could not compile the regular expression.")
	}
	
	if multiCommandSingle.MatchString(content) {
		log.Fatal("Error: Found multiple commands following a single word. Only one command is allowed.")
	}

	invalidCmdSingle, err := regexp.Compile(`\(\s*\(+\s*(up|low|cap)\s*\)+\s*\)|\)\s*\((up|low|cap)\)\(`)
	if err != nil {
		log.Fatal("Error: Could not compile the regular expression.")
	}

	if invalidCmdSingle.MatchString(content) {
		log.Fatal("Error: The command format is invalid. Command must be in the format (up), (low), or (cap) with no extra parentheses.")
	}

	invalidCmdNumbered, err := regexp.Compile(`\(\s*\(+\s*(up|low|cap)\s*,\s*-?\d+\s*\)+\s*\)|\)\s*\((up|low|cap),\s*-?\d+\)\(`)
	if err != nil {
		log.Fatal("Error: Could not compile the regular expression.")
	}

	if invalidCmdNumbered.MatchString(content) {
		log.Fatal("Error: The command format is invalid. Command must be in the format (up, <number>), (low, <number>), or (cap, <number>) with no extra parentheses.")
	}

	invalidCmdNumbered2, err := regexp.Compile(`\((up|low|cap),\s*(-?\d+)\)`)
	if err != nil {
		log.Fatal("Error: Could not compile the regular expression.")
	}

	found := invalidCmdNumbered2.FindAllStringSubmatch(content, -1)
	for _, match := range found {
		number, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal("Error: Invalid number format in command.")
		}
		if number <= 0 {
			log.Fatal("Error: The number in the command must be a positive integer.")
		}
	}

	numberedCmds := regexp.MustCompile(`\((up|low|cap)\s*,\s*(\d+)\)`)
	matches := numberedCmds.FindAllStringSubmatchIndex(content, -1)

	for _, match := range matches {
		commandStart := match[0]

		numStr := content[match[4]:match[5]]
		number, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal("Error: Invalid number format in command.")
		}

		precedingContent := content[:commandStart]
		words := strings.Fields(precedingContent)
		wordCount := len(words)

		if number <= 0 {
			log.Fatal("Error: The number in the command must be a positive integer.")
		}

		if wordCount < number {
			fmt.Printf("Warning: Not enough words before the command '%s' to apply. Only found %d word(s) but command requires %d.\n", content[match[0]:match[1]], wordCount, number)
			return true
		}
	}
	return true
}
