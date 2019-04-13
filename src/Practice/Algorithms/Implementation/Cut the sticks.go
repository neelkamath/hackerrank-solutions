package main

import "fmt"

func smallest(slice []int) int {
	smallest := slice[0]
	for _, v := range slice[1:] {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}

func shorten(slice []int, cut int) []int {
	var shortened []int
	for _, v := range slice {
		length := v - cut
		if length != 0 {
			shortened = append(shortened, length)
		}
	}
	return shortened
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scanf("%d", &arr[i])
	}

	for len(arr) > 0 {
		fmt.Println(len(arr))
		arr = shorten(arr, smallest(arr))
	}
}
