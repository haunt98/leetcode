package main

import (
	"fmt"
)

func main() {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
}

func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			// delete i
			withoutI := strictPalindrome(s[i+1 : j+1])
			// delete j
			withoutJ := strictPalindrome(s[i:j])
			if withoutI || withoutJ {
				return true
			}

			return false
		}

		i++
		j--
	}

	return true
}

func strictPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
