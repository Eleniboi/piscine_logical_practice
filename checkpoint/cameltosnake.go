package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CamelToSnakeCase(s string) string {

	var result strings.Builder

	if s == "" {
		return ""
	}

	for i, char := range s {

		if unicode.IsUpper(char) && unicode.IsUpper(rune(s[i+1])) {
			return s
		}
		if char >= 'A' && char <= 'Z' && i > 0 {
			result.WriteString("_")
			result.WriteRune(char + 32)
			continue
		}
		result.WriteRune(char)
	}
	return result.String()
}

func main() {

	fmt.Println(CamelToSnakeCase("lowerCamelCase"))
}

// func CamelToSnakeCase(s string) string {
// 	if s == ""{
// 		return ""
// 	}
// 	var result strings.Builder

// 	for i, ch := range s {
// 		if unicode.IsUpper(ch) && unicode.IsUpper(rune(s[i+1])){
// 			return s
// 		}
// 		if ch >= 'A' && ch <= 'Z' && i > 0 {
// 			result.WriteByte('_')
// 			result.WriteRune(ch+32)
// 			continue
// 		}

// 		result.WriteRune(ch)
// 	}

// 	return result.String()
// }

// //HelloWor
