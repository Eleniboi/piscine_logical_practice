package main

import (
	"fmt"
	"strings"
)

func CamelToSnakeCase(s string) string {

	if s == "" {
		return ""
	}
	var result strings.Builder

	for i, ch := range s {

		if ch >= 'a' || ch <= 'z' && ch >= 'A' || ch <= 'Z' {

			if ch >= 'A' || ch <= 'Z' && i > 0 {
				result.WriteString("_")

			}
			result.WriteString(string(ch))

		}
	}
	return result.String()
}

func main() {

	fmt.Println(CamelToSnakeCase("samUel"))
}


//HelloWor