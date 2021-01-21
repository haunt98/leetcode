package main

import "fmt"

func main() {
	fmt.Println("1. Two Sum")
}

func twoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return nil
}
