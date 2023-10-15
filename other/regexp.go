package main

import (
	"fmt"
	"regexp"
)

var p = fmt.Println

func main() {
	r, _ := regexp.Compile("p([a-z]+)ch")
	s := "peach punch peach bike"

	p(r.MatchString(s))               // true
	p(r.FindString(s))                // peach
	p(r.FindAllString(s, -1))         // [peach punch peach]
	p(r.ReplaceAllString(s, "fruit")) // fruit fruit fruit bike
	p(s)                              // peach punch peach bike
}
