package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	s := "https://github.com/rrrr/PHILIPP111007?tab=repositories&user=Phil"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u)        // equals to s
	fmt.Println(u.Scheme) // https
	fmt.Println(u.Host)   // github.com
	fmt.Println(u.Path)   // /rrrr/PHILIPP111007

	rawQuery := u.RawQuery
	fmt.Println(rawQuery) // tab=repositories&user=Phil

	// first path
	for _, kv := range strings.Split(rawQuery, "&") {
		kv := strings.Split(kv, "=")
		fmt.Println(kv[0], " ", kv[1]) // tab repositories
		// user Phil
	}

	// or do this - quicker and effective :)
	m, _ := url.ParseQuery(rawQuery)
	fmt.Println(m) // map[tab:[repositories] user:[Phil]]
}
