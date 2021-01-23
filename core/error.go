package main

import (
	"errors"
	"fmt"
)

func main() {
	middle, total, err := multiply(10, 5)

	if err != nil {
		fmt.Println("Error.", err)
	} else {
		fmt.Println(total)
		fmt.Println("median of original numbers", middle)
	}

}

func multiply(start, end int) (float64, int, error) {

	if start > end {
		return 0, 0, errors.New("Invalid")
	}

	total := 1
	for i := start; i <= end; i++ { // multiplies all numbers between set range (3 to 5)
		total *= i
	}

	middle := (float64(start) + float64(end)) / 2
	return middle, total, nil
}
