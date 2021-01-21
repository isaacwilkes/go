package main

import (
	"fmt"
)

type I interface {
	Add() int
}

func main() {
	var a I
	a = MyInt(5)


	fmt.Println(a.Add())
}

type MyInt int

func (c MyInt) Add() int {
	c = c + c
	return int(c)
}
