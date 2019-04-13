package main

import "fmt"

func main() {
	var d1, d2, m1, m2, y1, y2 int
	fmt.Scanf("%d %d %d\n%d %d %d", &d2, &m2, &y2, &d1, &m1, &y1)
	earlierY := y2 < y1
	earlierM := y2 == y1 && m2 < m1
	earlierOrSameD := y2 == y1 && m2 == m1 && d2 <= d1
	switch {
	case earlierY || earlierM || earlierOrSameD:
		fmt.Println(0)
	case y2 == y1 && m2 == m1 && d2 > d1:
		fmt.Println(15 * (d2 - d1))
	case y2 == y1 && m2 > m1:
		fmt.Println(500 * (m2 - m1))
	default:
		fmt.Println(10000)
	}
}
