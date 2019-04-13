package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	scores := make([]int, n)
	for i := range scores {
		fmt.Scanf("%d", &scores[i])
	}

	var highestCount, lowestCount int
	highest := scores[0]
	lowest := scores[0]
	for _, v := range scores[1:] {
		switch {
		case v > highest:
			highestCount++
			highest = v
		case v < lowest:
			lowestCount++
			lowest = v
		}
	}
	fmt.Print(highestCount, lowestCount)
}
