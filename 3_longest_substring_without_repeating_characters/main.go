package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
}

const (
	leftState  = "left"
	rightState = "right"
)

func lengthOfLongestSubstring(s string) int {
	track := make(map[string]int)

	var left, right int

	// init move right
	state := rightState

	var countRepeat int

	maxLen := 0

	for right < len(s) {
		// fmt.Println(left, right, state)

		if state == rightState {
			rr := string(s[right])

			track[rr]++

			// only count one time when move from 1 -> 2
			if track[rr] == 2 {
				countRepeat++
			}
		}

		if state == leftState {
			ll := string(s[left-1])

			track[ll]--

			// only count one time when move from 2 -> 1
			if track[ll] == 1 {
				countRepeat--
			}
		}

		// fmt.Println(track, countRepeat)

		if countRepeat == 0 {
			if right-left+1 > maxLen {
				maxLen = right - left + 1
			}

			right++
			state = rightState
		} else {
			left++
			state = leftState
		}
	}

	return maxLen
}
