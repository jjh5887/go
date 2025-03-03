package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(">", i)
	}

	//for {
	//	fmt.Println("infinite loop")
	//}

	loc := []string{"a", "b", "c"}
	// idx 무시하려면 _ 로
	for idx, data := range loc {
		fmt.Println(idx, data)
	}

	sum := 0
	for sum < 100 {
		sum += 1
	}
	fmt.Println("sum: ", sum)

	sum2 := 0
	for {
		if sum2 > 100 {
			break
		}
		sum2 += 1
	}
	fmt.Println("sum2: ", sum2)

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("i: ", i)
	}

	for i, j := 0, 0; i < 10; i, j = i+1, j+10 {
		fmt.Println(i, j)
	}
}
