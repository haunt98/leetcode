package main

import "fmt"

func main() {
	fmt.Println("11. Container With Most Water")
}

func maxArea(height []int) int {
	max := -1

	i := 0
	j := len(height) - 1
	for i < j {
		max = max2(max, min2(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return max
}

func maxAreaSlow(height []int) int {
	max := -1

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			water := min2(height[i], height[j]) * (j - i)
			if max < water {
				max = water
			}
		}
	}

	return max
}

func min2(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max2(a, b int) int {
	if a > b {
		return a
	}

	return b
}
