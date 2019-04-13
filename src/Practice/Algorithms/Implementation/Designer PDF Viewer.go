package main

import "fmt"

func max(m map[string]int) int {
	var largest int
	for _, v := range m {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	h := make(map[string]int)
	for _, v := range "abcdefghijklmnopqrstuvwxyz" {
		var height int
		fmt.Scanf("%d", &height)
		h[string(v)] = height
	}

	var word string
	fmt.Scanf("%s", &word)

	heights := make(map[string]int)
	for _, v := range word {
		str := string(v)
		heights[str] = h[str]
	}

	fmt.Print(max(heights) * len(word))
}
