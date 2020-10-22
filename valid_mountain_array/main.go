package main

func main() {}

func validMountainArray_1(A []int) bool {
	i := 0
	for i < len(A)-1 {
		if A[i] == A[i+1] {
			return false
		}

		if A[i] > A[i+1] {
			break
		}

		i++
	}

	if i == 0 || i == len(A)-1 {
		return false
	}

	j := i + 1
	for j < len(A)-1 {
		if A[j] <= A[j+1] {
			return false
		}

		j++
	}

	return true
}
