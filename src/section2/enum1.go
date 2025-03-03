package main

import "fmt"

func main() {
	const (
		JAN = 1
		FEB = 2
		MAR = 3
		APR = 4
		MAY = 5
	)

	fmt.Println(JAN, FEB, MAR, APR, MAY)

	const (
		A = iota + 1
		B
		C
		D
		E
	)

	fmt.Println(A, B, C, D, E)
}
