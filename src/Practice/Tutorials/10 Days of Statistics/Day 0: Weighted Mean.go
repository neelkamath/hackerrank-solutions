package main

import "fmt"

func calculate(x, w []float64) float64 {
	var numerator, denominator float64
	for i, v := range x {
		numerator += v * w[i]
	}
	for _, v := range w {
		denominator += v
	}
	return numerator / denominator
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	x := make([]float64, n)
	for i := range x {
		fmt.Scanf("%f", &x[i])
	}
	w := make([]float64, n)
	for i := range w {
		fmt.Scanf("%f", &w[i])
	}
	fmt.Printf("%.1f\n", calculate(x, w))
}
