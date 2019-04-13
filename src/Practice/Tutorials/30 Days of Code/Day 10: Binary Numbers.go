package main

import (
	"fmt"
	"strconv"
)

func binConv(n int) string {
	bin := ""
	for n > 0 {
		bin = strconv.Itoa(n%2) + bin
		n /= 2
	}
	return bin
}

func countOnes(s string) int {
	var currMax, max int
	for _, v := range s {
		if v == '1' {
			currMax += 1
		} else {
			if currMax > max {
				max = currMax
			}
			currMax = 0
		}
	}
	if currMax > max {
		max = currMax
	}
	return max
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(countOnes(binConv(n)))
}
