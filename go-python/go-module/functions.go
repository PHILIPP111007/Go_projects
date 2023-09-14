package main

// go build -o functions.so -buildmode=c-shared functions.go

import "C"

//export go_loop
func go_loop() {
	for i := 0; i < 1_000_000; i++ {
	}
}

//export go_sum
func go_sum(x int64, y int64) int64 {
	return x + y
}

func main() {}
