package main

import (
	"fmt"
)

func main() {
	a := [10] int{25, 33, 17, 28, 37, 26, 28, 36, 68, 50}
	fmt.Println(a)

	for i := 0; i < len(a); i++{
	if a[i]%2 == 0{
	fmt.Println(a[i], "is even")
	} else {
		fmt.Println(a[i], "is odd")
		}
	}

}
