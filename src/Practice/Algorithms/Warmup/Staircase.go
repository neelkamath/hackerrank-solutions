package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	for row := 1; row <= n; row++ {
		line := ""
		for i := 0; i < n-row; i++ {
			line += " "
		}
		for i := 0; i < row; i++ {
			line += "#"
		}
		fmt.Println(line)
	}
}
