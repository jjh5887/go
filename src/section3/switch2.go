package main

import "fmt"

func main() {
	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("2, 4, 6")
		fallthrough // go는 break가 기본이라 비활성화를 하려면 fallthrough
	case 1, 3, 5:
		fmt.Println("1, 3, 5")
		fallthrough
	default:
		fmt.Println("default")
	}

}
