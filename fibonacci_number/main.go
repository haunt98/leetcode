package main

func main() {}

func fib_1(N int) int {
	if N <= 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	return fib_1(N-1) + fib_1(N-2)
}

func fib_2(N int) int {
	if N <= 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	arr := make([]int, N+1)
	arr[0] = 0
	arr[1] = 1

	for i := 2; i <= N; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr[N]
}
