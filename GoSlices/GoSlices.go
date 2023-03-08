package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[0] = "b"
	s[0] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)
	
	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoDim := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDim[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDim[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoDim)
}
