package main

import (
	"fmt"
	"time"
)

func main() {
	countdown(2, 10)
}

func countdown(start, end int) {
	for i := start; i >= end; i-- {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
		time.Sleep(1 * time.Second)
	}

	if end > start {
		fmt.Println("starting number greater than ending number")
	}
}
