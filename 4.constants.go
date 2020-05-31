package main

import (
	"fmt"
	"math"
)

const s string = "Pradyumna"

func main() {
	fmt.Println(s)

	const n = 50000
	fmt.Println(n)

	const m = 5e20 / n
	// Const has no types until given one
	fmt.Println(m)
	fmt.Println(int64(m))

	// Type based on context
	fmt.Println(math.Sin(m))
}
