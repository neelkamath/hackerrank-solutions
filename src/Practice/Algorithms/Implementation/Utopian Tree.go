package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	n := make([]int, t)
	for i := range n {
		fmt.Scanf("%d", &n[i])
	}

	for _, v := range n {
		h := 1
		for i := 0; i < v; i++ {
			if i%2 == 1 {
				h++
			} else {
				h *= 2
			}
		}
		fmt.Println(h)
	}
}
