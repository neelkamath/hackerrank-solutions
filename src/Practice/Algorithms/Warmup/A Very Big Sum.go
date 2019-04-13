package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	nums := make([]int, n)
	for i := range nums {
		fmt.Scanf("%d", &nums[i])
	}

	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	fmt.Print(sum)
}
