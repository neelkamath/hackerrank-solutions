package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N)
	for i := range A {
		fmt.Scanf("%d", &A[i])
	}

	for i := N - 1; i > -1; i-- {
		fmt.Printf("%d ", A[i])
	}
}
