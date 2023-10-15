package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var print = fmt.Println

type University struct {
	Name          string `json:"name"`
	StudentsCount int    `json:"students_count"`
}
type Person struct {
	Name       string     `json:"name"`
	Age        int        `json:"age"`
	University University `json:"university"`
}

type Students map[string]Person

func workWithStudents() {
	s := `
	{
		"name": "Phil", 
		"age": 23,
		"university": {
			"name": "Secha",
			"students_count": 1000
		}
	}
	`
	person := Person{}
	json.Unmarshal([]byte(s), &person)
	person.Name = "Alex"
	newPerson, _ := json.Marshal(&person)

	print(person)                                                                           // {Phil 23 {Secha 1000}}
	print(person.Age, person.Name, person.University.Name, person.University.StudentsCount) // 23 Phil Secha 1000
	print(string(newPerson))                                                                // {"name":"Alex","age":23,"university":{"name":"Secha","students_count":1000}}

	// list of students
	s = `
	[
		{"name":"Alex","age":10,"university":{"name":"Secha","students_count":1000}},
		{"name":"Phil","age":11,"university":{"name":"Oxford","students_count":10000}}
	]
	`
	students_1 := []Person{}
	json.Unmarshal([]byte(s), &students_1)
	print(students_1) // [{Alex 10 {Secha 1000}} {Phil 11 {Oxford 10000}}]

	// dict of students with id
	s = `
	{
		"1": {"name":"Alex","age":10,"university":{"name":"Secha","students_count":1000}},
		"2": {"name":"Phil","age":11,"university":{"name":"Oxford","students_count":10000}}
	}
	`
	students_2 := Students{}
	json.Unmarshal([]byte(s), &students_2)
	print(students_2) // map[1:{Alex 10 {Secha 1000}} 2:{Phil 11 {Oxford 10000}}]
}

func workWithTypes() {
	s := `{"integer": 100, "float": 6.14, "list": ["123", "dqw123"], "boolean": true}`
	var data map[string]interface{}
	json.Unmarshal([]byte(s), &data)

	fmt.Printf("%T\n", data["integer"]) // float64 !!!!!!!
	fmt.Printf("%T\n", data["float"])   // float64
	fmt.Printf("%T\n", data["boolean"]) // bool
	fmt.Printf("%T\n", data["list"])    // []interface {}

	print(data["list"]) // [123 dqw123]
	// print(data["list"][1]) // cannot index data["list"] (map index expression of type interface{})
	// arr := []string(data["list"]) // cannot convert data["list"] (map index expression of type interface{}) to type []string: need type assertion

	strs := data["list"].([]interface{})
	str1 := strs[0].(string)
	fmt.Printf("%T\n", str1) // string

	i := int(data["integer"].(float64))
	fmt.Printf("%T\n", i) // int
}

func main() {
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d) // {"apple":5,"lettuce":7}
}
