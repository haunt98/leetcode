package main

func main() {}

func sortedSquares_1(A []int) []int {
	for i := range A {
		A[i] *= A[i]
	}

	insertionSort(A)

	return A
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j >= 1; j-- {
			if arr[j] >= arr[j-1] {
				break
			}

			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
