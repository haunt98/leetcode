package main

import "fmt"

func main() {
	fmt.Println("204. Count Primes")
}

func countPrimes(n int) int {
	count := 0

	marks := make([]bool, n)

	for i := 2; i < n; i++ {
		if marks[i] == true {
			continue
		}

		count++
		for j := i; j < n; j += i {
			marks[j] = true
		}
	}

	return count
}
