package main

import "fmt"

func main() {
	var (
		name      string = "John"
		height    int    = 170
		weight    float32
		isRunning bool
	)

	height = 180
	weight = 70.5
	isRunning = true

	fmt.Println(name, height, weight, isRunning)
}
