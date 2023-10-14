package main

import (
	"fmt"
)

func arrays() {
	// var numbers_10 [10]int = [10]int{1, 2, 3, 4, 5, 6}
	numbers_10 := [10]int{1, 2, 3, 4, 5, 6}
	numbers_undef := [...]int{1, 2, 3} // [...] means define a lenght of an array by a Go compiler

	fmt.Println(numbers_10, numbers_undef)
}

func main() {}
