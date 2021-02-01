package main

import "fmt"

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

const (
	leftState  = "left"
	rightState = "right"
)

// move right until windows get all needs
// move left until windows no more satisfy needs
func minWindow(s string, t string) string {
	needs, windows := getNeedsWindows(t)

	var left, right int

	// count == len(t) then windows match needs
	count := 0

	// move left or right
	// init right state
	state := rightState

	var minResult string
	isSetMin := false

	for left <= right && right < len(s) {
		// fmt.Println(left, right, state)
		// fmt.Println(windows)

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
			// check previous left
			ll := string(s[left-1])

			if _, ok := needs[ll]; ok {
				if windows[ll] <= needs[ll] {
					count--
				}

				windows[ll]--
			}
		}

		// fmt.Println("count", count)

		if count == len(t) {
			if !isSetMin {
				minResult = s[left : right+1]
				isSetMin = true
			} else {
				if len(minResult) > right-left+1 {
					minResult = s[left : right+1]
				}
			}

			left++
			state = leftState
		} else {
			right++
			state = rightState
		}
	}

	return minResult
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
