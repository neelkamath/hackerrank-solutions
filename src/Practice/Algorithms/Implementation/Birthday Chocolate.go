package main

import "fmt"

func sum(nums []int) int {
	val := 0
	for _, v := range nums {
		val += v
	}
	return val
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	s := make([]int, n)
	for i := range s {
		fmt.Scanf("%d", &s[i])
	}

	var d, m int
	fmt.Scanf("%d %d", &d, &m)

	ways := 0
	for i := 0; i < n; i++ {
		if i+m > len(s) {
			break
		}
		if d == sum(s[i:i+m]) {
			ways++
		}
	}
	fmt.Print(ways)
}
