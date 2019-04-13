package main

import "fmt"

func index(slice []int, num int) int {
	for i, v := range slice {
		if v == num {
			return i
		}
	}
	return -1
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	p := make([]int, n)
	for i := range p {
		fmt.Scanf("%d", &p[i])
	}

	for x := 1; x <= n; x++ {
		i := index(p, x) + 1
		fmt.Println(index(p, i) + 1)
	}
}
