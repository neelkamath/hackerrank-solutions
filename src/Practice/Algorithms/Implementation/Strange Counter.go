package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	cycle := 3
	value := cycle
	time := 1
	for value == cycle && time+cycle <= t {
		time += cycle
		cycle *= 2
		value = cycle
	}
	value -= t - time
	fmt.Println(value)
}
