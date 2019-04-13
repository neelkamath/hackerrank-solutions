package main

import "fmt"

func max(slice []int) int {
	var largest int
	for _, v := range slice {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	height := make([]int, n)
	for i := range height {
		fmt.Scanf("%d", &height[i])
	}

	potions := max(height) - k
	if potions < 0 {
		fmt.Print(0)
	} else {
		fmt.Print(potions)
	}
}
