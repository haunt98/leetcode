package main

func main() {}

func removeDuplicates_1(nums []int) int {
	i := 0
	j := i + 1

	for j < len(nums) {
		if nums[j] == nums[i] {
			j++
			continue
		}

		nums[i+1] = nums[j]
		i++
		j++
	}

	return i + 1
}
