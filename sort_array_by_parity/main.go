package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity_1([]int{3, 1, 2, 4}))
	fmt.Println(sortArrayByParity_2([]int{0, 2}))
}

func sortArrayByParity_1(A []int) []int {
	i := 0
	j := i + 1

	for j < len(A) {
		if A[i]%2 == 0 {
			i++
			j++
			continue
		}

		if A[j]%2 != 0 {
			j++
			continue
		}

		A[i], A[j] = A[j], A[i]
		i++
		j++
	}

	return A
}

func sortArrayByParity_2(A []int) []int {
	i := 0
	j := len(A) - 1

	for i < j {
		if A[i]%2 == 0 {
			i++
			continue
		}

		if A[j]%2 != 0 {
			j--
			continue
		}

		A[i], A[j] = A[j], A[i]
		i++
		j--

	}

	return A
}
