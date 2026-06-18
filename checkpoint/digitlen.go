package main

func DigitLen(n, base int) int {

	if base < 2 || base > 32 {
		return -1
	}

	if n < 0 {
		n = -(n)
	}

	count := 0

	for n != 0 {
		n = n / base
		count++
	}

	return count
}

// func main() {

// 	fmt.Println(DigitLen(100, 38))
// }
