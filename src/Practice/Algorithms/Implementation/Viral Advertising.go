package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	shared := 5
	liked := shared / 2
	cumulative := liked
	for i := 1; i < n; i++ {
		shared = liked * 3
		liked = shared / 2
		cumulative += liked
	}
	fmt.Print(cumulative)
}
