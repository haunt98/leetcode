package main

import "fmt"

func main() {
	arr := []int{3, 2, 2, 3}
	fmt.Println(removeElement_1(arr, 3))
}

func removeElement_1(nums []int, val int) int {
	index := 0
	newLen := len(nums)

	for {
		if index >= newLen {
			break
		}

		if nums[index] != val {
			index++
			continue
		}

		for j := index; j < newLen-1; j++ {
			nums[j] = nums[j+1]
		}
		newLen--
	}

	return newLen
}
