package main

import (
	"fmt"
	"math"
)

func main() {
	var mealCost float64
	var tipPercent, taxPercent int
	fmt.Scanf("%f\n%d\n%d", &mealCost, &tipPercent, &taxPercent)
	tip := mealCost * float64(tipPercent) / 100
	tax := mealCost * float64(taxPercent) / 100
	fmt.Println(math.Round(mealCost + tip + tax))
}
