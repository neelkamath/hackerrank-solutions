package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	nums := make([]int, n)
	for i := range nums {
		fmt.Scanf("%d", &nums[i])
	}

	max := nums[0]
	for _, v := range nums[1:] {
		if v > max {
			max = v
		}
	}

	count := 0
	for _, v := range nums {
		if v == max {
			count++
		}
	}

	fmt.Print(count)
}
