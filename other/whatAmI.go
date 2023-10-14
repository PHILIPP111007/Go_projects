package main

import "fmt"

func whatAmI(i interface{}) {

	switch t := i.(type) {
	case bool:
		fmt.Println("I am bool.")
	case int:
		fmt.Println("I am integer.")
	default:
		fmt.Printf("I am %T.\n", t)
	}
}

func main() {
	whatAmI(true)
	whatAmI(10)
	whatAmI(1.1)
}
