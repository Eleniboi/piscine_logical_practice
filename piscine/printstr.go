package main

import (
	"fmt"
)

func PrintStr(s string) {

	for _, ch := range s {
		fmt.Printf("%c\n", ch)
	}
}
func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

func StrRev(s string) string {

	
}
func main() {

	// PrintStr("lsssl")
	// a := 3
	// b := 6
	// Swap(&a, &b)
	// fmt.Println(a)
	// fmt.Println(b)
	fmt.Println(StrRev("samuel"))
}
