package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 1

type chrAlleles []string
type chrsDict map[string]chrAlleles

const chrsCount int = 10

func fillDict(i int, chrs chrsDict) {
	defer wg.Done() // 3

	time.Sleep(time.Millisecond * 1000)
	chr := fmt.Sprintf("chr%d", i)
	fmt.Println(chr)
	chrs[chr] = chrAlleles{"a", "t", "g", "c"}
}

func main() {
	chrs := chrsDict{}
	wg.Add(chrsCount) // 2

	for i := 1; i <= chrsCount; i++ {
		// wg.Add(1) // 2
		go fillDict(i, chrs)
	}

	wg.Wait()         // 4
	fmt.Println(chrs) // map[chr1:[a t g c] chr10:[a t g c] ... ]
}
