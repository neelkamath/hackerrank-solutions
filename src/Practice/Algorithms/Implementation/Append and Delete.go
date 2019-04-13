package main

import "fmt"

func rmLast(s string) string {
	return s[:len(s)-1]
}

func main() {
	var s, t string
	var k int
	fmt.Scanf("%s\n%s\n%d", &s, &t, &k)

	if k >= len(s)+len(t) {
		fmt.Println("Yes")
		return
	}
	for k > 0 {
		switch {
		case len(t) > len(s):
			if t[:len(s)] == s {
				s += string(t[len(s)])
			} else {
				s = rmLast(s)
			}
		case len(s) > len(t):
			s = rmLast(s)
		case len(s) == len(t):
			s = rmLast(s)
		}
		k--
	}
	if s == t {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
