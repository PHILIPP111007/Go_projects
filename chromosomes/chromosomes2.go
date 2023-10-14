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

func fillChannel(i int, chrs chan<- chrsDict) {
	defer wg.Done() // 3

	time.Sleep(time.Millisecond * 1000)
	chr := fmt.Sprintf("chr%d", i)
	chrs <- chrsDict{
		chr: chrAlleles{"a", "t", "g", "c"},
	}
}

func main() {
	chrs := make(chan chrsDict, chrsCount)
	wg.Add(chrsCount) // 2

	for i := 1; i <= chrsCount; i++ {
		go fillChannel(i, chrs)
	}

	wg.Wait() // 4
	close(chrs)

	var chrNum int
	for chr := range chrs {
		fmt.Println(chr)
		chrNum++
	}

	fmt.Println("chromosomes number: ", chrNum)
}
