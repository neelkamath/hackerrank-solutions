package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	a := make([]int, n)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}

	pairs := 0
	for i, first := range a {
		for j, second := range a {
			if i < j && (first+second)%k == 0 {
				pairs++
			}
		}
	}
	fmt.Print(pairs)
}
