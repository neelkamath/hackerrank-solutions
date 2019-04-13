package main

import "fmt"

func main() {
	var h, m, s int
	var meridiem string
	fmt.Scanf("%d:%d:%d%s", &h, &m, &s, &meridiem)

	if h == 12 && m == 0 && s == 0 && meridiem == "AM" {
		h = 0
	} else if meridiem == "AM" && h == 12 {
		h -= 12
	} else if meridiem == "PM" && h != 12 {
		h += 12
	}
	fmt.Printf("%02d:%02d:%02d", h, m, s)
}
