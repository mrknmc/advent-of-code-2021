package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)
	fs1 := make([]int64, 0)
	for fscanner.Scan() {
		i, _ := strconv.ParseInt(fscanner.Text(), 10, 64)
		fs1 = append(fs1, i)
	}
	count := 0
	var last_sum int64
	for i := 2; i < len(fs1); i++ {
		sum := fs1[i-1] + fs1[i-2] + fs1[i]
		if sum > last_sum && last_sum != 0 {
			count++
		}
		last_sum = sum
	}
	fmt.Println(count)
}
