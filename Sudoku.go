package main

import (
	"fmt"
)

/*const (
	s = 9 //length of sudoku
	b = 3 //length of small box
)*/
func inRange(board [][]int) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] > 0 && board[i][j] <= 9 {
				return true
			}
		}
	}
	return false
}

func smallBox(board [][]int, s int) bool {
	for i := 0; i < s-2; i += 3 {
		for j := 0; j < s-2; j += 3 {
			currVal := board[i][j]
			for k := j + 1; k < s; k++ {
				if currVal == board[i][k] {
					return false
				}
			}
			for k := i + 1; k < s; k++ {
				if currVal == board[k][j] {
					return false
				}
			}
		}
	}
	return true
}

func validSudoku(board [][]int, s int) bool {
	if !inRange(board) {
		return false
	}
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			currVal := board[i][j]
			for k := j + 1; k < s; k++ {
				if currVal == board[i][k] {
					return false
				}
			}
			for k := i + 1; k < s; k++ {
				if currVal == board[k][j] {
					return false
				}
			}
			//count:=0

		}
	}
	if !smallBox(board, s) {
		return false
	}
	return true
}
func main() {
	board := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	if validSudoku(board, 9) {
		fmt.Println("valid")
	} else {
		fmt.Println("Invalid")
	}
}
