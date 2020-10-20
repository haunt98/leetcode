package main

func main() {}

func merge_1(nums1 []int, m int, nums2 []int, n int) {
	totalLen := m
	maxLen := m + n
	nums1Index := 0
	nums2Index := 0

	for {
		if totalLen == maxLen {
			return
		}

		if nums1Index >= totalLen {
			nums1[nums1Index] = nums2[nums2Index]
			totalLen++
			nums1Index++
			nums2Index++
			continue
		}

		if nums1[nums1Index] <= nums2[nums2Index] {
			nums1Index++
			continue
		}

		for i := totalLen; i > nums1Index; i-- {
			nums1[i] = nums1[i-1]
		}

		nums1[nums1Index] = nums2[nums2Index]
		totalLen++
		nums1Index++
		nums2Index++
	}
}
