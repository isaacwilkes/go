package main

import (
	"errors"
	"fmt"
)

//SortNum ...
func SortNum(numbs []int) ([]int, error) {

	if len(numbs) > 10 {
		return nil, errors.New("slice contains too many numbers")
	}

	for i := 0; i < len(numbs); i++ {
		for n := 1; n < len(numbs)-i; n++ {

			if numbs[n] < 0 {
				return nil, errors.New("slice contains negative number")
			}

			if numbs[n] < numbs[n-1] {
				numbs[n], numbs[n-1] = numbs[n-1], numbs[n]
			}
		}
	}
	return numbs, nil
}

func main() {

	a, err := SortNum([]int{23, 12, 23123, 45, 1, 5, 13, 16, 23, 25, 22, 5234})

	if err != nil {
		fmt.Println("error occured:", err)
	} else {
		fmt.Println(a)
	}
}
