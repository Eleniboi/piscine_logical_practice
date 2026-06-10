package main

import (
	"fmt"
	"strings"
)

func PrintStr(s string) {

	for _, ch := range s {
		fmt.Printf("%c\n", ch)
	}
}
func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

// StrRev is serving as a helper function to  sliceRev
func StrRev(s string) string {

	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// sliceStrRev reverse each string in a slice with the help of StrRev Which has the main reversal logic
func sliceStrRev(s []string) string {

	var result strings.Builder
	for _, char := range s {
		result.WriteString(StrRev(char))
	}
	return result.String()
}

// sliceRev reverse a slice
func sliceRev(s []string) string {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return strings.Join(s, " ")
}

// func main() {

// 	// PrintStr("lsssl")
// 	// a := 3
// 	// b := 6
// 	// Swap(&a, &b)
// 	// fmt.Println(a)
// 	// fmt.Println(b)
// 	fmt.Println(sliceRev([]string{"fine","you are","samuel", "bro",}))
// }
