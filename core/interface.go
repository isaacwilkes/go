package main

import (
	"fmt"
)

type I interface {
	Add() int
}

func main() {
	var a I
	b := MyInt(5)
	a = b

	fmt.Println(a.Add())
}

type MyInt int

func (c MyInt) Add() int {
	c = c + c
	return int(c)
}
