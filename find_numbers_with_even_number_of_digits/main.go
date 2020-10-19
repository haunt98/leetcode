package main

func main() {}

func findNumbers_1(nums []int) int {
	count := 0

	for i := range nums {
		if hasEvenNumberDigits(nums[i]) {
			count++
		}
	}

	return count
}

func hasEvenNumberDigits(num int) bool {
	if num == 0 {
		return false
	}

	numberDigits := 0

	for num > 0 {
		num /= 10
		numberDigits++
	}

	return numberDigits%2 == 0
}
