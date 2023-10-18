package main

import (
	"fmt"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

var p = fmt.Println

func get() {
	resp, err := http.Get("https://gobyexample.com/http-client")
	check(err)
	defer resp.Body.Close()
	p(resp, "\n")
	p(resp.Header, "\n")
	p(resp.Body, "\n")
	p(resp.Status, "\n")
	p(resp.StatusCode, "\n")
	p(resp.Location())
}

func home(w http.ResponseWriter, req *http.Request) {
	s := `
	<h1>Hello world</h1>
	<p>wpfkewpfk</p>fewfewf owefjoefjeofjew
	`
	fmt.Fprint(w, s)
}

func main() {
	port := "127.0.0.1:8000"

	http.HandleFunc("/", home)
	http.ListenAndServe(port, nil)
}
