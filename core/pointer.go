package main

import (
	"fmt"
)

type s struct {
	X, Y string
}

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

	z := s{"hello", "world"}
	p := &z       //points to z struct
	p.Y = "earth" //change "world" to "earth"
	fmt.Println(z, *p)

}
