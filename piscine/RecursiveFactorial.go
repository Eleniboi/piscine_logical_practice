package main 


import (
	"fmt"
)
//is a function that calls it self with the smaller version of it problem
func RecursiveFactorial(nb int) int{

	if nb == 1{
		return 1
	}

	if nb < 0 || nb > 12{
		return 0
	}

	n := nb * RecursiveFactorial(nb - 1)
	return n
}

func main() {
	
	fmt.Println(RecursiveFactorial(5))

}