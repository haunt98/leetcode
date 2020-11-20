package main

func main() {}

func smallerNumbersThanCurrent_1(nums []int) []int {
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}

			if nums[j] < nums[i] {
				count++
			}
		}
		result[i] = count
	}

	return result
}

func smallerNumbersThanCurrent_2(nums []int) []int {
	appears := make([]int, 101)

	// count
	for i := 0; i < len(nums); i++ {
		appears[nums[i]]++
	}

	for i := 1; i < len(appears); i++ {
		appears[i] += appears[i-1]
	}

	// 0 1 2 3 4
	// 2 2 1 2 4
	// ->
	// 0 1 2 3 4
	// 2 4 5 7 11
	// < 1 have 2
	// < 2 have 4
	// < 3 have 5
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			result[i] = 0
			continue
		}

		result[i] = appears[nums[i]-1]
	}

	return result
}
