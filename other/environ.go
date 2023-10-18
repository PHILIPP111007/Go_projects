package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("USER_NAME", "Phil")
	user_name := os.Getenv("USER_NAME")
	foo := os.Getenv("FOO")           // we dont have the foo environ variable
	fmt.Println(user_name)            // Phil
	fmt.Printf("%v - %T\n", foo, foo) // "" - string
}
