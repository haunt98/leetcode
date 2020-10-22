package main

func main() {}

func checkIfExist_1(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		isEven := arr[i]%2 == 0
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == arr[i]*2 {
				return true
			}

			if isEven && arr[j] == arr[i]/2 {
				return true
			}
		}
	}

	return false
}
