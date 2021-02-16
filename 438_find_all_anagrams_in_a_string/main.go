package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
	fmt.Println(findAnagrams("acdcaeccde", "c"))
}

const (
	leftState  = "left"
	rightState = "right"
)

func findAnagrams(s string, p string) []int {
	needs, windows := getNeedsWindows(p)
	// fmt.Println(needs, windows)

	var left, right int

	// if count == len(s) mean windows match needs
	var count int

	// init move right
	state := rightState

	result := make([]int, 0, 10)

	for right < len(s) {
		// fmt.Println(left, right, state)

		if state == rightState {
			rr := string(s[right])

			if _, ok := needs[rr]; ok {
				if windows[rr] < needs[rr] {
					count++
				}

				windows[rr]++
			}
		}

		if state == leftState {
			ll := string(s[left-1])

			if _, ok := needs[ll]; ok {
				if windows[ll] <= needs[ll] {
					count--
				}

				windows[ll]--
			}
		}

		// fmt.Println(windows, count)

		if count == len(p) {
			if right-left+1 == len(p) {
				result = append(result, left)
			}

			left++
			state = leftState
		} else {
			right++
			state = rightState
		}
	}

	return result
}

// break v to needs, windows
// v = ABCA
// needs = {A: 2; B: 1; C: 1}
// windows = {A: 0, B: 0, C: 0}
func getNeedsWindows(v string) (needs map[string]int, windows map[string]int) {
	needs = make(map[string]int)
	windows = make(map[string]int)
	for _, c := range v {
		cc := string(c)

		if _, ok := needs[cc]; !ok {
			needs[cc] = 0
		}
		needs[cc]++

		windows[cc] = 0
	}
	return needs, windows
}
