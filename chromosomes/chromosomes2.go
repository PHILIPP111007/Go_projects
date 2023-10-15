package main

// In this version we will fill chromosomes channel
// and increment atomic counter

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup // 1

type chrAlleles []string
type chrsDict map[string]chrAlleles
type Container struct {
	chrsNum atomic.Uint32
	chrs    chan chrsDict
}

const chrsCount int = 10

func fillChannel(i int, c *Container) {
	defer wg.Done() // 3

	time.Sleep(time.Millisecond * 1000)
	chr := fmt.Sprintf("chr%d", i)
	c.chrsNum.Add(1) // increment atomic counter
	c.chrs <- chrsDict{
		chr: chrAlleles{"a", "t", "g", "c"},
	}
}

func main() {
	c := Container{chrs: make(chan chrsDict, chrsCount)}
	wg.Add(chrsCount) // 2

	for i := 1; i <= chrsCount; i++ {
		go fillChannel(i, &c)
	}

	wg.Wait() // 4
	close(c.chrs)

	for chr := range c.chrs {
		fmt.Println(chr)
	}

	fmt.Println("chromosomes number:", c.chrsNum.Load())
}
