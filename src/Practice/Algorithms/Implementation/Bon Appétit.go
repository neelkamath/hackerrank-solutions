package main

import "fmt"

func sum(nums []int) int {
	val := 0
	for _, num := range nums {
		val += num
	}
	return val
}

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	bill := make([]int, n)
	for i := range bill {
		fmt.Scanf("%d", &bill[i])
	}

	var b int
	fmt.Scanf("%d", &b)

	incorrect := sum(bill) / 2
	correct := (sum(bill) - bill[k]) / 2
	if b == correct {
		fmt.Print("Bon Appetit")
	} else {
		fmt.Print(incorrect - correct)
	}
}
