package main

import (
	"fmt"
	"log"
	"os"
)

func print_env() {
	lib := os.Environ()
	for _, v := range lib {
		fmt.Print(v, "\n\n\n")
	}
}

func create_file() {
	file, err := os.Create("hello.txt")
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Created:", file.Name()) // hello.txt

	for i := 0; i < 10; i++ {
		file.WriteString("Hello world!")
	}
}

func read_file() {
	file, err := os.Open("hello.txt")
	defer file.Close()

	if err != nil {
		log.Fatalln(err)
	}

	var text string

	// One cycle -> one readed byte and saved as a string to text variable
	for {
		b := make([]byte, 1)
		_, err := file.Read(b)

		// if err == EOF break
		if err != nil {
			break
		}
		text += string(b)
	}

	fmt.Println(text) // Hello world!
}

func main() {
	log.SetFlags(0)

	create_file()
	read_file()
}
