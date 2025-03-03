package main

import "fmt"

func main() {
	// 짧은 선언
	// 함수 안에서만 사용
	shortVar1 := 3
	shortVar2 := "Test"

	if i := 10; i > 5 {
		fmt.Println("Short Var Test Success!")
	}

	fmt.Println(shortVar1, shortVar2)
}
