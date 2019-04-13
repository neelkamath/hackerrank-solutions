package main

import "fmt"

func main() {
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N)
	for i := range A {
		fmt.Scanf("%d", &A[i])
	}
	for i := len(A) - 1; i >= 0; i-- {
		fmt.Printf("%d ", A[i])
	}
}
