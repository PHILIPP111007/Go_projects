package main

import (
	"fmt"
	"time"
)

func factorial(n int64) int64 {
	if n == 1 {
		return 1
	}
	return factorial(n-1) * n
}

func big_func(f func(int64) int64) {
	for i := 0; i < 1000000; i++ {
		f(70)
	}
}

func main() {
	t_1 := time.Now().UnixMilli()
	for i := 0; i < 7; i++ {
		big_func(factorial)
	}
	t_2 := time.Now().UnixMilli()
	fmt.Println(t_2 - t_1) // 3149 milliseconds

	t_1 = time.Now().UnixMilli()
	for i := 0; i < 7; i++ {
		go big_func(factorial) // `go` means create a goroutine
	}
	t_2 = time.Now().UnixMilli()
	fmt.Println(t_2 - t_1) // 0 millisecond
}
