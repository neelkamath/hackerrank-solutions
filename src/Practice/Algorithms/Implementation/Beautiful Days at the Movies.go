package main

import (
	"fmt"
	"math"
)

func digits(num int) []int {
	var digits []int
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}
	return digits
}

func rev(num int) int {
	reversed := 0
	nums := digits(num)
	for i, v := range nums {
		reversed += v * int(math.Pow10(len(nums)-i-1))
	}
	return reversed
}

func beautiful(day, divisor int) bool {
	quotient := float32(day-rev(day)) / float32(divisor)
	return float32(int(quotient)) == quotient
}

func main() {
	var i, j, k int
	fmt.Scanf("%d %d %d", &i, &j, &k)

	days := 0
	for day := i; day <= j; day++ {
		if beautiful(day, k) {
			days++
		}
	}
	fmt.Print(days)
}
