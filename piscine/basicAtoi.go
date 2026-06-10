package main

import "fmt"

func BasicAtoi(s string) int {
	result := 0

	for _, char := range s {

		// convert char to digit
		num := int(char - '0')

		// update result
		result = (result * 10) + num
	}

	return result
}

func main() {

	fmt.Println(BasicAtoi("000000122234"))
}
