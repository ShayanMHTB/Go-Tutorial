package main

import (
	"os"
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := os.Args[1]
	number, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Error during conversion.")
		return
	} else {
		if isPrime(number) == true {
			fmt.Println(number, "is a Prime.")
		} else {
			fmt.Println(number, "is not a Prime.")
		}
		return
	}
}

func isPrime(number int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(number)) + 1)); i++ {
		if number % i == 0 {
			return false
		}
	}
	return true
}
