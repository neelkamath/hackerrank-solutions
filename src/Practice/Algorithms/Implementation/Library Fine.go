package main

import "fmt"

func main() {
	var d1, m1, y1, d2, m2, y2 int
	fmt.Scanf("%d %d %d\n%d %d %d", &d1, &m1, &y1, &d2, &m2, &y2)
	switch {
	case d1 > d2 && m1 == m2 && y1 == y2:
		fmt.Print(15 * (d1 - d2))
	case m1 > m2 && y1 == y2:
		fmt.Print(500 * (m1 - m2))
	case y1 > y2:
		fmt.Print(10000)
	default:
		fmt.Print(0)
	}
}
