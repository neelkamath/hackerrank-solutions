package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var s string
	fmt.Scanf("%s", &s)

	valleys := 0
	altitude := 0
	for _, char := range s {
		switch char {
		case 'U':
			altitude++
		case 'D':
			if altitude == 0 {
				valleys++
			}
			altitude--
		}
	}
	fmt.Print(valleys)
}
