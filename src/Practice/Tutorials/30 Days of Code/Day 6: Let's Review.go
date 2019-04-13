package main

import "fmt"

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for i := 0; i < T; i++ {
		var S string
		fmt.Scanf("%s", &S)
		var evens, odds string
		for i := 0; i < len(S); i += 2 {
			evens += string(S[i])
		}
		for i := 1; i < len(S); i += 2 {
			odds += string(S[i])
		}
		fmt.Println(evens, odds)
	}
}
