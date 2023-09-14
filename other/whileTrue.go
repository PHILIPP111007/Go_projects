package main

import "fmt"

func main() {
	var i int = 0

	for {
		fmt.Println(i)
		if i > 100 {
			break
		}
		i++
	}
}
