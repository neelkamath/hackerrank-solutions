package main

import "fmt"

type data struct {
	students []int
	k        int
}

func (d data) String() string {
	return fmt.Sprintf("{students: %v, threshold: %v}", d.students, d.k)
}

func students(n int) []int {
	a := make([]int, n)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}
	return a
}

func nonPos(slice []int) int {
	var count int
	for _, v := range slice {
		if v <= 0 {
			count++
		}
	}
	return count
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	tests := make([]data, t)
	for index := range tests {
		var n, k int
		fmt.Scanf("%d %d", &n, &k)
		tests[index] = data{students(n), k}
	}

	for _, v := range tests {
		if nonPos(v.students) >= v.k {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
