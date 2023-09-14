package main

import (
	"fmt"
)

func main() {

	var s1 []string = []string{"1", "2", "3"}
	fmt.Println(s1) // [1 2 3]

	var s2 []string = make([]string, 3)
	fmt.Println(len(s2), s2) // 3 [  ]

	var s10 []string = make([]string, 10)
	fmt.Println(len(s10), s10) // 10 [         ]

	var s3 []string = make([]string, 3)
	s3 = append(s3, "Tom")
	s3[0] = "Phil"
	fmt.Println(len(s3), s3) // 4 [Phil   Tom]

	// исходный массив
	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users1 := initialUsers[2:6] // с 3-го по 6-й
	users2 := initialUsers[:4]  // с 1-го по 4-й
	users3 := initialUsers[3:]  // с 4-го до конца

	fmt.Println(users1) // ["Kate", "Sam", "Tom", "Paul"]
	fmt.Println(users2) // ["Bob", "Alice", "Kate", "Sam"]
	fmt.Println(users3) // ["Sam", "Tom", "Paul", "Mike", "Robert"]

	// удаляем 4-й элемент
	users := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	var n = 3
	users = append(users[:n], users[n+1:]...)
	fmt.Println(users) //["Bob", "Alice", "Kate", "Tom", "Paul", "Mike", "Robert"]

}
