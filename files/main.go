package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func check(err error) {
	if err != nil {
		// log.Fatalln(err)
		panic(err)
	}
}

func print_env() {
	lib := os.Environ()
	for _, v := range lib {
		fmt.Print(v, "\n\n")
	}
}

func create_file(file_path string) {
	file, err := os.Create(file_path)
	defer file.Close()

	check(err)

	fmt.Println("Created:", file.Name())                    // Created: hello.txt
	fmt.Println("File extension:", filepath.Ext(file_path)) // .txt

	for i := 1; i <= 10; i++ {
		s := fmt.Sprintf("[%d]: Hello world!\n", i)
		file.WriteString(s)
	}
}

func read_file_1(file_path string) {
	file, err := os.Open(file_path)
	defer file.Close()

	check(err)

	var text string

	// One cycle -> one readed byte and saved as a string to the text variable
	for {
		b := make([]byte, 1)
		_, err := file.Read(b)

		// if err == EOF break
		if err != nil {
			break
		}
		text += string(b)
	}
	fmt.Println(text)
	// [1]: Hello world!
	// [2]: Hello world!
	// [3]: Hello world!
	// [4]: Hello world!
	// [5]: Hello world!
	// [6]: Hello world!
	// [7]: Hello world!
	// [8]: Hello world!
	// [9]: Hello world!
	// [10]: Hello world!
}

func read_file_2(file_path string) {
	data, err := os.ReadFile(file_path)
	check(err)
	fmt.Println(string(data))
}

func delete_file(file_path string) {
	err := os.Remove(file_path)
	check(err)
}

func read_dir(dir_path string) {
	data, err := os.ReadDir(dir_path)
	check(err)
	fmt.Println(data)
}

// ---------------------------------

// ! incomplete (built-in function `filepath.Walk`)
func walk_dir_1(dir_path string) {
	// data, err := filepath.Walk(dir_path, )
}

// My implementation
func walk_dir_2(dir_path string) {
	data, err := os.ReadDir(dir_path)
	check(err)

	for _, p := range data {
		fmt.Println(p)
		if p.IsDir() {
			newDir := filepath.Join(dir_path, p.Name())
			walk_dir_2(newDir)
		}
	}
}

func main() {
	log.SetFlags(0)

	file_path := "hello.txt"
	create_file(file_path)
	read_file_1(file_path)
	read_file_2(file_path)
	time.Sleep(time.Millisecond * 1000)
	delete_file(file_path)

	read_dir("../other") // [- arrays.go - channelDirections.go - channels.go ...]
	read_dir("../")      // [- .DS_Store d .git/ - .gitignore - LICENSE - README.md d chromosomes/ d files/ d go-python/ d greeting/ d other/]

	walk_dir_2("../")
}
