package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		fmt.Println(getRow_1(i))
	}
}

// 1
// 1 1
// 1 2 1
// 1 3 3 1
func getRow_1(rowIndex int) []int {
	if rowIndex < 0 {
		return nil
	}

	if rowIndex == 0 {
		return []int{1}
	}

	prevRow := getRow_1(rowIndex - 1)

	lenCurRow := len(prevRow) + 1

	curRow := make([]int, lenCurRow)
	curRow[0] = 1
	curRow[len(curRow)-1] = 1
	for i := 1; i < lenCurRow-1; i++ {
		curRow[i] = prevRow[i-1] + prevRow[i]
	}

	return curRow
}
