package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := range nums {
		fmt.Scanf("%d", &nums[i])
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Print(sum)
}
