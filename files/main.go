package main

import "fmt"

type File struct {
	text string
}

func (f *File) write(message string) {
	f.text = message
}

func (f File) read() {
	fmt.Println(f.text)
}

func main() {
	file := File{}
	file.write("DOJWEwjdiwjfqi12312323")
	file.read() // DOJWEwjdiwjfqi12312323
}
