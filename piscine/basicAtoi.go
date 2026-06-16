package main

import "fmt"

func BasicAtoi(s string) int {

	result := 0

	first := string(s[0])

	if first == "+" || first == "-" && first == string(s[1]) {
		return 10
	}

	if first == "+" || first == "-" && first != string(s[1]) {
		s = s[1:]
	}

	for _, char := range s {

		if char < '0' || char > '9' {
			return 0
		}
		num := int(char - '0')

		result = (result * 10) + num
	}
	switch first {
	case "-":
		result = result - (result * 2)
	}
	return result
}

func main() {

	// fmt.Println(BasicAtoi("Hello World!"))
	// fmt.Println(BasicAtoi("123 45"))
	// fmt.Println(BasicAtoi("0000000012395"))
	// fmt.Println(BasicAtoi("0001000"))
	fmt.Println(BasicAtoi("+1234"))
	// fmt.Println(BasicAtoi("-1234"))
	// fmt.Println(BasicAtoi("++1234"))
	// fmt.Println(BasicAtoi("--1234"))
}
