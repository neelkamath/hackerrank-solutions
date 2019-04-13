package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d", &n, &k)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scanf("%d", &arr[i])
	}
	var spProblems, page int
	for _, problems := range arr {
		problem := 1
		for problem <= problems {
			page++
			for i := 0; i < k && problem <= problems; i, problem = i+1, problem+1 {
				if problem == page {
					spProblems++
				}
			}
		}
	}
	fmt.Println(spProblems)
}
