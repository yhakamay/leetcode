package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1, m, nums2, n := []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3
	merge(nums1, m, nums2, n)
	fmt.Printf("nums1: %v\n", nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i, j := m, 0; i < m+n; {
		nums1[i] = nums2[j]
		i++
		j++
	}

	sort.Ints(nums1)
}
