package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	switch {
	case n%2 == 1:
		fmt.Println("Weird")
	case n >= 2 && n <= 5:
		fmt.Println("Not Weird")
	case n >= 6 && n <= 20:
		fmt.Println("Weird")
	default:
		fmt.Println("Not Weird")
	}
}
