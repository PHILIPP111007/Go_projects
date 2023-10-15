// [Timers](timers) are for when you want to do
// something once in the future - _tickers_ are for when
// you want to do something repeatedly at regular
// intervals. Here's an example of a ticker that ticks
// periodically until we stop it.

package main

import (
	"fmt"
	"time"
)

func tickChannel() {
	reqCount := 5
	requests := make(chan int, reqCount)

	for i := 0; i < reqCount; i++ {
		requests <- i
	}

	close(requests)
	limiter := time.Tick(time.Millisecond * 100)

	for i := range requests {
		<-limiter
		fmt.Println("request", i, time.Now().UTC())
	}

}

func main() {
	// Tickers use a similar mechanism to timers: a
	// channel that is sent values. Here we'll use the
	// `select` builtin on the channel to await the
	// values as they arrive every 500ms.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// Tickers can be stopped like timers. Once a ticker
	// is stopped it won't receive any more values on its
	// channel. We'll stop ours after 1600ms.
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

	tickChannel()
	fmt.Println("Done")
}

// $ go run tickers.go
// Tick at 2012-09-23 11:29:56.487625 -0700 PDT
// Tick at 2012-09-23 11:29:56.988063 -0700 PDT
// Tick at 2012-09-23 11:29:57.488076 -0700 PDT
// Ticker stopped
