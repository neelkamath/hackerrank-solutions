package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	var i2 uint64
	var d2 float64
	fmt.Scanf("%d\n%f", &i2, &d2)
	scanner.Scan()
	s2 := scanner.Text()
	fmt.Printf("%d\n%.1f\n%s%s\n", i+i2, d+d2, s, s2)
}
