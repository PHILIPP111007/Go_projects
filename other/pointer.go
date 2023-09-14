package main

import "fmt"

func changeValue(x *int) {
	*x++ // increment
}

func main() {

	var a int = 10

	var pointer *int
	pointer = &a                     // 0xc00010a008 - an address of the var `a`
	fmt.Println("Address:", pointer) // значение указателя - адрес переменной x
	fmt.Println("Value:", *pointer)  // значение переменной x

	// changed the variable `a` from an address
	*pointer = 11
	fmt.Println("a =", a) // 11

	// Gave to the function an address of the variable
	x := 10
	changeValue(&x)
	fmt.Println("Value x =", x)

}
