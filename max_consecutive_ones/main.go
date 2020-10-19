package main

func main() {}

func findMaxConsecutiveOnes_1(nums []int) int {
	maxConsecutive := 0
	lenNums := len(nums)

	for i := 0; i < lenNums; {
		countConsecutive := 0

		j := i
		for ; j < lenNums; j++ {
			if nums[j] != 1 {
				break
			}

			countConsecutive++
		}
		i = j + 1

		if maxConsecutive < countConsecutive {
			maxConsecutive = countConsecutive
		}
	}

	return maxConsecutive
}
