package main

import (
	"fmt"
	"time"
)

func timer(f func(int) int) {
	time_1 := time.Now().UnixMilli()
	for i := 0; i < 1000000; i++ {
		f(70)
	}
	time_2 := time.Now().UnixMilli()
	fmt.Println(time_2 - time_1)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return factorial(n-1) * n
}

func main() {
	timer(factorial)
}
