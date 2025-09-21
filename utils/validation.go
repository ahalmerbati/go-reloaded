package utils

import (
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func ValidateInputFiles() {
	if len(os.Args) != 3 {
		log.Fatal("Error: Invalid length of arguments")
	}

	sampleFile := os.Args[1]
	resultFile := os.Args[2]

	if !strings.HasSuffix(sampleFile, ".txt") || !strings.HasSuffix(resultFile, ".txt") {
		log.Fatal("Error: The file to be read from and the file to be written to have to both be a txt file")
	}

	if sampleFile == resultFile {
		log.Fatal("The input file cannot be the same as the output file")
	}

	if !CheckFileSize(sampleFile) {
		log.Fatal("Error: The file size exceeds the maximum limit of 100KB")
	}

	content, err := os.ReadFile(sampleFile)
	if err != nil {
		log.Fatal("Error: Could not read the file")
	}

	if !CheckAsciiAndPrintables(string(content)) {
		log.Fatal("Error: Cannot process content as it contains non-ASCII or non-printable characters.")
	}
}

// The function verifies that a given string contains only ASCII and printable characters and replaces all tabs and new lines with a space
func CheckAsciiAndPrintables(content string) bool {
	content = strings.ReplaceAll(content, "\n", " ")
	content = strings.ReplaceAll(content, "\t", " ")

	expression, err := regexp.Compile(`\s+`)
	if err != nil {
		log.Fatal("Error: Could not compile the regular expression")
	}
	content = expression.ReplaceAllString(content, " ")

	for _, r := range content {
		if r > unicode.MaxASCII || !unicode.IsPrint(r) {
			return false
		}
	}
	return true
}

// The function verifies if the file is within the maximum size limit of 100KB
func CheckFileSize(filePath string) bool {
	const maxSize int64 = 102400 // 100KB

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatal("Error: Could not get file information")
	}

	if fileInfo.Size() > maxSize {
		log.Fatal("Error: The file size exceeds the maximum limit of 100KB")
	}

	return true
}

func CheckCommand(command string) bool {
	cmd := strings.Trim(command, "()")

	for _, v := range cmd {
		if !unicode.IsLetter(v) && !unicode.IsNumber(v) && v != ',' && v != ' ' && v != 10 {
			return false
		}
	}
	return true
}

// The function removes all spaces within the bracjets for the command such that ( cap,  4) would become (cap,4)
func CheckCaseConverterSpacing(content string) string {
	var result strings.Builder
	var inParentheses bool

	for _, r := range content {
		if r == '(' {
			inParentheses = true
			result.WriteRune(r)
			continue
		}
		if r == ')' {
			inParentheses = false
			result.WriteRune(r)
			continue
		}
		if inParentheses && r == ' ' {
			continue
		}
		result.WriteRune(r)
	}

	return result.String()
}

// The function checks whether there is a punctuation (.,;:!?) next to the parantheses and if so places a space between them
func CheckASCIIBrackets(content string) string {
	arr := []rune(content)
	var result strings.Builder

	for i := 0; i < len(arr); i++ {
		if i != 0 && arr[i] == '(' && unicode.IsPrint(arr[i-1]) {
			result.WriteRune(' ')
		}

		result.WriteRune(arr[i])

		if i != len(arr)-1 && arr[i] == ')' && unicode.IsPrint(arr[i+1]) {
			result.WriteRune(' ')
		}
	}
	return result.String()
}
