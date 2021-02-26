package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("aba"))
	fmt.Println(isPalindrome("ab"))

	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("ac"))
	fmt.Println(longestPalindrome("aacabdkacaa"))
	fmt.Println(longestPalindrome("babaddtattarrattatddetartrateedredividerb"))
}

func longestPalindrome(s string) string {
	// return longestPalindromeBruteForce(s)
	return longestPalindromeDynamic(s)
}

func longestPalindromeDynamic(s string) string {
	// memo[i][j] == true mean s[i:j+1] is palindrome
	memo := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		memo[i] = make([]bool, len(s))
	}

	max := ""

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			memo[i][j] = isPalindromeWithMemo(s, i, j, memo)
			if !memo[i][j] {
				continue
			}

			if j-i+1 > len(max) {
				max = s[i : j+1]
			}
		}
	}

	return max
}

func isPalindromeWithMemo(s string, i, j int, memo [][]bool) bool {
	if memo[i][j] {
		return true
	}

	if i == j {
		memo[i][j] = true
		return memo[i][j]
	}

	if j == i+1 && s[i] == s[j] {
		memo[i][j] = true
		return memo[i][j]
	}

	if i > j {
		return false
	}

	if s[i] != s[j] {
		return false
	}

	memo[i][j] = isPalindromeWithMemo(s, i+1, j-1, memo)
	return memo[i][j]
}

func longestPalindromeBruteForce(s string) string {
	max := ""

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			sub := s[i : j+1]
			// fmt.Println(sub)

			if !isPalindrome(sub) {
				continue
			}

			if len(sub) > len(max) {
				max = s[i : j+1]
			}
		}
	}

	return max
}

func isPalindrome(s string) bool {
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
