package main

import (
	"fmt"
	"time"
)

func main() {
	a := 2
	fmt.Print(a, ",")
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}

	time := time.Now()
	switch {
	case time.Hour() < 10:
		fmt.Println("it's before 10")
	case time.Hour() < 12:
		fmt.Println("it's after 10, before 12")
	default:
		fmt.Println("its after 12")
	}
	fmt.Println(time.Hour())
}
