package main

import "fmt"

type kilometer int
type contactInfo struct {
	email        string
	phone_number string
}
type person struct {
	name        string
	age         int8
	contactInfo contactInfo
}

type student struct {
	person     person
	university string
	student_id int
}

func nodes() {
	// Хранения ссылки на структуру того же типа
	// При этом надо учитывать, что структура не может иметь поле,
	// которое представляет тип этой же структуры. Например:
	type node struct {
		value int
		next  *node // поле должно представлять указатель на структуру!!!
	}
}

func main() {

	var km_1 kilometer = 10
	fmt.Println(km_1) // 10

	phil := person{
		name: "Phil",
		age:  23,
		contactInfo: contactInfo{
			email:        "r.puk@example.com",
			phone_number: "+7999111000",
		},
	}

	fmt.Println(phil)                // {Phil 23}
	fmt.Println(phil.name, phil.age) // Phil 23

	phil_student := student{
		person:     phil,
		university: "Secha",
		student_id: 213148723416784,
	}

	fmt.Println(phil_student) // {{Phil 23 {r.puk@example.com +7999111000}} Secha 213148723416784}

}
