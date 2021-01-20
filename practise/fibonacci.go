package main

import (
	"fmt"
	"time"
)

//func fibonacci() func() int {
//	a, b := -1, 1
//	return func() int {
//		a, b = b, a+b
//		return b
//	}
//}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	start := time.Now()

	fmt.Printf("start time: %d:%d:%d:%d\n",
		start.Hour(),
		start.Minute(),
		start.Second(),
		start.Nanosecond())

	now := time.Now()
	g := fibonacci()
	for i := 0; i < 51; i++ {
		now := time.Now()

		h, m, s, n := now.Hour(), now.Minute(), now.Second(), now.Nanosecond()
		fmt.Println("fib:", g(), "iteration:", i, "time:", h, ":", m, ":", s, ":", n)
	}

	fmt.Printf("time ended: %d:%d:%d:%d\n",
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Nanosecond(),
	)

	elapsed := time.Since(start)
	fmt.Printf("program took: %s", elapsed) //calculated elapsed time

}
