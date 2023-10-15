package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(10)    // 0 <= n < 10
	f32 := rand.Float32() // 0.0 <= f32 < 1.0
	f64 := rand.Float64() // 0.0 <= f64 < 1.0
	fmt.Println(n)
	fmt.Println(f32)
	fmt.Println(f64)

	// in this case random ranges equals to each other because i set the seed
	for i := 1; i <= 2; i++ {
		src := rand.NewSource(42) // create rand source with seed
		r := rand.New(src)        // set random generator with source

		fmt.Print(i, " random range: ")
		for j := 0; j < 5; j++ {
			r.Intn(100)
			fmt.Print(r.Intn(100), " ")
		}
		fmt.Println()
	}
}

// go run random.go
// 7
// 0.042564146
// 0.5597652292995556
// 1 random range: 87 50 45 76 43
// 2 random range: 87 50 45 76 43
