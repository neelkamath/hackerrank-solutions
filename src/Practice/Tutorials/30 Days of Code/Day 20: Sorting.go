package main

import "fmt"

func sort(ints []int) int {
	totalSwaps := 0
	endPos := len(ints) - 1
	for endPos > 0 {
		swaps := 0
		for i, v := range ints[:endPos] {
			if v > ints[i+1] {
				ints[i], ints[i+1] = ints[i+1], ints[i]
				swaps++
				endPos = i
			}
		}
		if swaps == 0 {
			break
		}
		totalSwaps += swaps
	}
	return totalSwaps
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}
	fmt.Println("Array is sorted in", sort(a), "swaps.")
	fmt.Println("First Element:", a[0])
	fmt.Println("Last Element:", a[len(a)-1])
}
