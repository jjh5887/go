package main

import "fmt"

func main() {
	// skip
	const (
		_ = iota
		A
		B
		C
	)

	const (
		_ = iota + 0.75*2
		DEFAULT
		_
		GOLD
	)

	fmt.Println(A, B, C)
	fmt.Println(DEFAULT, GOLD)
}
