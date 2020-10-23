package main

func main() {}

func replaceElements_1(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		max := arr[i+1]
		for j := i + 2; j < len(arr); j++ {
			if max < arr[j] {
				max = arr[j]
			}
		}
		arr[i] = max
	}

	arr[len(arr)-1] = -1

	return arr
}
