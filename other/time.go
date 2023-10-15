package main

import (
	"fmt"
	"time"
)

var print = fmt.Println

func main() {
	now := time.Now()
	print(now.UTC())   // 2023-10-15 15:22:29.657658 +0000 UTC
	print(now.Clock()) // 18 22 29
	print(now.Date())  // 2023 October 15

	now = now.Add(time.Hour * 24 * 365)
	print(now.UTC())     // 2024-10-14 15:22:29.657658 +0000 UTC
	print(now.Clock())   // 18 22 29
	print(now.Date())    // 2024 October 14
	print(now.Weekday()) // Monday

	diff := now.Sub(time.Date(2000, 1, 10, 10, 0, 0, 0, time.UTC))
	print(diff)               // 217061h27m20.645947s
	print(diff.Abs())         // 217061h27m56.794185s
	print(diff.Hours())       // 217061.4657761625
	print(diff.Nanoseconds()) // 781421276794185000
}
