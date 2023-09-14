package main

import "fmt"

func main() {

	// map - it is like a dict in Python

	var people map[string]int = map[string]int{
		"Tom":  14,
		"Phil": 23,
	}

	fmt.Println(people) // map[Phil:23 Tom:14]

	fmt.Println(people["Tom"])  // 14
	fmt.Println(people["Alex"]) // 0

	// ==>
	var name string = "Alex"
	if people[name] != 0 {
		fmt.Printf("%v is here.\n", name)
	} else {
		fmt.Printf("%v is NOT here.\n", name)

		people[name] = 32
		fmt.Println(people[name]) // 32
	}

	delete(people, name)
	fmt.Println(people[name]) // 0

	for key, value := range people {
		fmt.Println(key, value)
	}
}
