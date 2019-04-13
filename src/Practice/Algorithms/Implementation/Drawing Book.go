package main

import "fmt"

func main() {
	var n, p int
	fmt.Scanf("%d\n%d", &n, &p)

	if p == 1 || p == n || (p%2 == 0 && p == n-1) {
		fmt.Print(0)
	} else {
		front := p / 2
		var back int
		if p%2 == 0 {
			back = (n - p) / 2
		} else {
			back = (n - p + 1) / 2
		}
		if front < back {
			fmt.Print(front)
		} else {
			fmt.Print(back)
		}
	}
}
