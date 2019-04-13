package main

import "fmt"

func getLargest(arr []int32) int32 {
	largest := arr[0]
	for _, v := range arr[1:] {
		if v > largest {
			largest = v
		}
	}
	return largest
}

// Complete the countingSort function below.
func countingSort(arr []int32) []int32 {
	var sorted []int32
	if len(arr) > 0 {
		counts := make([]int32, getLargest(arr)+1)
		for _, v := range arr {
			counts[v]++
		}
		for index, value := range counts {
			part := make([]int32, value)
			for i := int32(0); i < value; i++ {
				part[i] = int32(index)
			}
			sorted = append(sorted, part...)
		}
	}
	return sorted
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int32, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	sorted := countingSort(arr)
	for _, v := range sorted {
		fmt.Printf("%d ", v)
	}
}
