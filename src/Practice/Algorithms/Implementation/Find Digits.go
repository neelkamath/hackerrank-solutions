package main

import "fmt"

func getDigits(num int) []int {
	var slice []int
	for num > 0 {
		slice = append(slice, num%10)
		num /= 10
	}
	return slice
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	tests := make([]int, t)
	for i := range tests {
		fmt.Scanf("%d", &tests[i])
	}

	for _, test := range tests {
		count := 0
		for _, v := range getDigits(test) {
			if v != 0 && test%v == 0 {
				count++
			}
		}
		fmt.Println(count)
	}
}
