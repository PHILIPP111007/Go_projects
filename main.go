package main

import (
	"fmt"
	"os"
)

// func main() {
// 	const N = 10
// 	c := make(chan int8, N)
// 	for i := 1; i <= N; i++ {
// 		go func() {
// 			fmt.Println(i)
// 			c <- 1
// 		}()
// 	}
// 	for range N {
// 		<-c
// 	}
// 	fmt.Println("End.")
// }

// func main() {
// 	c := make(chan int)

// 	c <- 10

// 	fmt.Println(<-c)
// }

// func main() {
// 	var file_path string = "/Users/phil/Downloads/Roshchin_performance_rewiev.md"

// 	file, err := os.Open(file_path)
// 	if err != nil {
// 		fmt.Println("Unable to open this file, ", err)
// 	}
// 	defer file.Close()

// 	io.Copy(os.Stdout, file)
// }

// func main() {
// 	file, err := os.Create("confeve.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	defer file.Close()
// 	fmt.Fprint(file, "Сегодня ")
// 	fmt.Fprintln(file, "хорошая погода")
// }

// func main() {
// 	var name string

// 	fmt.Print("Введите ваше имя: ")
// 	fmt.Fscan(os.Stdin, &name)
// 	fmt.Println(name)
// }

// func main() {
// 	resp, _ := http.Get("https://google.com/")
// 	defer resp.Body.Close()

// 	io.Copy(os.Stdout, resp.Body)

// }

func writeData(filename string) {
	// начальные данные
	var name string = "Tom"
	var age int = 24

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintln(file, name)
	fmt.Fprintln(file, age)
}
func readData(filename string) {

	var name string
	var age int

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fscanln(file, &name)
	fmt.Fscanln(file, &age)
	fmt.Println(name, age)
}
