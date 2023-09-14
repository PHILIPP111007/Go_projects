package main

import "fmt"

// func takes undefined numder of integer arguments
func add_from_int_args(numbers ...int) {
	var sum int = 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Println("\nSum of numbers =", sum)

}

func add_from_int_slice(numbers ...int) {
	var sum int = 0
	for _, n := range numbers {
		sum += n
	}
	fmt.Println("\nSum of numbers =", sum)
}

func add(x int, y int) int {
	return x + y
}

// this function takes two int args and one function as argument
func takes_one_func(x, y int, operator func(int, int) int) {
	fmt.Println(operator(x, y))
}

func return_lambda_func() func() int {
	f := func() int { return 1 } // var f func(int, int) int = func(x, y int) int{ return 1}
	return f
}

func finish() {
	fmt.Println("Program has been finished")
}

func division(a int, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func main() {
	// Оператор defer позволяет выполнить определенную функцию в конце программы,
	// при этом не важно, где в реальности вызывается эта функция
	defer finish()

	division(1, 0)

	add_from_int_args(1, 2, 3)
	numbers := []int{1, 2, 3} // it is a slice
	fmt.Println(len(numbers), numbers)
	add_from_int_slice(numbers...)

	var add_func func(int, int) int = add
	fmt.Println(add_func(1, 2))

	takes_one_func(10, 10, add)
}
