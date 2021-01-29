package main

import "fmt"

type numbs struct {
	Num1, Num2 int
}

var m map[int]numbs

func main() {
	m = make(map[int]numbs)
	m[1] = numbs{
		1, 2,
	}
	fmt.Println(m[1])
	fmt.Println(m)

	m2 := make(map[string]string)

	m2["First"] = "One"
	fmt.Println(m2["First"])

	m2["Second"] = "Two"
	fmt.Println(m2["Second"])

	fmt.Println(m2)

	delete(m2, "Second")
	fmt.Println(m2)

	a, ok := m2["Second"]
	fmt.Println("second present:", ok, a)
}
