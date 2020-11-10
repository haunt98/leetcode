package main

import "fmt"

func main() {
	fmt.Println(myPow_2(2.0, 10))
}

func myPow_1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		return 1 / myPow_1(x, -n)
	}

	return x * myPow_1(x, n-1)
}

func myPow_2(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		return 1 / myPow_2(x, -n)
	}

	// x ^ n = (x^2) ^ (n/2)
	if n%2 == 0 {
		return myPow_2(x*x, n>>1)
	}

	return x * myPow_2(x*x, n>>1)
}
