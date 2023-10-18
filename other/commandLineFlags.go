package main

import (
	"flag"
	"fmt"
)

func main() {
	wordFl := flag.String("word", "foo", "a string")
	boolFl := flag.Bool("create", false, "a bool")
	numFl := flag.Int("num", 1, "a number")
	flag.Parse()

	fmt.Println("word:", *wordFl)
	fmt.Println("bool:", *boolFl)
	fmt.Println("number:", *numFl)
}

// Run code with flags (-word, -create, -num) or with -h:
//
// Usage of ./commandLineFlags:
//   -create
//         a bool
//   -num int
//         a number (default 1)
//   -word string
//         a string (default "foo")
