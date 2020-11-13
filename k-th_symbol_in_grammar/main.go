package main

func main() {}

// 0 -> 01
// 1 -> 10
func kthGrammar_1(N int, K int) int {
	if N == 1 {
		return 0
	}

	prev := kthGrammar_1(N-1, (K+1)/2)

	if K%2 != 0 {
		return prev
	}

	return 1 - prev
}
