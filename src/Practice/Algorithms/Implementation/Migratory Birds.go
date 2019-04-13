package main

import "fmt"

func get(key, value *int, vals map[int]int) {
	for k, v := range vals {
		*key = k
		*value = v
		delete(vals, k)
		return
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Scanf("%d", &arr[i])
	}

	frequencies := make(map[int]int)
	for _, id := range arr {
		count := frequencies[id]
		count++
		frequencies[id] = count
	}

	var id, count int
	get(&id, &count, frequencies)
	for k, v := range frequencies {
		if v > count || (v == count && k < id) {
			id = k
			count = v
		}
	}
	fmt.Print(id)
}
