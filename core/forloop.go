package main

import (
	"fmt"
)

func main() {
	for i := 0; i<20; i++{
	fmt.Println(i)
	i += i
	}
	
	x := 0
	for x <= 3 {
	fmt.Println(x)
	x = x + 1
	}
	
}

