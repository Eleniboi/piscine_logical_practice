package main

func IterativePower(nb int, power int) int {

	if power == 0 {
		return 1
	}

	if power < 0 {
		return 0
	}

	var result = 1

	for range power {
		result *= nb
	}
	return result
}

// func main() {

// 	fmt.Println(IterativePower(4, 3))
// }
