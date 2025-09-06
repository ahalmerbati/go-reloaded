package functions

import (
	"fmt"
	"regexp"
	"strings"
)

func CaseConverter(content, pattern string, converter func(string) string) string {
	expression, err := regexp.Compile(pattern)
	if err != nil {
		return "The regex pattern is invalid"
	}
	
	var result strings.Builder
	index := 0
	found := expression.FindAllStringSubmatchIndex(content, -1)
	
	for _, match := range found {
		result.WriteString(content[index:match[0]])
		
		word := content[match[2]:match[3]]
		
		convertedWord := converter(word)
		
		result.WriteString(convertedWord)
		index = match[1]
	}
	result.WriteString(content[index:])
	
	if strings.Contains(result.String(), strings.Split(pattern, `\s*`)[1]) {
		fmt.Println("Some words were not fully converted due to invalid input or other ASCII characters.")
	}
	
	return result.String()
}

func UpperCaseConverter(content string) string {
	return CaseConverter(content, `(\b[a-zA-Z]+\b)\s*\(up\)`, strings.ToUpper)
}

func LowerCaseConverter(content string) string {
	return CaseConverter(content, `(\b[a-zA-Z]+\b)\s*\(low\)`, strings.ToLower)
}

func CapitalizedCaseConverter(content string) string {
	return CaseConverter(content, `(\b[a-zA-Z]+\b)\s*\(cap\)`, capitalize)
}

func capitalize(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(str[:1])
}

// func UpperCaseConverter(content string) string {
	// 	expression, err := regexp.Compile(`(\b[a-zA-Z]+\b)\s*\(up\)`)
	// 	if err != nil {
		// 		return "The regex pattern is invalid"
// 	}

// 	var result strings.Builder

// 	found := expression.FindAllStringSubmatchIndex(content, -1)
// 	index := 0

// 	for _, match := range found {
// 		result.WriteString(content[index:match[0]])

// 		capitalizedWord := strings.ToUpper(content[match[2]:match[3]])

// 		result.WriteString(fmt.Sprint(capitalizedWord))

// 		index = match[1]
// 	}
// 	result.WriteString(content[index:])

// 	if strings.Contains(result.String(), "(up)") {
// 		fmt.Println("Some words were not fully converted due to invalid input or other ASCII characters.")
// 	}

// 	return result.String()
// }

// func LowerCaseConverter(content string) string {
// 	expression, err := regexp.Compile(`(\b[a-zA-Z]+\b)\s*\(low\)`)
// 	if err != nil {
// 		return "The regex pattern is invalid"
// 	}

// 	var result strings.Builder
// 	index := 0

// 	found := expression.FindAllStringSubmatchIndex(content, -1)

// 	for _, match := range found {
// 		result.WriteString(content[index:match[0]])

// 		lowerCaseWord := strings.ToLower(content[match[2]:match[3]])

// 		result.WriteString(fmt.Sprint(lowerCaseWord))

// 		index = match[1]
// 	}
// 	result.WriteString(content[index:])

// 	if strings.Contains(result.String(), "(low)") {
// 		fmt.Println("Some words were not fully converted due to invalid input or other ASCII characters.")
// 	}

// 	return result.String()
// }

// func CapitalizedCaseConverter(content string) string {
// 	expression, err := regexp.Compile(`(\b[a-zA-Z]+\b)\s*\(cap\)`)
// 	if err != nil {
// 		return "The regex pattern is invalid"
// 	}

// 	var result strings.Builder
// 	index := 0

// 	found := expression.FindAllStringSubmatchIndex(content, -1)

// 	for _, match := range found {
// 		result.WriteString(content[index:match[0]])

// 		index2 := match[2]
// 		capitalizedCase := strings.ToUpper(string(content[index2]))

// 		result.WriteString(fmt.Sprint(capitalizedCase))

// 		result.WriteString(content[index2+1 : match[3]])

// 		index = match[1]
// 	}
// 	result.WriteString(content[index:])

// 	if strings.Contains(result.String(), "(cap)") {
// 		fmt.Println("Some words were not fully converted due to invalid input or other ASCII characters.")
// 	}

// 	return result.String()
// }