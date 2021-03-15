package main

import (
	"fmt"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))
	fmt.Println(convert("AB", 1))
	fmt.Println(convert("ABC", 1))
}

const (
	moveDown = "down"
	moveUp   = "up"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	zigzag := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		zigzag[i] = make([]string, len(s))
	}

	// x is row index, y is col index
	x := -1
	y := 0

	// counter is from 0 to len(s)
	counter := 0

	// decide moving down or moving up
	state := moveDown

	for counter < len(s) {
		if state == moveDown {
			x++
			if x == numRows-1 {
				state = moveUp
			}
		} else {
			x--

			y++
			if x == 0 {
				state = moveDown
			}
		}

		// fmt.Println(x, y, string(s[counter]), counter)

		zigzag[x][y] = string(s[counter])
		counter++

	}

	// combine zigzag
	result := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(s); j++ {
			result += zigzag[i][j]
		}
	}

	return result
}
