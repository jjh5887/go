package main

import "fmt"

func main() {
	const a int = 10
	const b string = "Hello"

	// 함수는 불가능
	//const c = getHeight()

	const d = 10.5
	const e = true

	fmt.Println(a, b, d, e)
}

func getHeight() int {
	return 10
}
