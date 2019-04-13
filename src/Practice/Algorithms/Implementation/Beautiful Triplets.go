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

func main() {
	var n, d int
	fmt.Scanf("%d %d", &n, &d)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scanf("%d", &arr[i])
	}

	count := 0
	for _, v := range arr {
		if contains(arr, v+2*d) && contains(arr, v+d) {
			count++
		}
	}
	fmt.Println(count)
}
