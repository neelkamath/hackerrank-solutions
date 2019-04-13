package main

import (
	"fmt"
	"sort"
)

func sum(nums []int) int {
	val := 0
	for _, v := range nums {
		val += v
	}
	return val
}

func main() {
	nums := make([]int, 5)
	for i := range nums {
		fmt.Scanf("%d", &nums[i])
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Print(sum(nums[:4]), sum(nums[1:]))
}
