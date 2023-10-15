package main

//! incomplete
// There is no used `text/templates` or `http/templates` modules

import "fmt"

func main() {
	s := "Value is %v"
	v1 := []string{"Alex", "Phil"}

	fmt.Println(fmt.Sprintf(s, v1)) // Value is [Alex Phil]
}
