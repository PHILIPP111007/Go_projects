package main

import "fmt"

func create_unbuffered_channel() {
	intChannel := make(chan int) // initilize unbuffered channel with `nil` value

	go func() {
		intChannel <- 5
	}()

	fmt.Println(intChannel)   // 0xc000072060 - a channel address
	fmt.Println(<-intChannel) // 5 - got a channel value
}

// -----------------------------------

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return factorial(n-1) * n
}

func call_factorial_1(n int, intChannel chan int) {
	intChannel <- factorial(n) // Loads result to the channel
}

func call_factorial_2(val int) {
	fmt.Println(factorial(val)) // 720
}

func call_function() {
	intChannel := make(chan int)

	go call_factorial_1(3, intChannel)
	call_factorial_2(<-intChannel) // is waiting for `call_factorial_1` end
}

// -----------------------------------

func create_buffered_channel() {
	intCh := make(chan int, 3)
	intCh <- 10
	intCh <- 3
	intCh <- 24
	fmt.Println(<-intCh) // 10
	fmt.Println(<-intCh) // 3
	fmt.Println(<-intCh) //24
}

// -----------------------------------

func main() {
	// create_unbuffered_channel()
	// create_buffered_channel()
	call_function()
}
