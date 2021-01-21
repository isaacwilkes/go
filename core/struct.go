package main

import (
	"fmt"
)

type s struct {
	A int
	B int
	C int
}

func main() {
	x := s{5, 2, 6}
	x.B = 3
	fmt.Println(x.B * x.A)
	p := &x
	p.A = 4
	fmt.Println(x)
}
