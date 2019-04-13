package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func printArr(arr []int32) {
	length := len(arr)
	for i := 0; i < length; i++ {
		if i == length-1 {
			fmt.Println(arr[i])
		} else {
			fmt.Printf("%d ", arr[i])
		}
	}
}

// Complete the insertionSort2 function below.
func insertionSort2(n int32, arr []int32) {
	for i := int32(1); i < n; i++ {
		var shift int32
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[i] {
				shift++
			}
		}
		value := arr[i]
		index := i - shift
		for k := i; k >= index+1; k-- {
			arr[k] = arr[k-1]
		}
		arr[index] = value
		printArr(arr)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	insertionSort2(n, arr)
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
