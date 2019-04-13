package main

import (
	"fmt"
	"strconv"
	"strings"
)

func rangeSlice(lower, upper int) []int {
	slice := make([]int, upper-lower+1)
	for i, v := 0, lower; v <= upper; i, v = i+1, v+1 {
		slice[i] = v
	}
	return slice
}

func kaprekars(slice []int) []int {
	var nums []int
	for _, v := range slice {
		if isKaprekar(v) {
			nums = append(nums, v)
		}
	}
	return nums
}

func isKaprekar(n int) bool {
	d := len(strconv.Itoa(n))
	sq := strconv.Itoa(n * n)
	var l, r int
	switch len(sq) {
	case 2 * d:
		l, _ = strconv.Atoi(sq[:d])
		r, _ = strconv.Atoi(sq[d:])
	case 2*d - 1:
		l, _ = strconv.Atoi(sq[:d-1])
		r, _ = strconv.Atoi(sq[d-1:])
	}
	return l+r == n
}

func stringify(slice []int) []string {
	strings := make([]string, len(slice))
	for i, v := range slice {
		strings[i] = strconv.Itoa(v)
	}
	return strings
}

func main() {
	var p, q int
	fmt.Scanf("%d\n%d", &p, &q)

	slice := rangeSlice(p, q)
	if nums := kaprekars(slice); len(nums) == 0 {
		fmt.Println("INVALID RANGE")
	} else {
		fmt.Println(strings.Join(stringify(nums), " "))
	}
}
