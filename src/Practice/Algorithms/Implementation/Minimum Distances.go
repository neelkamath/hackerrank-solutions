package main

import "fmt"

type position struct {
	index, value int
}

func getPositions(slice []int) [][]position {
	var positions [][]position
	for i, v := range slice {
		if count(slice, v) > 1 {
			if index, ok := in(v, positions); ok {
				positions[index] = append(positions[index], position{i, v})
			} else {
				positions = append(positions, []position{position{i, v}})
			}
		}
	}
	return positions
}

func count(slice []int, num int) int {
	seen := 0
	for _, v := range slice {
		if v == num {
			seen++
		}
	}
	return seen
}

func in(num int, positions [][]position) (index int, ok bool) {
	for i, pos := range positions {
		for _, v := range pos {
			if v.value == num {
				return i, true
			}
		}
	}
	return
}

func getDistances(posSlices [][]position) [][]int {
	distances := make([][]int, len(posSlices))
	for sliceIndex, posSlice := range posSlices {
		for i := 1; i < len(posSlice); i++ {
			distances[sliceIndex] = append(distances[sliceIndex], posSlice[i].index-posSlice[i-1].index)
		}
	}
	return distances
}

func shorten(slices [][]int) []int {
	slice := make([]int, len(slices))
	for i, v := range slices {
		slice[i] = shortest(v)
	}
	return slice
}

func shortest(slice []int) int {
	if len(slice) == 0 {
		return -1
	}

	val := slice[0]
	for _, v := range slice[1:] {
		if v < val {
			val = v
		}
	}
	return val
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := range a {
		fmt.Scanf("%d", &a[i])
	}

	positions := getPositions(a)
	distances := getDistances(positions)
	shortestDistances := shorten(distances)
	fmt.Print(shortest(shortestDistances))
}
