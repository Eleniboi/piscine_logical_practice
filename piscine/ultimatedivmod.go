package main 

import (
	"fmt"
)

func DivMod(a int, b int, div *int, mod *int) {

	*div = a/b
	*mod = a%b
}
func UltimatePointOne(n ***int) {
	***n = 1
}

func main() {
	//ultimatepointone this is the test process for it
	a := 13
	b := 2
	var div int
	var mod int

	DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)

	//pointont this is the test process for the function 
	n := 0
	p1 := &n
	p2 := &p1

	UltimatePointOne(&p2)
	fmt.Println(**p2)
}