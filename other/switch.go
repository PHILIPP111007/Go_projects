package main

import "fmt"

func my_switch() {
	a := 8
	switch a {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	}

	a = 87
	switch a {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	default:
		fmt.Println("значение переменной a не определено")
	}

	a = 5
	switch a {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	case 6, 5, 4:
		fmt.Println("a = 6 или 5 или 4, но это не точно")
	default:
		fmt.Println("значение переменной a не определено")
	}

}

func main() {}
