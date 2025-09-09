package functions

import (
	"fmt"
	"regexp"
	"strings"
)

func Punctuation(content, pattern1, pattern2 string) string {
	expression, err := regexp.Compile(pattern1)
	if err != nil {
		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
		return content
	}

	var result strings.Builder
	index := 0

	found := expression.FindAllStringSubmatchIndex(content, -1)

	if len(found) == 0 {
		fmt.Println("No instances of puncutation marks were found in the content. The original content is returned:")
		return content
	}

	for _, match := range found {
		result.WriteString(content[index:match[0]])
		result.WriteString(content[match[2]:match[3]])
		result.WriteString(" ")
		index = match[1]
	}
	result.WriteString(content[index:])

	finalString := strings.TrimSpace(result.String())

	expression2, err := regexp.Compile(pattern2)
	if err != nil {
		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
		return content
	}

	finalString = expression2.ReplaceAllString(finalString, "$1 ")

	return finalString
}

func SinglePunctuation(content string) string {
	return Punctuation(content, `\s*([.,!?;:])\s*`, `([.,!?;:])\s+`)
}

func GroupPunctuation(content string) string {
	return Punctuation(content, `\s*([.,!?:;]+)\s*`, `([.,!?;:]+)\s+`)
}

// func SinglePunctuation(content string) string {
// 	expression, err := regexp.Compile(`\s*([.,!?;:])\s*`)
// 	if err != nil {
// 		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
// 		return content
// 	}

// 	var result strings.Builder
// 	index := 0

// 	found := expression.FindAllStringSubmatchIndex(content, -1)

// 	if len(found) == 0 {
// 		fmt.Println("No instances of punctuation marks were found in the content. The original content is returned:")
// 		return content
// 	}

// 	for _, match := range found {
// 		result.WriteString(content[index:match[0]])
// 		result.WriteString(content[match[2]:match[3]])
// 		result.WriteString(" ")
// 		index = match[1]
// 	}

// 	result.WriteString(content[index:])

// 	expression2, err := regexp.Compile(`([.,!?;:])\s+`)
// 	if err != nil {
// 		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
// 		return content
// 	}

// 	finalString := expression2.ReplaceAllString(result.String(), "$1 ")

// 	return finalString
// }

// func GroupPunctuation(content string) string {
// 	expression, err := regexp.Compile(`\s*([.,!?:;]+)\s*`)
// 	if err != nil {
// 		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
// 		return content
// 	}

// 	var result strings.Builder
// 	index := 0
// 	found := expression.FindAllStringSubmatchIndex(content, -1)

// 	for _, match := range found {
// 		result.WriteString(content[index:match[0]])
// 		result.WriteString(content[match[2]:match[3]])
// 		result.WriteString(" ")
// 		index = match[1]
// 	}

// 	result.WriteString(content[index:])

// 	finalString := strings.TrimSpace(result.String())

// 	expression2, err := regexp.Compile(`([.,!?;:]+)\s+`)
// 	if err != nil {
// 		fmt.Println("There was an error compiling the regular expression. The original content is returned:")
// 		return content
// 	}

// 	finalString = expression2.ReplaceAllString(finalString, "$1 ")

// 	return finalString
// }
