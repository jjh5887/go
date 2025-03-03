package main

import "fmt"

func main() {
Loop1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Printf("%d %d\n", i, j)
		}
	}
}
