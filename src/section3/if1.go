package main

import "fmt"

func main() {
	var a int = 20
	b := 20

	if a >= 15 {
		fmt.Println("a is greater than 15")
	}

	if b >= 25 {
		fmt.Println("b is greater than 25")
	}

	// 짧은 변수 범위
	if c := 10; c >= 15 {
		fmt.Println("c is greater than 15")
	}
	// error
	// c += 20

	if a >= 66 {
		fmt.Println("a is greater than 66")
	} else {
		fmt.Println("a is less than 66")
	}

	if a >= 30 {
		fmt.Println("a is greater than 30")
	} else if a >= 20 && a < 30 {
		fmt.Println("a is greater than 20 and less than 30")
	} else if a < 10 {
		fmt.Println("a is greater than 10")
	} else {
		fmt.Println("a is less than 10")
	}
}
