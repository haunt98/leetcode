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

func sortedSquares_2(A []int) []int {
	positiveIndex := 0

	for i := 0; i < len(A); i++ {
		if A[i] >= 0 {
			positiveIndex = i
			break
		}
	}

	negativeIndex := positiveIndex - 1

	for i := range A {
		A[i] *= A[i]
	}

	result := make([]int, 0, len(A))

	for {
		positiveEnable := false
		var positiveValue int
		if positiveIndex < len(A) {
			positiveEnable = true
			positiveValue = A[positiveIndex]
		}

		negativeEnable := false
		var negativeValue int
		if negativeIndex >= 0 {
			negativeEnable = true
			negativeValue = A[negativeIndex]
		}

		if positiveEnable && negativeEnable {
			if positiveValue < negativeValue {
				result = append(result, positiveValue)
				positiveIndex++
			} else {
				result = append(result, negativeValue)
				negativeIndex--
			}
			continue
		}

		if positiveEnable {
			result = append(result, positiveValue)
			positiveIndex++
			continue
		}

		if negativeEnable {
			result = append(result, negativeValue)
			negativeIndex--
			continue
		}

		break
	}

	return result
}
