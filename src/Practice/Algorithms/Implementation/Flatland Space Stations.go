package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	stations := make([]int, m)
	for i := range stations {
		fmt.Scanf("%d", &stations[i])
	}
	maxDist := 0
	for city := 0; city < n; city++ {
		closest := int(math.Abs(float64(city - stations[0])))
		for _, station := range stations[1:] {
			dist := int(math.Abs(float64(city - station)))
			if dist < closest {
				closest = dist
			}
		}
		if closest > maxDist {
			maxDist = closest
		}
	}
	fmt.Println(maxDist)
}
