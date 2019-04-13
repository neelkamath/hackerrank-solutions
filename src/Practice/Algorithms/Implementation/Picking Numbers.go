package main

import "fmt"

func contains(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

func unique(slice []int) [][]int {
	var slices [][]int
	var added []int
	for _, v := range slice {
		if !contains(added, v) {
			slices = append(slices, []int{v})
			added = append(added, v)
		}
	}
	return slices
}

func largestLen(slices [][]int) int {
	var largest int
	for _, v := range slices {
		length := len(v)
		if length > largest {
			largest = length
		}
	}
	return largest
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}

	slices := unique(a)

	for _, v := range a {
		for i, slice := range slices {
			diff := slice[0] - v
			if diff == 1 || diff == 0 {
				slices[i] = append(slice, v)
			}
		}
	}

	fmt.Print(largestLen(slices) - 1)
}
