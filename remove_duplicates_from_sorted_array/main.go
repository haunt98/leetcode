package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates_1(arr))
	fmt.Println(arr)
}

func removeDuplicates_1(nums []int) int {
	newLen := len(nums)

	for i := 0; i < newLen; i++ {
		dupCount := 0
		for j := i + 1; j < newLen; j++ {
			if nums[j] != nums[i] {
				break
			}

			dupCount++
		}

		if dupCount == 0 {
			continue
		}

		newLen -= dupCount

		for j := i + 1; j < newLen; j++ {
			nums[j] = nums[j+dupCount]
		}

	}

	return newLen
}
