package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	ar := make([]int, n)
	for i := range ar {
		fmt.Scanf("%d", &ar[i])
	}

	socks := make(map[int]int)
	for _, sock := range ar {
		count := socks[sock]
		count++
		socks[sock] = count
	}
	pairs := 0
	for _, count := range socks {
		pairs += count / 2
	}
	fmt.Print(pairs)
}
