package main

import "fmt"

func add(c chan int) {
	sum := 5
	c <- sum
}

func main() {
	send := make(chan string)

	go func() { send <- "yo" }()

	a := <-send
	fmt.Println(a)

	c := make(chan int)
	go add(c)
	x := <-c
	fmt.Println(x + 10)

	ch := make(chan int, 5) // Buffered channel
	ch <- 1                 // sends to a channel block when buffer full
	ch <- 10                // receives block when buffer empty
	ch <- 100
	ch <- 1000
	ch <- 10000
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}
