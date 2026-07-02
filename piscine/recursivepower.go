package main


import (
	"fmt"
)

//recursivePower it called it self with the smaller version of it problem
// what happen internally 
// there is a base number which will be unchange only the power will decrease at each call 
// nb := 5, power := 2

// 5 * RP(5 x 2)
// 5 * RP(5 x 1)
// 5 x 5 = 25
func RecursivePower(nb, power int)  int {


	if power < 0{
		return 0
	}

	if power == 0{
		return 1
	}

	return nb * RecursivePower(nb, power-1)
}


func main() {

	fmt.Println(RecursivePower(5, 5))
}