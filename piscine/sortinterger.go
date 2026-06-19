package main

import "fmt"

//The sortinterTable is a fuction that sort out an unordered numbers into an ordered manner 
func SortIntegerTable(table []int) {

	//var result = make([]int, len(table))

	for pass := 0; pass < len(table); pass++{
		for i := 0; i < len(table)-1; i++{

			if table[i] > table[i+1]{
				temp := table[i]
				table[i] = table[i+1]
				table[i+1] = temp
			}
		}
	}
}
func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	SortIntegerTable(s)
	fmt.Println(s)
}
