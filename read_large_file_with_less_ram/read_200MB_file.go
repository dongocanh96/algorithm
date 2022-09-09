package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("random.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, _ := file.Stat()
	chunkSize := int(info.Size() / 20)

	bufR := bufio.NewReaderSize(file, chunkSize)

	for i := range [20]int{} {
		reader := make([]byte, chunkSize)
		_, err := bufR.Read(reader)
		if err != nil {
			panic(err)
		}
		byteSlice := bytes.Split(reader, []byte(" "))
		sliceInt := make([]int64, len(byteSlice))
		for j := range byteSlice {
			sliceInt[j], _ = strconv.ParseInt(string(byteSlice[j]), 10, 64)
		}

		sortedSlice := sort(sliceInt)

		writeFile(i, sortedSlice)
	}
}

func writeFile(i int, slice []int64) {
	fname := fmt.Sprintf("file_%v.txt", i)
	f, err := os.Create(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	datawriter := bufio.NewWriter(f)
	for _, data := range slice {
		inputData := fmt.Sprintf("%v\n", data)
		_, _ = datawriter.WriteString(inputData)
	}
}

func sort(slice []int64) []int64 {
	if len(slice) < 2 {
		return slice
	}
	first := sort(slice[:len(slice)/2])
	second := sort(slice[len(slice)/2:])
	return merge(first, second)
}

func merge(a []int64, b []int64) []int64 {
	final := []int64{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
