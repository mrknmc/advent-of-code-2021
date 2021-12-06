package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func sum_board(board [][]string) int {
	sum := 0
	for x := range board {
		for y := range board[0] {
			if board[x][y] != "" {
				num, _ := strconv.Atoi(board[x][y])
				sum += num
			}
		}
	}
	return sum
}

func main() {
	file, _ := os.Open("data.txt")
	fscanner := bufio.NewScanner(file)

	fscanner.Scan()
	nums := strings.Split(fscanner.Text(), ",")

	space_re := regexp.MustCompile(`\s+`)

	boards := make([][][]string, 0)

	board := make([][]string, 0)
	for fscanner.Scan() {
		text := fscanner.Text()
		if text == "" {
			if len(board) > 0 {
				boards = append(boards, board)
				board = make([][]string, 0)
			}
			continue
		}
		row := space_re.Split(strings.Trim(text, " "), -1)
		board = append(board, row)
	}

	for _, num := range nums {
		for index, board := range boards {
			if len(board) == 0 {
				// bingoed already
				continue
			}
			for x := range board {
				for y := range board[0] {
					if num == board[x][y] {
						board[x][y] = ""
					}
				}
			}
			for x := range board {
				all_y := true
				for y := range board[0] {
					if board[x][y] != "" {
						all_y = false
						break
					}
				}
				if all_y {
					boards[index] = [][]string{}
					fmt.Println(sum_board(board), num)
				}
			}
			for y := range board[0] {
				all_x := true
				for x := range board {
					if board[x][y] != "" {
						all_x = false
						break
					}
				}
				if all_x {
					boards[index] = [][]string{}
					fmt.Println(sum_board(board), num)
				}
			}
		}
	}

}
