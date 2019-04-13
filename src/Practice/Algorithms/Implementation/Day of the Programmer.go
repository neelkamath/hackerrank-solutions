package main

import "fmt"

func main() {
	var y int
	fmt.Scanf("%d", &y)

	m := 9
	var d int
	switch {
	case y > 1918:
		if y%400 == 0 || (y%4 == 0 && y%100 != 0) {
			d = 12
		} else {
			d = 13
		}
	case y < 1918:
		if y%4 == 0 {
			d = 12
		} else {
			d = 13
		}
	case y == 1918:
		if y%400 == 0 || (y%4 == 0 && y%100 != 0) {
			d = 25
		} else {
			d = 26
		}
	}
	fmt.Printf("%02d.%02d.%d", d, m, y)
}
