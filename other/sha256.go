package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "Hello world!"
	hashFunc := sha256.New()
	hashFunc.Write([]byte(s))
	byteSum := hashFunc.Sum(nil)

	fmt.Println(s)
	fmt.Printf("Hash for `%s`: %x\n", s, byteSum) // c0535e4be2b79ffd93291305436bf889314e4a3faec05ecffcbb7df31ad9e51a
}
