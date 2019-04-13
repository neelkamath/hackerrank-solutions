package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func intify(int32s []int32) []int {
	var ints []int
	for _, v := range int32s {
		ints = append(ints, int(v))
	}
	return ints
}

// Complete the marcsCakewalk function below.
func marcsCakewalk(calorie []int32) int64 {
	calories := intify(calorie)
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	var miles int64
	for j := 0; j < len(calorie); j++ {
		miles += int64(math.Pow(2, float64(j)) * float64(calories[j]))
	}
	return miles
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	calorieTemp := strings.Split(readLine(reader), " ")

	var calorie []int32

	for i := 0; i < int(n); i++ {
		calorieItemTemp, err := strconv.ParseInt(calorieTemp[i], 10, 64)
		checkError(err)
		calorieItem := int32(calorieItemTemp)
		calorie = append(calorie, calorieItem)
	}

	result := marcsCakewalk(calorie)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
