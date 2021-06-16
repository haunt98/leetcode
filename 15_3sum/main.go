package main

import "fmt"

func main() {
	fmt.Println()
}

func threeSum(nums []int) [][]int {
	sortArr(nums)
	// fmt.Println(nums)

	result := make([][]int, 0, 100)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 2 pointers

		j := i + 1
		if j >= len(nums) {
			break
		}

		k := len(nums) - 1

		// fmt.Println("init", i, j, k)

		for j < k {
			// fmt.Println("before", i, j, k)

			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}

			if k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
				continue
			}

			// fmt.Println("after", i, j, k)

			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				// fmt.Println("ok", i, j, k)
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return result
}

func sortArr(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
