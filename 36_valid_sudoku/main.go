package main

import "fmt"

func main() {
	fmt.Println("36. Valid Sudoku")
}

func isValidSudoku(board [][]byte) bool {
	fmt.Println(string(board[0][0]))

	if !ruleHorizontal(board) {
		return false
	}

	if !ruleVertical(board) {
		return false
	}

	if !ruleBlock(board) {
		return false
	}

	return true
}

func ruleHorizontal(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		marks := make(map[string]bool)

		for j := 0; j < 9; j++ {
			if !isValidMarks(marks, board[i][j]) {
				return false
			}
		}
	}

	return true
}

func ruleVertical(board [][]byte) bool {
	for j := 0; j < 9; j++ {
		marks := make(map[string]bool)

		for i := 0; i < 9; i++ {
			if !isValidMarks(marks, board[i][j]) {
				return false
			}
		}
	}

	return true
}

func ruleBlock(board [][]byte) bool {
	for base := 0; base < 9; base++ {
		marks := make(map[string]bool)

		i0 := base - base%3
		for i := i0; i < i0+3; i++ {
			j0 := base % 3 * 3
			for j := j0; j < j0+3; j++ {
				if !isValidMarks(marks, board[i][j]) {
					return false
				}
			}
		}
	}

	return true
}

func isValidMarks(marks map[string]bool, target byte) bool {
	if string(target) == "." {
		return true
	}

	if marks[string(target)] {
		return false
	}

	marks[string(target)] = true
	return true
}
