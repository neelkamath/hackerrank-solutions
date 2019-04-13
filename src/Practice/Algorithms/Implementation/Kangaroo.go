package main

import "fmt"

func main() {
	var x1, v1, x2, v2 int
	fmt.Scanf("%d %d %d %d", &x1, &v1, &x2, &v2)

	if v1 <= v2 {
		fmt.Print("NO")
		return
	}
	for {
		x1 += v1
		x2 += v2
		if x1 > x2 && v1 > v2 {
			fmt.Print("NO")
			return
		}
		if x1 == x2 {
			fmt.Print("YES")
			return
		}
	}
}
