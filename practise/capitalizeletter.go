package main

import (
	"fmt"
	"strings"
)

func main() {
	for i, letter := range "banana" {
		if i%2 == 0 {
			letterUpper := strings.ToUpper(string(letter))
			fmt.Println(letterUpper)
		} else {
			letterLower := strings.ToLower(string(letter))
			fmt.Println(letterLower)
		}
	}
}
