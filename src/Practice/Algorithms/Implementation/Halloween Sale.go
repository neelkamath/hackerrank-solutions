package main

import "fmt"

func main() {
	var p, d, m, s int
	fmt.Scanf("%d %d %d %d", &p, &d, &m, &s)

	games := 0
	for {
		s -= p
		if p > m {
			p -= d
			if p < m {
				p = m
			}
		}
		if s < 0 {
			break
		}
		games++
	}
	fmt.Println(games)
}
