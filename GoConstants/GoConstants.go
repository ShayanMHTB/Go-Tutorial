package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println("s:", s)
	
	const n = 500000000
	const d = 3e20 / n

	fmt.Println("d:", d)
	fmt.Println("d int64:", int64(d))
	fmt.Println("Sin(n):", math.Sin(n))
}
