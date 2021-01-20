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

	result := 0
	for x := 0; x < len(a); x++ {
		result += a[x]
	}
	fmt.Println(result)

	if result%2 == 0 {
		fmt.Println("the sum of numebers in array is even:", result)
	} else {
		fmt.Println("the sum of numebers in array is odd:", result)
	}

}
