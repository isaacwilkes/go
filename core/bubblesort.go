package main

import (
	"fmt"
)

//SortNum ...
func SortNum(numbs []int) []int {
	for i := 0; i < len(numbs); i++ {
		for n := 1; n < len(numbs)-i; n++ {
			if numbs[n] < numbs[n-1] {
				numbs[n], numbs[n-1] = numbs[n-1], numbs[n]
				//fmt.Println(numbs)
			}
		}
	}
	return numbs
}

func main() {
	a := []int{23, 12, 23123, 23, 1235, 3, 6, 88888, 34, 345, 3667, 863, 953, 400, 3431, 11}

	fmt.Println(SortNum(a))
}
