package main

import "fmt"

func shift(array []int) []int {
	return append(array, array[0])[1:]
}

func main() {
	var n, d int
	fmt.Scanf("%d %d", &n, &d)

	array := make([]int, n)
	for i := range array {
		fmt.Scanf("%d", &array[i])
	}

	for i := 0; i < d; i++ {
		array = shift(array)
	}
	for i, v := range array {
		str := "%d"
		if i != len(array)-1 {
			str += " "
		}
		fmt.Printf(str, v)
	}
}
