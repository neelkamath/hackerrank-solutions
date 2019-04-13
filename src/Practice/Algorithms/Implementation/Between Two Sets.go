package main

import (
	"fmt"
	"sort"
)

func input(size int) []int {
	nums := make([]int, size)
	for i := range nums {
		fmt.Scanf("%d", &nums[i])
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}

func factors(num int) []int {
	var f []int
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			f = append(f, i)
		}
	}
	return f
}

func contains(nums []int, num int) bool {
	for _, v := range nums {
		if v == num {
			return true
		}
	}
	return false
}

func factorOf(x, y int) bool {
	return contains(factors(y), x)
}

func factorsOf(nums []int, num int) bool {
	for _, v := range nums {
		if !factorOf(v, num) {
			return false
		}
	}
	return true
}

func factorOfAll(num int, nums []int) bool {
	for _, v := range nums {
		if !factorOf(num, v) {
			return false
		}
	}
	return true
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	a := input(n)
	b := input(m)

	count := 0
	for begin, end := a[len(a)-1], b[0]; begin <= end; begin++ {
		if factorsOf(a, begin) && factorOfAll(begin, b) {
			count++
		}
	}
	fmt.Print(count)
}
