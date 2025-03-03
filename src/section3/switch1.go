package main

import (
	"fmt"
)

func main() {
	a := -7
	switch {
	case a < 0:
		fmt.Println("a is negative")
	case a == 0:
		fmt.Println("a is zero")
	case a > 0:
		fmt.Println("a is positive")
	}

	switch b := 25; {
	case b < 0:
		fmt.Println("b is negative")
	case b == 0:
		fmt.Println("b is zero")
	case b > 0:
		fmt.Println("b is positive")
	}

	switch c := "go"; c {
	case "go":
		fmt.Println("c is go")
	case "java":
		fmt.Println("c is java")
	default:
		fmt.Println("c is unknown")
	}

	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("c is golang")
	case "java":
		fmt.Println("c is java")
	default:
		fmt.Println("c is unknown")
	}

	switch i, j := 1, 2; {
	case i < j:
		fmt.Println("i is less than j")
	case i == j:
		fmt.Println("i is equal to j")
	case i > j:
		fmt.Println("i is greater than j")
	}

}
