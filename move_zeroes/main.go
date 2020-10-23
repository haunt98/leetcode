package main

func main() {}

func moveZeroes_1(nums []int) {
	i := 0
	j := i + 1
	for j < len(nums) {
		if nums[i] != 0 {
			i++
			j++
			continue
		}

		if nums[j] == 0 {
			j++
			continue
		}

		nums[i], nums[j] = nums[j], 0
		i++
		j++
	}
}
