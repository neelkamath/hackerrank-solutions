package main

import "fmt"

func countMode(slice []int) int {
	m := getMap(slice)
	modeCount := first(m)
	for _, count := range m {
		if count > modeCount {
			modeCount = count
		}
	}
	return modeCount
}

func first(m map[int]int) int {
	for k, v := range m {
		delete(m, k)
		return v
	}
	return 0
}

func getMap(slice []int) map[int]int {
	m := make(map[int]int)
	for _, v := range slice {
		count := m[v]
		count++
		m[v] = count
	}
	return m
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scanf("%d", &arr[i])
	}

	count := countMode(arr)
	fmt.Println(len(arr) - count)
}
