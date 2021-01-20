package main

import (
	"fmt"
)

func main() {
	a := [15]int{25, 33, 17, 28, 37, 26, 28, 36, 68, 50, 34, 46, 474, 84, 29}
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		if a[i]%2 == 0 {
			fmt.Println(a[i], "is even number")
		} else {
			fmt.Println(a[i], "is odd number")
		}
	}

	sum := 0
	for x := 0; x < len(a); x++ {
		sum += a[x]
	}
	fmt.Println(sum)

	if sum%2 == 0 {
		fmt.Println("the sum of numbers in array is even:", sum)
	} else {
		fmt.Println("the sum of numbers in array is odd:", sum)
	}

}
