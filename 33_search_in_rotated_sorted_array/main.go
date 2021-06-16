package main

import "fmt"

func main() {
	fmt.Println("33. Search in Rotated Sorted Array")
}

// The idea is find the rotate index
// Then make the nums look like one
func search(nums []int, target int) int {
	rotateIndex := getRotateIndex(nums)

	result := binarySearchRotate(nums, target, rotateIndex, rotateIndex+len(nums)-1)
	if result != -1 {
		result %= len(nums)
	}
	return result
}

func getRotateIndex(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			return i + 1
		}
	}

	return 0
}

// Always check the rotate index
func binarySearchRotate(nums []int, target, from, to int) int {
	if from > to {
		return -1
	}

	if from == to {
		if nums[from%len(nums)] == target {
			return from
		}

		return -1
	}

	mid := from + (to-from)/2
	if nums[mid%len(nums)] == target {
		return mid
	} else if nums[mid%len(nums)] > target {
		return binarySearchRotate(nums, target, from, mid-1)
	} else {
		return binarySearchRotate(nums, target, mid+1, to)
	}
}
