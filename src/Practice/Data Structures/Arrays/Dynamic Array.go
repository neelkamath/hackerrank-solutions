package main

import "fmt"

func main() {
	var N, Q int
	fmt.Scanf("%d %d", &N, &Q)

	seqList := make([][]int, N)
	lastAnswer := 0
	var vals []int
	for i := 0; i < Q; i++ {
		var t, x, y int
		fmt.Scanf("%d %d %d", &t, &x, &y)
		index := (x ^ lastAnswer) % N
		seq := seqList[index]
		switch t {
		case 1:
			seqList[index] = append(seq, y)
		case 2:
			lastAnswer = seq[y%len(seq)]
			vals = append(vals, lastAnswer)
		}
	}

	for i, v := range vals {
		fmtStr := "%d"
		if i != len(vals)-1 {
			fmtStr += "\n"
		}
		fmt.Printf(fmtStr, v)
	}
}
