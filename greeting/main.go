package main

import (
	"fmt"
	"log"

	greeting "my-module/module"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) // disable time printing

	names := []string{
		"Phil",
		"iefh",
		"Natali",
	}

	for i, name := range names {

		// msg, err := greeting.Hello(name)
		var msg, err = greeting.Hello(name)

		if err != nil {
			var errMsg = fmt.Sprintf("%v at index %v.", err, i)
			log.Fatalln(errMsg)
		}
		names[i] = msg
	}

	fmt.Println(names)
	fmt.Println(len(names))
}
