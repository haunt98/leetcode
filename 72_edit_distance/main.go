package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("intention", "execution"))
	fmt.Println(minDistance("dinitrophenylhydrazine", "benzalphenylhydrazone"))
}

func minDistance(word1 string, word2 string) int {
	memo := make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		memo[i] = make([]int, len(word2))
		for j := 0; j < len(word2); j++ {
			memo[i][j] = -1
		}
	}

	return minDistanceWithMemo(word1, len(word1)-1, word2, len(word2)-1, memo)
}

// return min distance of word1[0:i+1] and word2[0:j+1]
func minDistanceRecursive(word1 string, i int, word2 string, j int) int {
	// insert all word2[0:j+1]
	if i == -1 {
		return j + 1
	}

	// delete all word1[0:i+1]
	if j == -1 {
		return i + 1
	}

	// alredy match this time, move forward
	if word1[i] == word2[j] {
		return minDistanceRecursive(word1, i-1, word2, j-1)
	}

	// word1 = abc
	// word2 = abd
	// insert mean word1 become abcd for abc[d] match ab[d]
	// then it become abc and ab for next min distance
	insertDistance := minDistanceRecursive(word1, i, word2, j-1) + 1

	// word1 = abc
	// word2 = abd
	// update mean word1 ab[c] -> ab[d]
	updateDistance := minDistanceRecursive(word1, i-1, word2, j-1) + 1

	// word1 = abc
	// word2 = abd
	// delete mean word1 ab[c] -> ab
	deleteDistance := minDistanceRecursive(word1, i-1, word2, j) + 1

	return minOf2(minOf2(insertDistance, updateDistance), deleteDistance)
}

// return min distance of word1[0:i+1] and word2[0:j+1]
func minDistanceWithMemo(word1 string, i int, word2 string, j int, memo [][]int) int {
	// fmt.Println(memo)

	// insert all word2[0:j+1]
	if i == -1 {
		return j + 1
	}

	// delete all word1[0:i+1]
	if j == -1 {
		return i + 1
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	// alredy match this time, move forward
	if word1[i] == word2[j] {
		memo[i][j] = minDistanceWithMemo(word1, i-1, word2, j-1, memo)
		return memo[i][j]
	}

	// word1 = abc
	// word2 = abd
	// insert mean word1 become abcd for abc[d] match ab[d]
	// then it become abc and ab for next min distance
	insertDistance := minDistanceWithMemo(word1, i, word2, j-1, memo) + 1

	// word1 = abc
	// word2 = abd
	// update mean word1 ab[c] -> ab[d]
	updateDistance := minDistanceWithMemo(word1, i-1, word2, j-1, memo) + 1

	// word1 = abc
	// word2 = abd
	// delete mean word1 ab[c] -> ab
	deleteDistance := minDistanceWithMemo(word1, i-1, word2, j, memo) + 1

	memo[i][j] = minOf3(insertDistance, updateDistance, deleteDistance)
	return memo[i][j]
}

func minOf2(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func minOf3(a, b, c int) int {
	return minOf2(minOf2(a, b), c)
}
