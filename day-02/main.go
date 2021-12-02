package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const ()

func main() {
	regex := regexp.MustCompile(`^(forward|up|down) (\d+)$`)
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)
	aim := 0
	depth := 0
	horizontal := 0
	for fscanner.Scan() {
		match := regex.FindStringSubmatch(fscanner.Text())
		fmt.Println(match)
		if match[1] == "up" {
			i, _ := strconv.Atoi(match[2])
			aim -= i
		} else if match[1] == "down" {
			i, _ := strconv.Atoi(match[2])
			aim += i
		} else {
			i, _ := strconv.Atoi(match[2])
			horizontal += i
			depth += aim * i
		}
	}
	fmt.Println(depth * horizontal)
}
