package main

import "fmt"

func factorial(N int) int {
	if N <= 1 {
		return 1
	}
	return N * factorial(N-1)
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	fmt.Println(factorial(N))
}
