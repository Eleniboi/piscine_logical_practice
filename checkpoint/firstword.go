package main

import (
	"fmt"
)

func FirstWord(s string) string {
	var result string

	foundWord := false

	for _, char := range s {

		if char != ' ' {
			foundWord = true
			result += string(char)
			continue
		}

		if foundWord {
			break
		}
	}

	return result + "\n"
}
func main() {
	fmt.Print(FirstWord("      hello"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello .........  bye"))
}

// func FirstWord(s string) string {

// 	if s == "" {
// 		return "\n"
// 	}

// 	var build strings.Builder

// 	for _, char := range s {

// 		if char == 32 {
// 			build.WriteString("\n")
// 			break
// 		}
// 		build.WriteRune(char)
// 	}
// 	return build.String()
// }
