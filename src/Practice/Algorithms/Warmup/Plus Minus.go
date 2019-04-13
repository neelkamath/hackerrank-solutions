package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	nums := make([]int, n)
	for i := range nums {
		fmt.Scanf("%d", &nums[i])
	}

	pos := 0
	neg := 0
	zeros := 0
	for _, v := range nums {
		switch {
		case v > 0:
			pos++
		case v < 0:
			neg++
		case v == 0:
			zeros++
		}
	}

	f := "%6f\n"
	fmt.Printf(f, float32(pos)/float32(n))
	fmt.Printf(f, float32(neg)/float32(n))
	fmt.Printf(f, float32(zeros)/float32(n))
}
