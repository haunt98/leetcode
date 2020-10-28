package main

import "fmt"

func main() {
	fmt.Println(findDisappearedNumbers_1([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDisappearedNumbers_1(nums []int) []int {
	m := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, ok := m[num]; ok {
			continue
		}

		m[num] = struct{}{}
	}

	arr := make([]int, 0, len(nums))

	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; ok {
			continue
		}

		arr = append(arr, i)
	}

	return arr
}
