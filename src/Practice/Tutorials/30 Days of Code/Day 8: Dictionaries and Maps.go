package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	entries := make(map[string]int)
	for i := 0; i < n; i++ {
		var name string
		var phoneNumber int
		fmt.Scanf("%s %d", &name, &phoneNumber)
		entries[name] = phoneNumber
	}
	for {
		var name string
		fmt.Scanf("%s", &name)
		if name == "" {
			break
		}
		if phoneNumber, ok := entries[name]; ok {
			fmt.Printf("%s=%d\n", name, phoneNumber)
		} else {
			fmt.Println("Not found")
		}
	}
}
