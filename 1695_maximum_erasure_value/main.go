package main

import "fmt"

func main() {
	fmt.Println(maximumUniqueSubarray([]int{4, 2, 4, 5, 6}))
	fmt.Println(maximumUniqueSubarray([]int{5, 2, 1, 2, 5, 2, 1, 2, 5}))
}

const (
	leftState  = "left"
	rightState = "right"
)

func maximumUniqueSubarray(nums []int) int {
	track := make(map[int]int)

	var left, right int

	// init move right
	state := rightState

	var countRepeat int
	var sum int

	maxSum := 0

	for right < len(nums) {
		// fmt.Println(left, right, state)

		if state == rightState {
			rr := nums[right]

			track[rr]++

			// only count one time when move from 1 -> 2
			if track[rr] == 2 {
				countRepeat++
			}

			sum += rr
		}

		if state == leftState {
			ll := nums[left-1]

			track[ll]--

			// only count one time when move from 2 -> 1
			if track[ll] == 1 {
				countRepeat--
			}

			sum -= ll
		}

		// fmt.Println(track, countRepeat)

		if countRepeat == 0 {
			if sum > maxSum {
				maxSum = sum
			}

			right++
			state = rightState
		} else {
			left++
			state = leftState
		}
	}

	return maxSum
}
