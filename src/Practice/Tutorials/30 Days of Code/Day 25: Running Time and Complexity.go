package main

import "fmt"

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n == 1 || n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 3; i <= n/3; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for i := 0; i < T; i++ {
		var n int
		fmt.Scanf("%d", &n)
		if isPrime(n) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}
