package main

import (
	"fmt"
	"time"
)

func send(s string) {
	for i := 0; i < 10; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go send("1")
	send("2")
}
