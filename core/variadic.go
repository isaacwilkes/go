package main

import (
	"fmt"
	"strings"
)

func add(numbs ...int) {
	fmt.Println("numbers:", numbs)
	sum := 0
	for _, i := range numbs {
		sum += i
	}
	fmt.Println("sum:", sum)

}

func capitalise(words ...string) {
	for _, word := range words {
		word = string(strings.ToUpper(string(word)))
		fmt.Println(word)
	}
}

func main() {
	add(1, 2, 3, 4, 5)
	add(1)
	add(21, 105, 3556, 1)
	add(100, 200, 300, 400, 500, 600)

	nums := []int{1, 2, 3, 4}
	add(nums...)

	capitalise("test", "testing")
	capitalise("running", "capital?")

	words := []string{"these", "words", "will", "become", "capital"}
	capitalise(words...)
}
