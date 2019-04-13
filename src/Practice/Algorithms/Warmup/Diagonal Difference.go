package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var nums [][]int
	for i := 0; i < n; i++ {
		nums = append(nums, make([]int, n))
	}
	for i := range nums {
		for j := range nums[i] {
			fmt.Scanf("%d", &nums[i][j])
		}
	}

	primary := 0
	for i := range nums {
		primary += nums[i][i]
	}
	secondary := 0
	for i := range nums {
		secondary += nums[i][len(nums)-i-1]
	}

	diff := primary - secondary
	if diff < 0 {
		diff = -diff
	}
	fmt.Print(diff)
}
