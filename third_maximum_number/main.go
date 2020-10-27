package main

import "fmt"

func main() {}

func thirdMax_1(nums []int) int {
	max_1 := nums[0]

	var max_2, max_3 int
	init_2 := false
	init_3 := false

	for i := 1; i < len(nums); i++ {
		fmt.Println(nums[i], max_1, max_2, init_2, max_3, init_3)

		if max_1 < nums[i] {
			if init_2 {
				max_3 = max_2
				init_3 = true
			}
			max_2 = max_1
			init_2 = true
			max_1 = nums[i]
			continue
		}

		if max_1 == nums[i] {
			continue
		}

		if !init_2 {
			max_2 = nums[i]
			init_2 = true
			continue
		}

		if max_2 < nums[i] {
			max_3 = max_2
			init_3 = true
			max_2 = nums[i]
			init_2 = true
			continue
		}

		if max_2 == nums[i] {
			continue
		}

		if !init_3 || (init_3 && max_3 < nums[i]) {
			max_3 = nums[i]
			init_3 = true
			continue
		}

		if max_3 == nums[i] {
			continue
		}
	}

	if !init_3 {
		return max_1
	}

	return max_3
}
