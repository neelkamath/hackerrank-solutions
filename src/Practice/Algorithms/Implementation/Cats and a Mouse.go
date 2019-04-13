package main

import "fmt"

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type postition struct {
	x, y, z int
}

func main() {
	var q int
	fmt.Scanf("%d", &q)

	queries := make([]postition, q)
	for i := range queries {
		fmt.Scanf("%d %d %d", &queries[i].x, &queries[i].y, &queries[i].z)
	}

	for _, v := range queries {
		xz := abs(v.x - v.z)
		yz := abs(v.y - v.z)
		switch {
		case xz == yz:
			fmt.Println("Mouse C")
		case yz > xz:
			fmt.Println("Cat A")
		case xz > yz:
			fmt.Println("Cat B")
		}
	}
}
