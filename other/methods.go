package main

import (
	"fmt"
	"log"
)

type Array []int
type person struct {
	name string
	age  int
}

/*
	-------- Person methods --------
*/

// function method insertion
func (p person) print() {
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "eats", meal)
}

func (p *person) updateAge(age int) {
	// To update user age, we need to send to this
	// function a pointer with variable address
	p.age = age
}

/*
	-------- Array methods --------
*/

const index_out_msg string = "Index out of range."

// Check if index is out of range.
func _check_index_and_return_len_arr(array Array, index int) int {

	log.SetFlags(0)
	var len_arr int = len(array)

	if index < 0 && index < -len_arr {
		log.Fatalln(index_out_msg)
	} else if index >= len_arr {
		log.Fatalln(index_out_msg)
	}
	return len_arr
}

func (array Array) __getitem__(index int) int {
	len_arr := _check_index_and_return_len_arr(array, index)

	if index < 0 {
		return array[len_arr+index]
	}
	return array[index]
}

func (array *Array) __delitem__(index int) {
	len_arr := _check_index_and_return_len_arr(*array, index)
	if index < 0 {
		*array = append((*array)[0:index+len_arr], (*array)[index+len_arr+1:len_arr]...)
	} else {
		*array = append((*array)[0:index], (*array)[index+1:len_arr]...)
	}
}

func (array Array) __listitems__() {
	for _, val := range array {
		fmt.Println(val)
	}
}

func getitem(array Array, index int) int {
	return array.__getitem__(index)
}
func delitem(array *Array, index int) {
	array.__delitem__(index)
}
func pop(array *Array) int {
	last_item := array.__getitem__(-1)
	array.__delitem__(-1)
	return last_item
}
func listitems(array Array) {
	array.__listitems__()
}

/*
	-------- Testing methods --------
*/

func TryArrays() {
	arr := Array{1, 2, 3, 4}

	fmt.Println(arr.__getitem__(0), getitem(arr, -1))
	listitems(arr)
	// fmt.Println(arr.__getitem__(-5)) // Index out of range.

	fmt.Println("First:", arr) // First: [1 2 3 4]
	delitem(&arr, 0)
	fmt.Println("Second:", arr) // Second: [2 3 4]
	arr.__delitem__(1)
	fmt.Println("Third:", arr) // Third: [2 4]
	arr.__delitem__(-2)
	fmt.Println("Forth:", arr) // Forth: [4]
	last_item := pop(&arr)
	fmt.Println("Fifth:", last_item, arr) // Fifth: 4 []
}

func TryPersons() {
	Phil := person{name: "Phil", age: 11}

	fmt.Println(Phil) // {Phil 11}
	Phil.updateAge(23)
	fmt.Println(Phil) // {Phil 23}
}

func main() {
	TryArrays()
	TryPersons()
}
