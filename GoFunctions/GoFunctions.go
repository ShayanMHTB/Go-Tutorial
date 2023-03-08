package main

import "fmt"

func add2Nums(a int, b int) int {
	return a + b
}

func add3Nums(a, b, c int) int {
	return a + b + c
}

func main() {
	res := add2Nums(1, 2)
	fmt.Println("1 + 2 =", res)

	res = add3Nums(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)
}
