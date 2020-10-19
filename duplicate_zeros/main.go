package main

func main() {}

func duplicateZeros_1(arr []int) {
	for i := 0; i < len(arr); {
		if arr[i] != 0 {
			i++
			continue
		}

		// a b c -> 0 a b

		for j := len(arr) - 1; j > i+1; j-- {
			arr[j] = arr[j-1]
		}

		if i+1 < len(arr) {
			arr[i+1] = 0
		}

		i += 2
	}
}
