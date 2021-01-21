package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2

	c := &a            // c holds the memory address of a
	fmt.Println(*c, b) // *c reads a through pointer
	fmt.Println(c)     // memory address of a

	*c = 3 //change value of a through pointer
	fmt.Println(a)

	d := &b
	*d = *d * 4     // multiply b through pointer d
	fmt.Println(*d) // print
	fmt.Println(b)  //value of b is now 8

}
