package main

import (
	"fmt"
	"os"
)

// reads user input from a console
func inf_reading() {

	var code string

	for {
		b := make([]byte, 1)
		_, err := os.Stdin.Read(b)

		// if err != nil || bytes.Equal(b, []byte("\n")) && string(text[len(text)-1]) != "\\") {
		if err != nil || (b[0] == []byte("\n")[0] && string(code[len(code)-1]) != "\\") {
			break
		}
		code += string(b)
	}

	fmt.Println(code)
}

func read_and_make_vars() {
	var name string
	var age int
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	fmt.Println(name, age)
}

func main() {
	// inf_reading()
	read_and_make_vars()
}
