package main

func IterativeFactorial(nb int) int {

	//this controls negative number and overflow too
	//factorial of number greater than 12 will not fit into int 34
	if nb < 0 || nb > 12 {
		return 0
	}

	result := 1

	for i := nb; i >= 1; i-- {

		result *= i
	}
	return result
}

// func main() {

// 	fmt.Println(IterativeFactorial(5))
// }

// if nb <= 0{
// 	return 0
// }

// var result = 1

// for i := 1; i <= nb; i++ {

// 		result *= i
// 		fmt.Println(result)
// }
// return result
